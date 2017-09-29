package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/jackc/pgx"
	elogo "github.com/kortemy/elo-go"
	"github.com/pkg/errors"
)

type Player struct {
	pid     int64
	blizzid int64
	winner  bool
}

type Game struct {
	id   int64
	mode Mode
}

type Update struct {
	pid   int64
	score int
}

/*
Elo calculation algorithm:

For each patch:

1. Fetch all games, ordered by date.
2. Fetch all players, indexable by game.
3. Make a map from playerid + game mode -> score. QM inits to 1500. When a game mode that isn't QM is first encountered, take the highest game mode score available.
4. For each game, update the skill rating of the player in the database and the map.

Repeat for next patch. Group by patch to keep memory limited.
*/
func elo(pool *pgx.ConnPool) error {
	var patch string

	const stmtName = "updatePlayers"
	_, err := pool.Prepare(stmtName, "update players set skill = $1 where id = $2")
	if err != nil {
		return err
	}

	for {
		start := time.Now()
		if err := pool.QueryRow("select id from builds where id > $1 order by id limit 1", patch).Scan(&patch); err == pgx.ErrNoRows {
			return nil
		} else if err != nil {
			return err
		}
		fmt.Println("\nstart", patch, "at", start)

		// fetch players into memory
		players := make(map[int64][]Player, 100000)

		{
			rows, err := pool.Query("SELECT id, game, blizzid, winner FROM players WHERE build = $1", patch)
			if err != nil {
				return errors.Wrap(err, "fetch players")
			}
			var game int64
			var p Player
			for rows.Next() {
				if err := rows.Scan(&p.pid, &game, &p.blizzid, &p.winner); err != nil {
					return errors.Wrap(err, "scan")
				}
				players[game] = append(players[game], p)
			}
			if err := rows.Err(); err != nil {
				return errors.Wrap(err, "rows")
			}
			fmt.Println("fetched players", time.Since(start))
		}

		g, ctx := errgroup.WithContext(context.Background())
		done := ctx.Done()
		gameCh := make(chan Game, 100)
		updateCh := make(chan Update, 1000)
		g.Go(func() error {
			defer func() { fmt.Println("fetch games done", time.Since(start)) }()
			defer close(gameCh)

			// fetch games in order
			games, err := pool.QueryEx(ctx, "select id, mode from games where build = $1 order by time", nil, patch)
			if err != nil {
				return err
			}
			var g Game
			for games.Next() {
				if err := games.Scan(&g.id, &g.mode); err != nil {
					return errors.Wrap(err, "scan")
				}

				select {
				case <-done:
					return ctx.Err()
				case gameCh <- g:
				}
			}
			return games.Err()
		})
		g.Go(func() error {
			elo := elogo.NewElo()
			count := 0
			scores := NewScores()
			defer close(updateCh)
			for game := range gameCh {
				var winSum, loseSum int
				ps := players[game.id]
				for _, p := range ps {
					score := scores.Get(game.mode, p.blizzid)
					if p.winner {
						winSum += score
					} else {
						loseSum += score
					}
				}
				winSum /= 5
				loseSum /= 5
				delta := elo.RatingDelta(winSum, loseSum, 1)
				for _, p := range ps {
					if p.winner {
						scores[game.mode][p.blizzid] += delta
					} else {
						scores[game.mode][p.blizzid] -= delta
					}
					u := Update{
						pid:   p.pid,
						score: scores[game.mode][p.blizzid],
					}
					select {
					case <-done:
						return ctx.Err()
					case updateCh <- u:
					}
				}
				count++
				if count%10000 == 0 {
					fmt.Println("progress:", count)
				}
			}

			return nil
		})

		for i := 0; i < 5; i++ {
			g.Go(func() error {
				for u := range updateCh {
					if _, err := pool.ExecEx(ctx, stmtName, nil, u.score, u.pid); err != nil {
						return err
					}
				}
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return errors.Wrap(err, "wait")
		}
		fmt.Println(patch, "took", time.Since(start))
	}
}

type Scores map[Mode]map[int64]int

func NewScores() Scores {
	s := make(Scores)
	for m := range modeNames {
		s[m] = make(map[int64]int)
	}
	return s
}

func (s Scores) Get(m Mode, player int64) int {
	const defaultScore = 1500
	score, ok := s[m][player]
	if ok {
		return score
	}
	if m == 1 {
		s[m][player] = defaultScore
		return defaultScore
	}
	score = s.Get(m-1, player)
	s[m][player] = score
	return score
}
