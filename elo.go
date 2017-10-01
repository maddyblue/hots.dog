package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/VividCortex/gohistogram"
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
func elo(dburl string) error {
	config, err := pgx.ParseURI(dburl)
	if err != nil {
		return err
	}
	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: 20,
	}
	pool, err := pgx.NewConnPool(poolConfig)
	if err != nil {
		return err
	}
	defer pool.Close()

	var patch string

	const stmtUpdatePlayers = "updatePlayers"
	if _, err := pool.Prepare(stmtUpdatePlayers, "update players set skill = $1 where id = $2"); err != nil {
		return errors.Wrap(err, "make update players")
	}
	const stmtUpdateSkills = "updateSkills"
	if _, err := pool.Prepare(stmtUpdateSkills, "upsert into playerskills (blizzid, build, mode, skill) VALUES ($1, $2, $3, $4)"); err != nil {
		return errors.Wrap(err, "make update playerskills")
	}
	scores := NewScores()
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
			fmt.Println("fetched players", len(players), time.Since(start))
		}

		ctx := context.Background()
		{
			g, gCtx := errgroup.WithContext(ctx)
			done := gCtx.Done()
			gameCh := make(chan Game, 100)
			updateCh := make(chan Update, 1000)
			g.Go(func() error {
				defer close(gameCh)

				// fetch games in order
				games, err := pool.QueryEx(gCtx, "select id, mode from games where build = $1 order by time", nil, patch)
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
						return gCtx.Err()
					case gameCh <- g:
					}
				}
				fmt.Println("fetch games done", time.Since(start))
				return games.Err()
			})
			g.Go(func() error {
				elo := elogo.NewElo()
				defer close(updateCh)
				count := 0
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
						sc := scores.Get(game.mode, p.blizzid)
						if p.winner {
							sc += delta
						} else {
							sc -= delta
						}
						scores[game.mode][p.blizzid] = patchscore{
							patch: patch,
							score: sc,
						}
						u := Update{
							pid:   p.pid,
							score: scores[game.mode][p.blizzid].score,
						}
						select {
						case <-done:
							return gCtx.Err()
						case updateCh <- u:
						}
					}
					count++
					if count%1000 == 0 {
						fmt.Println("progress:", count, time.Since(start))
					}
				}
				fmt.Println("elo done", count, time.Since(start))
				return nil
			})
			for i := 0; i < poolConfig.MaxConnections; i++ {
				g.Go(func() error {
					for u := range updateCh {
						if _, err := pool.ExecEx(gCtx, stmtUpdatePlayers, nil, u.score, u.pid); err != nil {
							return err
						}
					}
					return nil
				})
			}
			if err := g.Wait(); err != nil {
				return errors.Wrap(err, "wait")
			}
		}
		{
			// Update playerskills.
			fmt.Println("starting update playerskills", time.Since(start))
			g, gCtx := errgroup.WithContext(ctx)
			done := gCtx.Done()
			type playerskill struct {
				mode    Mode
				blizzid int64
				score   int
			}
			skillCh := make(chan playerskill, 100)
			g.Go(func() error {
				defer close(skillCh)
				var skill playerskill
				n := 0
				for mode, blizzids := range scores {
					skill.mode = mode
					for blizzid, ps := range blizzids {
						// Don't update playerskills if no games played this patch.
						if ps.patch != patch {
							continue
						}
						skill.blizzid = blizzid
						skill.score = ps.score
						select {
						case <-done:
							return gCtx.Err()
						case skillCh <- skill:
						}
						n++
					}
				}
				fmt.Println("updated skills", n, time.Since(start))
				return nil
			})
			for i := 0; i < poolConfig.MaxConnections; i++ {
				g.Go(func() error {
					for s := range skillCh {
						if _, err := pool.ExecEx(gCtx, stmtUpdateSkills, nil, s.blizzid, patch, s.mode, s.score); err != nil {
							return errors.Wrap(err, "update skills")
						}
					}
					return nil
				})
			}
			if err := g.Wait(); err != nil {
				return errors.Wrap(err, "wait")
			}
		}
		fmt.Println("getting stats", time.Since(start))
		for m, blizzids := range scores {
			// For each mode, get all the scores and sort them.
			const buckets = 20
			hist := gohistogram.NewHistogram(buckets)
			for _, sc := range blizzids {
				if sc.patch == patch {
					hist.Add(float64(sc.score))
				}
			}
			if hist.Count() == 0 {
				continue
			}
			s := Stats{
				Count:    int(hist.Count()),
				Mean:     int(hist.Mean()),
				StdDev:   int(math.Sqrt(hist.Variance())),
				Quantile: make(map[int]int),
			}
			for i := 0; i <= 100; i += 100 / buckets {
				q := float64(i) / 100
				s.Quantile[i] = int(hist.Quantile(q))
			}
			b, err := json.Marshal(s)
			if err != nil {
				return err
			}
			if _, err := pool.Exec("upsert into skillstats (build, mode, data) values ($1, $2, $3)",
				patch, m, b,
			); err != nil {
				return err
			}
		}
		fmt.Println(patch, "took", time.Since(start))
	}
}

type patchscore struct {
	score int
	patch string
}

// game mode -> blizzid -> score
type Scores map[Mode]map[int64]patchscore

func NewScores() Scores {
	s := make(Scores)
	for m := range modeNames {
		s[m] = make(map[int64]patchscore)
	}
	return s
}

func (s Scores) Get(m Mode, player int64) int {
	const defaultScore = 1500
	for ; m > 0; m-- {
		ps, ok := s[m][player]
		if ok {
			return ps.score
		}
	}
	return defaultScore
}

type Stats struct {
	Count    int
	Mean     int
	StdDev   int
	Quantile map[int]int
}
