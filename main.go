package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

var (
	flagInit = flag.Bool("init", false, "drop database before starting")
	initDB   = true
)

func main() {
	flag.Parse()

	const dbName = "hots"
	db := mustInitDB(fmt.Sprintf("postgresql://root@localhost:26257/%s?sslmode=disable", dbName))
	defer db.Close()

	h := &hotsContext{
		db: db,
		x:  sqlx.NewDb(db, "postgres"),
	}

	if *flagInit {
		// Don't exit on panic; prevents modd from spinning.
		defer func() {
			return
			if r := recover(); r != nil {
				fmt.Printf("%+v", r)
				select {}
			}
		}()

		if initDB {
			time.Sleep(time.Second * 2)
			if _, err := db.Exec(fmt.Sprintf("drop database if exists %s cascade; create database %[1]s", dbName)); err != nil {
				panic(err)
			}
		}
	}

	mustMigrate(db)

	wrap := func(f func(context.Context, *http.Request, httprouter.Params) (interface{}, error)) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			log.Printf("request: %s", r.URL)
			ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
			res, err := f(ctx, r, ps)
			if err != nil {
				log.Printf("%+v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if res == nil {
				return
			}
			w.Header().Add("Content-Type", "application/json")
			enc := json.NewEncoder(w)
			enc.SetIndent("", "\t")
			if err := enc.Encode(res); err != nil {
				log.Print(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}

	router := httprouter.New()
	router.GET("/api/init", wrap(h.Init))
	router.GET("/api/get-winrates", wrap(h.GetWinrates))
	router.GET("/api/get-build-winrates/:hero", wrap(h.GetBuildWinrates))
	router.GET("/api/next-block", wrap(h.NextBlock))

	const addr = "localhost:4001"

	if *flagInit && initDB {
		go mustInitDevData(addr, db)
	}

	log.Fatal(http.ListenAndServe(addr, router))
}

type hotsContext struct {
	db *sql.DB
	x  *sqlx.DB
}

func (h *hotsContext) Init(ctx context.Context, _ *http.Request, _ httprouter.Params) (interface{}, error) {
	maps, heroes, err := h.getNames(ctx)
	if err != nil {
		return nil, err
	}
	builds, err := h.getBuilds(ctx)
	if err != nil {
		return nil, err
	}
	return struct {
		Modes  map[Mode]string
		Builds []Build
		Maps   []string
		Heroes []string
	}{
		Modes:  modeNames,
		Builds: builds,
		Maps:   maps,
		Heroes: heroes,
	}, err
}

func httpGet(ctx context.Context, url string) (*http.Response, error) {
	start := time.Now()
	defer func() {
		log.Printf("GET %s took %s", url, time.Since(start))
	}()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	return http.DefaultClient.Do(req)
}

func (h *hotsContext) NextBlock(ctx context.Context, _ *http.Request, _ httprouter.Params) (interface{}, error) {
	return nil, h.nextBlock(ctx)
}

func (h *hotsContext) nextBlock(ctx context.Context) error {
	var lastID int
	const configLastID = "last_id"
	if err := h.x.GetContext(ctx, &lastID, `
		SELECT i
		FROM config
		WHERE key = $1
	`, configLastID); err == sql.ErrNoRows {
		// lastID = 0
	} else if err != nil {
		return err
	}

	var lastGameID int
	if err := h.x.GetContext(ctx, &lastGameID, `
	SELECT id
		FROM games
		ORDER BY id DESC
		LIMIT 1
	`); err == nil {
		lastID = lastGameID
	}

	// Since min_id is the first result returned, we need to increment what was
	// returned.
	lastID++

	resp, err := httpGet(ctx, fmt.Sprintf("http://hotsapi.net/api/v1/replays?min_id=%d", lastID))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var replays Replays
	if err := json.NewDecoder(resp.Body).Decode(&replays); err != nil {
		return errors.Wrap(err, "decoding replays")
	}
	ch := make(chan *Replay, 5)
	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		for r := range ch {
			if err := h.getReplay(gCtx, r); err != nil {
				return err
			}
			lastID = r.ID
		}
		return nil
	})
	group.Go(func() error {
		defer close(ch)
		for _, r := range replays {
			resp, err := httpGet(ctx, fmt.Sprintf("http://hotsapi.net/api/v1/replays/%d", r.ID))
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			var replay Replay
			if err := json.NewDecoder(resp.Body).Decode(&replay); err != nil {
				return errors.Wrap(err, "decoding replay")
			}
			select {
			case ch <- &replay:
			case <-gCtx.Done():
				return ctx.Err()
			}
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		return err
	}
	_, err = h.db.ExecContext(ctx, `UPSERT INTO config (key, i) VALUES ($1, $2)`, configLastID, lastID)
	return err
}
}

func (h *hotsContext) getReplay(ctx context.Context, r *Replay) error {
	mode, ok := gameModes[r.GameType]
	if !ok {
		return errors.Errorf("unknown game type: %s", r.GameType)
	}
	if _, err := h.addBuild(ctx, r); err != nil {
		return err
	}
	if _, err := h.db.Exec(`UPSERT INTO maps (name) VALUES ($1) RETURNING NOTHING`, r.GameMap); err != nil {
		return err
	}
	for _, p := range r.Players {
		if _, err := h.db.Exec(`UPSERT INTO heroes (name) VALUES ($1) RETURNING NOTHING`, p.Hero); err != nil {
			return err
		}
	}
	txn, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer txn.Rollback()
	if _, err := txn.Exec(fmt.Sprintf(
		`INSERT INTO games (id, url, time, length, build, map, mode, region) VALUES %s RETURNING NOTHING`,
		makeValues(8)),
		r.ID, r.URL, time.Time(r.GameDate), r.GameLength, r.GameVersion, r.GameMap, mode, r.Region,
	); err != nil {
		return errors.Wrap(err, "insert into games")
	}
	for _, p := range r.Players {
		if _, err := txn.Exec(`UPSERT INTO battletags (id, name) VALUES ($1, $2) RETURNING NOTHING`, p.BlizzID, p.Battletag); err != nil {
			return err
		}
		if _, err := txn.Exec(fmt.Sprintf(`INSERT INTO players (
			game,
			blizzid,
			hero,
			hero_level,
			team,
			winner,
			build,
			length,
			map,
			mode,
			region
		) VALUES %s RETURNING NOTHING`, makeValues(11)),
			r.ID,
			p.BlizzID,
			p.Hero,
			p.HeroLevel,
			p.Team,
			p.Winner,
			r.GameVersion,
			r.GameLength,
			r.GameMap,
			mode,
			r.Region,
		); err != nil {
			return err
		}
	}
	return txn.Commit()
}

type Replays []Replay

type Replay struct {
	ID          int      `json:"id"`
	Filename    string   `json:"filename"`
	Size        int      `json:"size"`
	GameType    string   `json:"game_type"`
	GameDate    hotsTime `json:"game_date"`
	GameLength  int      `json:"game_length"`
	GameMap     string   `json:"game_map"`
	GameVersion string   `json:"game_version"`
	Region      int      `json:"region"`
	Fingerprint string   `json:"fingerprint"`
	URL         string   `json:"url"`
	Players     []struct {
		Battletag string `json:"battletag"`
		Hero      string `json:"hero"`
		HeroLevel int    `json:"hero_level"`
		Team      int    `json:"team"`
		Winner    bool   `json:"winner"`
		Region    int    `json:"region"`
		BlizzID   int    `json:"blizz_id"`
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

func (h *hotsContext) GetBuildWinrates(ctx context.Context, r *http.Request, ps httprouter.Params) (interface{}, error) {
	args := map[string]string{
		"build": r.FormValue("build"),
		"hero":  ps.ByName("hero"),
		"map":   r.FormValue("map"),
		"mode":  r.FormValue("mode"),
	}
	wrs, err := h.getBuildWinrates(ctx, args)
	if err != nil {
		return nil, errors.Wrap(err, "getBuildWinrates")
	}
	ret := struct {
		Current  map[int]map[string]Total
		Previous map[int]map[string]Total
	}{
		Current: wrs,
	}
	if prevBuild, ok := h.getBuildBefore(ctx, args["build"]); ok {
		args["build"] = prevBuild
		prevWrs, err := h.getBuildWinrates(ctx, args)
		if err != nil {
			return nil, errors.Wrapf(err, "fetch previous build: %v", prevBuild)
		}
		ret.Previous = prevWrs
	}
	return ret, nil
}

func (h *hotsContext) getBuildWinrates(ctx context.Context, args map[string]string) (map[int]map[string]Total, error) {
	if args["build"] == "" {
		return nil, errors.New("build required")
	}
	if args["hero"] == "" {
		return nil, errors.New("hero required")
	}

	groups := []string{"name", "tier", "winner"}
	var wheres []string
	var params []interface{}
	for _, key := range []string{"build", "hero", "map", "mode"} {
		v := args[key]
		if v == "" {
			continue
		}
		wheres = append(wheres, fmt.Sprintf("%s = $%d", key, len(params)+1))
		groups = append(groups, key)
		params = append(params, v)
	}

	buf := bytes.NewBufferString(`SELECT COUNT(*) count, "name", winner, tier`)
	buf.WriteString(" FROM talents")
	if len(wheres) > 0 {
		fmt.Fprintf(buf, " WHERE %s", strings.Join(wheres, " AND "))
	}
	buf.WriteString(" GROUP BY ")
	buf.WriteString(strings.Join(groups, ", "))
	query := buf.String()

	tally := make(map[int]map[string]Total)
	for i := 1; i <= 7; i++ {
		tally[i] = make(map[string]Total)
	}

	var winrates []struct {
		Count  int
		Name   string
		Tier   int
		Winner bool
	}
	if err := h.x.Select(&winrates, query, params...); err != nil {
		return nil, errors.Wrap(err, "select")
	}
	for _, wr := range winrates {
		t := tally[wr.Tier][wr.Name]
		if wr.Winner {
			t.Wins += wr.Count
		} else {
			t.Losses += wr.Count
		}
		tally[wr.Tier][wr.Name] = t
	}
	return tally, nil
}

func (h *hotsContext) GetWinrates(ctx context.Context, r *http.Request, _ httprouter.Params) (interface{}, error) {
	args := map[string]string{
		"build": r.FormValue("build"),
		"map":   r.FormValue("map"),
		"mode":  r.FormValue("mode"),
	}
	wrs, err := h.getWinrates(ctx, args)
	if err != nil {
		return nil, errors.Wrap(err, "getWinrates")
	}
	ret := struct {
		Current  map[string]Total
		Previous map[string]Total
	}{
		Current: wrs,
	}
	if prevBuild, ok := h.getBuildBefore(ctx, args["build"]); ok {
		args["build"] = prevBuild
		prevWrs, err := h.getWinrates(ctx, args)
		if err != nil {
			return nil, errors.Wrapf(err, "getWinrates previous: %v", prevBuild)
		}
		ret.Previous = prevWrs
	}
	return ret, nil
}

func (h *hotsContext) getWinrates(ctx context.Context, args map[string]string) (map[string]Total, error) {
	if args["build"] == "" {
		return nil, errors.New("build required")
	}

	groups := []string{"hero", "winner"}
	var wheres []string
	var params []interface{}
	for _, key := range []string{"build", "map", "mode"} {
		v := args[key]
		if v == "" {
			continue
		}
		wheres = append(wheres, fmt.Sprintf("%s = $%d", key, len(params)+1))
		groups = append(groups, key)
		params = append(params, v)
	}

	buf := bytes.NewBufferString("SELECT COUNT(*) count, hero, winner")
	buf.WriteString(" FROM players")
	if len(wheres) > 0 {
		fmt.Fprintf(buf, " WHERE %s", strings.Join(wheres, " AND "))
	}
	buf.WriteString(" GROUP BY ")
	buf.WriteString(strings.Join(groups, ", "))
	query := buf.String()

	tally := make(map[string]Total)

	var winrates []struct {
		Hero   string
		Count  int
		Winner bool
	}
	fmt.Println("\n", query, "\n")
	if err := h.x.Select(&winrates, query, params...); err != nil {
		return nil, errors.Wrap(err, "select from players")
	}
	for _, wr := range winrates {
		t := tally[wr.Hero]
		if wr.Winner {
			t.Wins += wr.Count
		} else {
			t.Losses += wr.Count
		}
		tally[wr.Hero] = t
	}
	return tally, nil
}

type Total struct {
	Wins, Losses int
}

func (h *hotsContext) getBuildBefore(ctx context.Context, id string) (build string, ok bool) {
	err := h.x.GetContext(ctx, &build, `
		SELECT id
		FROM builds
		WHERE id < $1
		ORDER BY id DESC
		LIMIT 1
	`, id)
	return build, err == nil
}

func (h *hotsContext) getBuilds(ctx context.Context) (builds []Build, err error) {
	err = h.x.SelectContext(ctx, &builds, "SELECT * from builds ORDER BY id DESC")
	return builds, err
}

func (h *hotsContext) getNames(ctx context.Context) (maps, heroes []string, err error) {
	if err := h.x.SelectContext(ctx, &maps, "SELECT * FROM maps ORDER BY name"); err != nil {
		return nil, nil, err
	}
	if err := h.x.SelectContext(ctx, &heroes, "SELECT * FROM heroes ORDER BY name"); err != nil {
		return nil, nil, err
	}
	return
}

func (h *hotsContext) getNamesAsMaps(ctx context.Context) (maps, heroes map[string]bool, err error) {
	mapNames, heroNames, err := h.getNames(ctx)
	maps = make(map[string]bool)
	for _, n := range mapNames {
		maps[n] = true
	}
	heroes = make(map[string]bool)
	for _, n := range heroNames {
		heroes[n] = true
	}
	return maps, heroes, err
}

func (h *hotsContext) addBuild(ctx context.Context, replay *Replay) (Build, error) {
	var build Build
	tx, err := h.x.BeginTx(ctx, nil)
	if err != nil {
		return build, err
	}

	id := replay.GameVersion
	start := time.Time(replay.GameDate)
	if err = h.x.Get(&build, "SELECT * FROM builds WHERE id = $1", id); err == sql.ErrNoRows {
		build.ID = id
		build.Start = start
		build.Finish = start
		_, err = tx.Exec("INSERT INTO builds (id, start, finish) VALUES ($1, $2, $3)", build.ID, build.Start, build.Finish)
	} else if err == nil {
		buildChanged := false
		if start.Before(build.Start) {
			build.Start = start
			buildChanged = true
		}
		if start.After(build.Finish) {
			build.Finish = start
			buildChanged = true
		}
		if buildChanged {
			_, err = tx.Exec("UPDATE builds SET start = $1, finish = $2 WHERE id = $3", build.Start, build.Finish, id)
		}
	}
	if err != nil {
		tx.Rollback()
		return build, err
	}
	return build, tx.Commit()
}
