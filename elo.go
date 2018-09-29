package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"sort"
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
func (h *hotsContext) elo() error {
	if err := h.updateInit(context.Background()); err != nil {
		return err
	}
	init := h.getInit()
	// Hard code to use the last 3 versions. This is bad but fine for now.
	updateAfterPatch := init.Builds[2].ID

	config, err := pgx.ParseURI(h.dburl)
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
	totalGames := make(map[modeRegion]struct {
		total  int
		recent int
	})

	var daysTime = time.Hour * 24 * time.Duration(daysOld)
	since := time.Now().Add(-daysTime)
	notBefore := time.Now().Add(-time.Hour * 24 * 500)
	fmt.Println("recent is since", since, ", days:", daysOld)

	for i := len(init.Builds) - 1; i >= 0; i-- {
		build := init.Builds[i]
		updatePatch := build.ID >= updateAfterPatch || *flagInit
		start := time.Now()
		patch := init.config.build(build.ID)
		isRecent := build.Start.After(since)
		fmt.Println("\nstart", build.ID, "at", start, "patch", patch)
		if build.Start.Before(notBefore) {
			fmt.Println("skipping because too old")
			continue
		}

		// fetch players into memory
		players := make(map[int64][]Player, 10000)

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
					return errors.Wrap(err, "fetch games in order")
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
				return errors.Wrap(games.Err(), "games.Err")
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
						rp := regionPlayer{game.region, p.blizzid}
						rating := scores.Get(game.mode, rp)
						if p.winner {
							teams[0].AddPlayer(p.blizzid, rating)
						} else {
							teams[1].AddPlayer(p.blizzid, rating)
						}
						mr := modeRegion{game.mode, rp}
						tg := totalGames[mr]
						tg.total++
						if isRecent {
							tg.recent++
						}
						totalGames[mr] = tg
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
						args := []interface{}{ratingToSkill(u.score), u.game, u.blizzid}
						if _, err := pool.ExecEx(gCtx, stmtUpdatePlayers, nil, args...); err != nil {
							return errors.Wrapf(err, "stmtUpdatePlayers: %v", args)
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
						if _, err := pool.ExecEx(gCtx, stmtUpdateSkills, nil, s.region, s.blizzid, patch, s.mode, ratingToSkill(s.score)); err != nil {
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
	fmt.Println("begin leaderboard")
	// Calculate leaderboard
	type leaderboardSkill struct {
		blizzid int64
		skill   float64
		total   int
		recent  int
	}
	var skills []leaderboardSkill
	const maxPlayers = 100
	for m := range init.Modes {
		for r := range init.Regions {
			skills = skills[:0]
			for rp, ps := range scores[m] {
				tg := totalGames[modeRegion{m, rp}]
				if rp.region != r || tg.recent < leaderboardMinGames {
					continue
				}
				skills = append(skills, leaderboardSkill{
					blizzid: rp.blizzid,
					skill:   ratingToSkill(ps.score),
					total:   tg.total,
					recent:  tg.recent,
				})
			}
			sort.Slice(skills, func(i, j int) bool {
				return skills[i].skill > skills[j].skill
			})
			fmt.Println("found", len(skills), "for mode", m, "region", r)
			tx, err := pool.Begin()
			if err != nil {
				return err
			}
			if _, err := tx.Exec("delete from leaderboard where region = $1 and mode = $2", r, m); err != nil {
				return err
			}
			for i, s := range skills {
				if i >= 100 {
					break
				}
				if _, err := tx.Exec(`
					INSERT INTO leaderboard
					(region, mode, rank, blizzid, skill, total, recent)
					VALUES
					($1, $2, $3, $4, $5, $6, $7)
					RETURNING NOTHING`,
					r,
					m,
					i+1,
					s.blizzid,
					s.skill,
					s.total,
					s.recent,
				); err != nil {
					return err
				}
			}
			if err := tx.Commit(); err != nil {
				return err
			}
		}
	}
	return nil
}

func ratingToSkill(r skills.Rating) float64 {
	return r.Mean()
}

type patchscore struct {
	score skills.Rating
	patch string
}

type modeRegion struct {
	mode Mode
	rp   regionPlayer
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
