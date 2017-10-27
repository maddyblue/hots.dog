package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/lib/pq"
	"github.com/pkg/errors"

	"cloud.google.com/go/storage"
)

const (
	HotsApi        = "https://hotsapi.net/api/v1"
	configJSON     = "config.json"
	dirGame        = "game"
	dirPlayer      = "player"
	dirUnprocessed = "unprocessed"
	perFile        = 10000
	perRequest     = 100
	configBase     = "%09d.csv"
)

func (h *hotsContext) updateDB() error {
	ctx := context.Background()
	cl, err := storage.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "new client")
	}
	bucket := cl.Bucket(*flagImport)
	for {
		err := h.updateDBNext(bucket)
		if err == storage.ErrObjectNotExist {
			fmt.Println("sleeping")
			time.Sleep(time.Minute * 10)
		} else if err != nil {
			return err
		}
	}
}

func (h *hotsContext) updateDBNext(bucket *storage.BucketHandle) error {
	ctx := context.Background()
	const nextUpdateKey = "next-update"
	var start int
	if err := h.x.Get(&start, `SELECT i FROM config WHERE key = $1`, nextUpdateKey); err != nil {
		return err
	}
	file := fmt.Sprintf(configBase, start)
	fmt.Println("updating block", start, file)

	if _, err := h.db.Exec(`DELETE FROM games WHERE ID >= $1 AND ID < $1 + $2`, start, perFile); err != nil {
		return errors.Wrap(err, "clear games")
	}
	for {
		res, err := h.db.Exec(`DELETE FROM players WHERE game >= $1 AND game < $1 + $2 limit 1000`, start, perFile)
		if err != nil {
			return errors.Wrap(err, "clear players")
		}
		count, err := res.RowsAffected()
		if err != nil {
			return errors.Wrap(err, "clear players")
		}
		if count == 0 {
			break
		}
		fmt.Println("CLEARED", count, "players")
	}

	fmt.Println("fetching csvs")
	var games, players [][]string
	{
		r, err := bucket.Object("game/" + file).NewReader(ctx)
		if err != nil {
			return err
		}
		defer r.Close()
		games, err = csv.NewReader(r).ReadAll()
		if err != nil {
			return err
		}
	}
	{
		r, err := bucket.Object("player/" + file).NewReader(ctx)
		if err != nil {
			return err
		}
		defer r.Close()
		players, err = csv.NewReader(r).ReadAll()
		if err != nil {
			return err
		}
	}

	copyin := func(data [][]string, table string, cols []string) error {
		const size = 500
		vv := make([]interface{}, len(cols))
		fmt.Println(table, len(data))
		for len(data) > 0 {
			nextLen := size
			if nextLen > len(data) {
				nextLen = len(data)
			}
			next := data[:nextLen]
			data = data[nextLen:]
			fmt.Println(start, table, len(next), len(data))
			count := 0
			if err := retry(func() error {
				if count > 0 {
					fmt.Println("retry", count, len(data))
				}
				count++
				txn, err := h.db.Begin()
				if err != nil {
					return errors.Wrap(err, "begin")
				}
				defer txn.Rollback()
				stmt, err := txn.Prepare(pq.CopyIn(table, cols...))
				if err != nil {
					return errors.Wrap(err, "prepare")
				}
				defer stmt.Close()
				for _, r := range next {
					for ii, v := range r {
						vv[ii] = v
					}
					_, err = stmt.Exec(vv...)
					if err != nil {
						return errors.Wrap(err, "exec")
					}
				}
				if _, err := stmt.Exec(); err != nil {
					return errors.Wrap(err, "exec")
				}
				if err := stmt.Close(); err != nil {
					return errors.Wrap(err, "close")
				}
				if err := txn.Commit(); err != nil {
					return errors.Wrap(err, "commit")
				}
				return nil
			}); err != nil {
				return errors.Wrap(err, "retry")
			}
		}
		return nil
	}
	if err := copyin(games, "games", []string{"id", "mode", "time", "map", "length", "build", "bans"}); err != nil {
		return errors.Wrap(err, "copy games")
	}
	if err := copyin(players, "players", []string{"game", "mode", "time", "map", "length", "build", "hero", "hero_level", "team", "winner", "blizzid", "skill", "battletag", "talents"}); err != nil {
		return errors.Wrap(err, "copy players")
	}
	if _, err := h.db.Exec(`UPDATE config SET i = $1 WHERE key = $2`, start+perFile, nextUpdateKey); err != nil {
		return errors.Wrap(err, "update config")
	}

	return nil
}

func updateNew(bucketName string) error {
	ctx := context.Background()
	cl, err := storage.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "new client")
	}
	bucket := cl.Bucket(bucketName)
	config, err := getConfig(ctx, bucket)
	if err != nil {
		return errors.Wrap(err, "get config")
	}
	for {
		_, err := updateNextGroup(ctx, bucket, config)
		if err == ErrNotEnough {
			fmt.Println("next block not done, sleeping")
			time.Sleep(time.Hour)
		} else if err != nil {
			log.Println(err)
			time.Sleep(time.Hour)
		}
	}
}

var ErrNotEnough = errors.New("not enough results")

func updateNextGroup(ctx context.Context, bucket *storage.BucketHandle, config *groupConfig) (int, error) {
	start := config.NextID
	until := start + perFile
	fmt.Println("UNG", start, until)
	if start%perFile != 0 {
		return 0, errors.Errorf("bad start id: %d", start)
	}
	base := fmt.Sprintf(configBase, start)
	gw := bucket.Object(path.Join(dirGame, base)).NewWriter(ctx)
	pw := bucket.Object(path.Join(dirPlayer, base)).NewWriter(ctx)
	gc := csv.NewWriter(gw)
	pc := csv.NewWriter(pw)
	{
		idCh := make(chan int)
		replaysCh := make(chan []byte, 5)
		replayCh := make(chan Replay, 100)
		g, ctx := errgroup.WithContext(ctx)
		done := ctx.Done()
		// Generate IDs.
		g.Go(func() error {
			defer close(idCh)
			for id := start; id < start+perFile; id += perRequest {
				select {
				case <-done:
					return ctx.Err()
				case idCh <- id:
				}
			}
			return nil
		})
		// Fetch blocks of data, send just bytes on the chan so we can start next request asap.
		g.Go(func() error {
			defer close(replaysCh)
			return groupWorkers(ctx, cap(replaysCh), func(ctx context.Context) error {
				for id := range idCh {
					url := fmt.Sprintf(HotsApi+"/replays?with_players=true&min_id=%d", id)
					b, err := httpGet(ctx, url)
					if err != nil {
						return errors.Wrapf(err, "http get: %s", b)
					}
					select {
					case <-done:
						return ctx.Err()
					case replaysCh <- b:
					}
				}
				return nil
			})
		})
		// Unmarshal bytes, send as replays.
		g.Go(func() error {
			defer close(replayCh)
			for b := range replaysCh {
				var replays Replays
				if err := json.Unmarshal(b, &replays); err != nil {
					return errors.Wrap(err, "json decode")
				}
				for _, r := range replays {
					if r.ID >= until {
						return nil
					}
					select {
					case <-done:
						return ctx.Err()
					case replayCh <- r:
					}
				}
				if len(replays) != perRequest {
					return ErrNotEnough
				}
			}
			return nil
		})
		g.Go(func() error {
			seen := make(map[int]bool)
			for r := range replayCh {
				if seen[r.ID] {
					continue
				}
				seen[r.ID] = true
				mode, ok := gameModes[r.GameType]
				if !ok {
					continue
				}
				/*
					if !r.Processed {
						if err := bucket.Object(path.Join(dirUnprocessed, strconv.Itoa(r.ID))).NewWriter(ctx).Close(); err != nil {
							return errors.Wrapf(err, "write unprocessed: %d", r.ID)
						}
						continue
					}
				*/
				var bans []string
				for _, bs := range r.Bans {
					for _, b := range bs {
						if b != "" {
							bans = append(bans, config.hero(b))
						}
					}
				}
				config.addBuild(r.GameVersion, time.Time(r.GameDate))
				common := []string{
					fmt.Sprint(r.ID),
					fmt.Sprint(mode),
					time.Time(r.GameDate).Format(time.RFC3339Nano),
					config.gamemap(r.GameMap),
					fmt.Sprint(r.GameLength),
					config.build(r.GameVersion),
				}
				if err := gc.Write(append(common,
					fmt.Sprintf("{%s}", strings.Join(bans, ",")),
				)); err != nil {
					return errors.Wrap(err, "gc write")
				}
				for _, p := range r.Players {
					var talents []string
					for _, t := range []string{
						p.Talents.Num1,
						p.Talents.Num4,
						p.Talents.Num7,
						p.Talents.Num10,
						p.Talents.Num13,
						p.Talents.Num16,
						p.Talents.Num20,
					} {
						if t == "" {
							break
						}
						talents = append(talents, config.talent(t))
					}
					if err := pc.Write(append(common,
						config.hero(p.Hero),
						fmt.Sprint(p.HeroLevel),
						fmt.Sprint(p.Team),
						fmt.Sprint(p.Winner),
						fmt.Sprint(p.BlizzID),
						"0", // skill
						p.Battletag,
						fmt.Sprintf("{%s}", strings.Join(talents, ",")),
					)); err != nil {
						return errors.Wrap(err, "gc write")
					}
				}
			}
			return nil
		})
		if err := g.Wait(); err != nil {
			gw.CloseWithError(err)
			pw.CloseWithError(err)
			return 0, err
		}
	}
	gc.Flush()
	pc.Flush()
	if err := gc.Error(); err != nil {
		return 0, errors.Wrap(err, "gc flush")
	}
	if err := pc.Error(); err != nil {
		return 0, errors.Wrap(err, "pc flush")
	}
	if err := gw.Close(); err != nil {
		return 0, errors.Wrap(err, "gw close")
	}
	if err := pw.Close(); err != nil {
		return 0, errors.Wrap(err, "pw close")
	}
	cw := bucket.Object(configJSON).NewWriter(ctx)
	config.NextID = until
	if err := json.NewEncoder(cw).Encode(config); err != nil {
		return 0, errors.Wrap(err, "json encode")
	}
	if err := cw.Close(); err != nil {
		return 0, errors.Wrap(err, "write config")
	}
	return config.NextID, nil
}

// groupWorkers creates num worker go routines in an error group.
func groupWorkers(ctx context.Context, num int, f func(context.Context) error) error {
	group, ctx := errgroup.WithContext(ctx)
	for i := 0; i < num; i++ {
		group.Go(func() error {
			return f(ctx)
		})
	}
	return group.Wait()
}

type groupConfig struct {
	sync.Mutex
	readonly bool
	NextID   int
	Map      map[string]map[string]string
	Builds   map[string]dateRange
}

type dateRange struct {
	Start time.Time
	End   time.Time
}

func (g *groupConfig) addBuild(name string, date time.Time) {
	if g.Builds == nil {
		g.Builds = make(map[string]dateRange)
	}
	b, ok := g.Builds[name]
	if !ok {
		b.Start = date
		b.End = date
	} else if date.Before(b.Start) {
		b.Start = date
	} else if date.After(b.End) {
		b.End = date
	}
	g.Builds[name] = b
}

func (g *groupConfig) build(name string) string {
	return g.get("build", name)
}

func (g *groupConfig) gamemap(name string) string {
	return g.get("map", name)
}

func (g *groupConfig) talent(name string) string {
	return g.get("talent", name)
}

func (g *groupConfig) hero(name string) string {
	return g.get("hero", name)
}

func (g *groupConfig) get(group, name string) string {
	if !g.readonly {
		g.Lock()
		defer g.Unlock()
		if g.Map == nil {
			g.Map = make(map[string]map[string]string)
		}
		if g.Map[group] == nil {
			g.Map[group] = make(map[string]string)
		}
	} else {
		if g.Map == nil || g.Map[group] == nil {
			return name
		}
	}
	id := g.Map[group][name]
	if !g.readonly && id == "" {
		id = strconv.Itoa(len(g.Map[group]) + 1)
		g.Map[group][name] = id
	}
	return id
}

func (g *groupConfig) toMap(group string) map[string]string {
	return g.Map[group]
}

func getConfig(ctx context.Context, bucket *storage.BucketHandle) (*groupConfig, error) {
	var config groupConfig
	r, err := bucket.Object(configJSON).NewReader(ctx)
	if err == storage.ErrObjectNotExist {
		// ignore
	} else if err != nil {
		return nil, errors.Wrap(err, "read config")
	} else if err := json.NewDecoder(r).Decode(&config); err != nil {
		return nil, errors.Wrap(err, "decode config")
	}
	return &config, nil
}

type Replays []Replay

type Replay struct {
	ID          int        `json:"id"`
	Filename    string     `json:"filename"`
	Size        int        `json:"size"`
	GameType    string     `json:"game_type"`
	GameDate    hotsTime   `json:"game_date"`
	GameMap     string     `json:"game_map"`
	GameLength  int        `json:"game_length"`
	GameVersion string     `json:"game_version"`
	Fingerprint string     `json:"fingerprint"`
	Region      int        `json:"region"`
	Processed   bool       `json:"processed"`
	URL         string     `json:"url"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
	Bans        [][]string `json:"bans"`
	Players     []struct {
		Hero      string `json:"hero"`
		HeroLevel int    `json:"hero_level"`
		Team      int    `json:"team"`
		Winner    bool   `json:"winner"`
		BlizzID   int    `json:"blizz_id"`
		Party     int    `json:"party"`
		Silenced  bool   `json:"silenced"`
		Battletag string `json:"battletag"`
		Talents   struct {
			Num1  string `json:"1"`
			Num4  string `json:"4"`
			Num7  string `json:"7"`
			Num10 string `json:"10"`
			Num13 string `json:"13"`
			Num16 string `json:"16"`
			Num20 string `json:"20"`
		} `json:"talents"`
		Score struct {
			Level                  int         `json:"level"`
			Kills                  int         `json:"kills"`
			Assists                int         `json:"assists"`
			Takedowns              int         `json:"takedowns"`
			Deaths                 int         `json:"deaths"`
			HighestKillStreak      int         `json:"highest_kill_streak"`
			HeroDamage             int         `json:"hero_damage"`
			SiegeDamage            int         `json:"siege_damage"`
			StructureDamage        int         `json:"structure_damage"`
			MinionDamage           int         `json:"minion_damage"`
			CreepDamage            int         `json:"creep_damage"`
			SummonDamage           int         `json:"summon_damage"`
			TimeCcEnemyHeroes      int         `json:"time_cc_enemy_heroes"`
			Healing                int         `json:"healing"`
			SelfHealing            int         `json:"self_healing"`
			DamageTaken            interface{} `json:"damage_taken"`
			ExperienceContribution int         `json:"experience_contribution"`
			TownKills              int         `json:"town_kills"`
			TimeSpentDead          int         `json:"time_spent_dead"`
			MercCampCaptures       int         `json:"merc_camp_captures"`
			WatchTowerCaptures     int         `json:"watch_tower_captures"`
			MetaExperience         int         `json:"meta_experience"`
		} `json:"score"`
	} `json:"players"`
}

type hotsTime time.Time

func (h *hotsTime) UnmarshalText(text []byte) error {
	const format = "2006-01-02 15:04:05"
	t, err := time.Parse(format, string(text))
	if err != nil {
		return err
	}
	*h = hotsTime(t)
	return nil
}
