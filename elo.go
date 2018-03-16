package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/ChrisHines/GoSkills/skills"
	"github.com/ChrisHines/GoSkills/skills/trueskill"
	"github.com/VividCortex/gohistogram"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

type Player struct {
	blizzid int64
	winner  bool
}

type Game struct {
	id     int64
	mode   Mode
	region int
}

type Update struct {
	game    int64
	blizzid int64
	score   skills.Rating
}

// https://www.reddit.com/r/heroesofthestorm/comments/82p1h5/ranked_player_distribution/dvbzypz/
var skillQuantiles = []int{
	0,
	7,
	42,
	77,
	92,
	99,
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
func (h *hotsContext) elo(dburl, updateAfterPatch string) error {
	if err := h.updateInit(context.Background()); err != nil {
		return err
	}
	init := h.getInit()

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

	const stmtUpdatePlayers = "updatePlayers"
	if _, err := pool.Prepare(stmtUpdatePlayers, "update players set skill = $1 where game = $2 and blizzid = $3"); err != nil {
		return errors.Wrap(err, "make update players")
	}
	const stmtUpdateSkills = "updateSkills"
	if _, err := pool.Prepare(stmtUpdateSkills, "upsert into playerskills (region, blizzid, build, mode, skill) VALUES ($1, $2, $3, $4, $5)"); err != nil {
		return errors.Wrap(err, "make update playerskills")
	}
	scores := NewScores()
	for i := len(init.Builds) - 1; i >= 0; i-- {
		build := init.Builds[i]
		updatePatch := build.ID >= updateAfterPatch
		start := time.Now()
		patch := init.config.build(build.ID)
		fmt.Println("\nstart", build.ID, "at", start, "patch", patch)

		// fetch players into memory
		players := make(map[int64][]Player, 100000)

		{
			rows, err := pool.Query("SELECT game, blizzid, winner FROM players WHERE build = $1", patch)
			if err != nil {
				return errors.Wrap(err, "fetch players")
			}
			var game int64
			var p Player
			for rows.Next() {
				if err := rows.Scan(&game, &p.blizzid, &p.winner); err != nil {
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
				games, err := pool.QueryEx(gCtx, "select id, mode, region from games where build = $1 order by time", nil, patch)
				if err != nil {
					return err
				}
				var g Game
				for games.Next() {
					if err := games.Scan(&g.id, &g.mode, &g.region); err != nil {
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
				defer close(updateCh)
				count := 0
				var calc trueskill.TwoTeamCalc
				teams := make([]skills.Team, 2)
				for game := range gameCh {
					count++
					teams[0] = skills.NewTeam()
					teams[1] = skills.NewTeam()
					ps := players[game.id]
					for _, p := range ps {
						rating := scores.Get(game.mode, regionPlayer{game.region, p.blizzid})
						if p.winner {
							teams[0].AddPlayer(p.blizzid, rating)
						} else {
							teams[1].AddPlayer(p.blizzid, rating)
						}
					}
					if teams[0].PlayerCount() != 5 || teams[1].PlayerCount() != 5 {
						continue
					}
					ratings := calc.CalcNewRatings(skills.DefaultGameInfo, teams, 1, 2)
					for _, p := range ps {
						rating, ok := ratings[p.blizzid]
						if !ok {
							return errors.Errorf("unfound id: %d", p.blizzid)
						}
						scores[game.mode][regionPlayer{game.region, p.blizzid}] = patchscore{
							patch: patch,
							score: rating,
						}
						if updatePatch {
							u := Update{
								game:    game.id,
								blizzid: p.blizzid,
								score:   scores[game.mode][regionPlayer{game.region, p.blizzid}].score,
							}
							select {
							case <-done:
								return gCtx.Err()
							case updateCh <- u:
							}
						}
					}
					if updatePatch && count%1000 == 0 {
						fmt.Println("progress:", count, time.Since(start))
					}
					// See if we are done.
					select {
					case <-done:
						return gCtx.Err()
					default:
					}
				}
				fmt.Println("elo done", count, time.Since(start))
				return nil
			})
			for i := 0; i < poolConfig.MaxConnections; i++ {
				g.Go(func() error {
					for u := range updateCh {
						if _, err := pool.ExecEx(gCtx, stmtUpdatePlayers, nil, ratingToInt(u.score), u.game, u.blizzid); err != nil {
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
		if updatePatch {
			// Update playerskills.
			fmt.Println("starting update playerskills", time.Since(start))
			g, gCtx := errgroup.WithContext(ctx)
			done := gCtx.Done()
			type playerskill struct {
				region  int
				mode    Mode
				blizzid int64
				score   skills.Rating
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
						skill.region = blizzid.region
						skill.blizzid = blizzid.blizzid
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
						if _, err := pool.ExecEx(gCtx, stmtUpdateSkills, nil, s.region, s.blizzid, patch, s.mode, ratingToInt(s.score)); err != nil {
							return errors.Wrap(err, "update skills")
						}
					}
					return nil
				})
			}
			if err := g.Wait(); err != nil {
				return errors.Wrap(err, "wait")
			}
			fmt.Println("getting stats", time.Since(start))
			for m, blizzids := range scores {
				// For each mode, get all the scores and sort them.
				const buckets = 50
				hist := gohistogram.NewHistogram(buckets)
				for _, sc := range blizzids {
					if sc.patch == patch {
						hist.Add(sc.score.Mean())
					}
				}
				if hist.Count() == 0 {
					continue
				}
				s := Stats{
					Count:    int(hist.Count()),
					Mean:     hist.Mean(),
					StdDev:   math.Sqrt(hist.Variance()),
					Quantile: make(map[int]float64),
				}
				for _, q := range skillQuantiles {
					s.Quantile[q] = hist.Quantile(float64(q) / 100)
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
		}
		fmt.Println(build.ID, "took", time.Since(start))
	}
	return nil
}

func ratingToInt(r skills.Rating) int64 {
	return meanToInt(r.Mean())
}

func meanToInt(f float64) int64 {
	return int64(f * meanMult)
}

const meanMult = 1e6

type patchscore struct {
	score skills.Rating
	patch string
}

type regionPlayer struct {
	region  int
	blizzid int64
}

// game mode -> blizzid -> score
type Scores map[Mode]map[regionPlayer]patchscore

func NewScores() Scores {
	s := make(Scores)
	for m := range modeNames {
		s[m] = make(map[regionPlayer]patchscore)
	}
	return s
}

func (s Scores) Get(m Mode, player regionPlayer) skills.Rating {
	for ; m > 0; m-- {
		ps, ok := s[m][player]
		if ok {
			return ps.score
		}
	}
	return skills.DefaultGameInfo.DefaultRating()
}

type Stats struct {
	Count    int
	Mean     float64
	StdDev   float64
	Quantile map[int]float64
}
