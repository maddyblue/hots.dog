package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/nfnt/resize"
	"github.com/pkg/errors"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

var (
	flagInit      = flag.Bool("init", false, "drop database before starting")
	flagAddr      = flag.String("addr", ":4001", "address to serve; HTTP redirect address if -autocert is set")
	flagAutocert  = flag.String("autocert", "", "domain name to autocert")
	flagAcmedir   = flag.String("acmedir", "", "optional acme directory; can be used to configure dev letsencrypt")
	flagCockroach = flag.String("cockroach", "postgresql://root@localhost:26257/?sslmode=disable", "cockroach connection URL")
	flagUpdate    = flag.Bool("update", false, "run hotsapi updater")
	flagElo       = flag.Bool("elo", false, "run elo update")
	initDB        = false
)

func main() {
	flag.Parse()

	if fromEnv := os.Getenv("ADDR"); fromEnv != "" {
		*flagAddr = fromEnv
	}
	if fromEnv := os.Getenv("COCKROACH"); fromEnv != "" {
		*flagCockroach = fromEnv
	}
	if fromEnv := os.Getenv("AUTOCERT"); fromEnv != "" {
		*flagAutocert = fromEnv
	}
	if fromEnv := os.Getenv("ACMEDIR"); fromEnv != "" {
		*flagAcmedir = fromEnv
	}

	const dbName = "hots"
	dbURL, err := url.Parse(*flagCockroach)
	if err != nil {
		log.Fatal(err)
	}
	dbURL.Path = dbName

	if *flagElo {
		if err := elo(dbURL.String()); err != nil {
			log.Fatalf("%+v", err)
		}
		return
	}

	db := mustInitDB(dbURL.String())
	defer db.Close()

	h := &hotsContext{
		db: db,
		x:  sqlx.NewDb(db, "postgres"),
	}

	if *flagUpdate {
		if err := h.update(); err != nil {
			log.Fatal(err)
		}
		return
	}

	h.mu.cache = make(map[string]cache)

	cacheTime := time.Minute * 10

	if *flagInit {
		cacheTime = -time.Second
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

	// Loop to allow for the server to start up.
	for i := 0; ; i++ {
		_, err := db.Exec(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s`, dbName))
		if err == nil {
			break
		}
		if i > 10 {
			panic(err)
		}
		fmt.Println("waiting:", err)
		time.Sleep(time.Second)
	}
	mustMigrate(db)
	if err := generateHeroes(db); err != nil {
		panic(err)
	}
	if err := h.updateInit(context.Background()); err != nil {
		panic(err)
	}

	wrap := func(f func(context.Context, *http.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
			url := r.URL.String()
			start := time.Now()
			useGzip := strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
			if useGzip {
				w.Header().Add("Content-Encoding", "gzip")
			}
			canCache := r.Method == "GET"
			defer func() { log.Printf("%s: %s", url, time.Since(start)) }()
			if canCache {
				h.mu.RLock()
				c, ok := h.mu.cache[url]
				h.mu.RUnlock()
				if ok && c.until > start.Unix() {
					w.Header().Add("Content-Type", "application/json")
					if useGzip {
						w.Write(c.gzip)
					} else {
						w.Write(c.data)
					}
					return
				}
				var data, gz []byte
				var until time.Time
				if err := h.db.QueryRowContext(ctx,
					"SELECT data, gzip, until FROM cache WHERE id = $1 AND until > now()",
					url,
				).Scan(&data, &gz, &until); err == nil {
					w.Header().Add("Content-Type", "application/json")
					if useGzip {
						w.Write(gz)
					} else {
						w.Write(data)
					}
					h.mu.Lock()
					h.mu.cache[url] = cache{
						until: until.Unix(),
						data:  data,
						gzip:  gz,
					}
					h.mu.Unlock()
					return
				}
			}

			res, err := f(ctx, r)
			if err != nil {
				log.Printf("%s: %+v", url, err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if res == nil {
				return
			}
			w.Header().Add("Content-Type", "application/json")
			b, err := json.Marshal(res)
			if err != nil {
				log.Printf("%s: JSON encoding error: %v", url, err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			var gz bytes.Buffer
			gzw, _ := gzip.NewWriterLevel(&gz, gzip.BestCompression)
			if _, err := gzw.Write(b); err != nil {
				log.Printf("%s: gzip write error: %v", url, err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := gzw.Close(); err != nil {
				log.Printf("%s: gzip close error: %v", url, err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if useGzip {
				w.Write(gz.Bytes())
			} else {
				w.Write(b)
			}
			if canCache {
				until := start.Add(cacheTime)
				h.mu.Lock()
				h.mu.cache[url] = cache{
					until: until.Unix(),
					data:  b,
					gzip:  gz.Bytes(),
				}
				h.mu.Unlock()
				go func() {
					if _, err := h.db.Exec("UPSERT INTO cache (id, until, data, gzip) VALUES ($1, $2, $3, $4)",
						url,
						until,
						b,
						gz.Bytes(),
					); err != nil {
						log.Printf("update cache table: %s: %v", url, err)
					}
				}()
			}
		}
	}

	http.Handle("/api/init", wrap(h.Init))
	http.Handle("/api/get-winrates", wrap(h.GetWinrates))
	//http.Handle("/api/get-build-winrates/:hero", wrap(h.GetBuildWinrates))
	http.Handle("/api/clear-cache", wrap(h.ClearCache))
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	if *flagInit && initDB {
		go mustInitDevData(*flagAddr, db)
	}

	go func() {
		for range time.Tick(time.Hour) {
			if err := h.updateInit(context.Background()); err != nil {
				log.Printf("could not update init data: %+v", err)
			}
		}
	}()

	if *flagAutocert != "" {
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				u := r.URL
				u.Scheme = "https"
				u.Host = *flagAutocert
				log.Printf("https redirect to %s", u)
				// 301
				http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
			})
			log.Fatal(http.ListenAndServe(*flagAddr, mux))
		}()
		log.Printf("AUTOCERT on: %s", *flagAutocert)
		if *flagAcmedir != "" {
			log.Printf("ACMEDIR: %s", *flagAcmedir)
		}
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(*flagAutocert),
			Cache:      dbCache{db},
			Client: &acme.Client{
				DirectoryURL: *flagAcmedir,
			},
		}
		s := &http.Server{
			Addr:      ":https",
			TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
		}
		log.Fatal(s.ListenAndServeTLS("", ""))
	} else {
		log.Printf("HTTP listen on addr: %s", *flagAddr)
		log.Fatal(http.ListenAndServe(*flagAddr, nil))
	}
}

const autocertPrefix = "autocert-"

type dbCache struct {
	*sql.DB
}

func (db dbCache) Get(ctx context.Context, key string) ([]byte, error) {
	var data []byte
	if err := db.QueryRowContext(ctx, "SELECT s FROM config WHERE key = $1", autocertPrefix+key).Scan(&data); err == sql.ErrNoRows {
		return nil, autocert.ErrCacheMiss
	} else if err != nil {
		return nil, err
	}
	return data, nil
}

func (db dbCache) Put(ctx context.Context, key string, data []byte) error {
	_, err := db.ExecContext(ctx, "UPSERT INTO config (key, s) VALUES ($1, $2)", autocertPrefix+key, data)
	return err
}

func (db dbCache) Delete(ctx context.Context, key string) error {
	_, err := db.ExecContext(ctx, "DELETE FROM config WHERE key = $1", autocertPrefix+key)
	return err
}

type hotsContext struct {
	db   *sql.DB
	x    *sqlx.DB
	init initData

	mu struct {
		sync.RWMutex
		cache map[string]cache
	}
}

type cache struct {
	until int64
	data  []byte
	gzip  []byte
}

type initData struct {
	Modes      map[Mode]string
	Builds     []Build
	Maps       []string
	Heroes     []Hero
	buildStats map[string]map[Mode]Stats
}

func (h *hotsContext) Init(ctx context.Context, _ *http.Request) (interface{}, error) {
	return h.init, nil
}

func (h *hotsContext) updateInit(ctx context.Context) error {
	maps, heroes, err := h.getNames(ctx)
	if err != nil {
		return err
	}
	builds, err := h.getBuilds(ctx)
	if err != nil {
		return err
	}
	bs := make(map[string]map[Mode]Stats)
	{
		rows, err := h.db.QueryContext(ctx, "SELECT * FROM skillstats")
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var build string
			var mode Mode
			var data []byte
			if err := rows.Scan(&build, &mode, &data); err != nil {
				return err
			}
			var s Stats
			if err := json.Unmarshal(data, &s); err != nil {
				return err
			}
			if _, ok := bs[build]; !ok {
				bs[build] = make(map[Mode]Stats)
			}
			bs[build][mode] = s
		}
		if err := rows.Err(); err != nil {
			return err
		}
	}
	h.init = initData{
		Modes:      modeNames,
		Builds:     builds,
		Maps:       maps,
		Heroes:     heroes,
		buildStats: bs,
	}
	return nil
}

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

func httpGet(ctx context.Context, url string) ([]byte, error) {
	start := time.Now()
	defer func() {
		log.Printf("GET %s took %s", url, time.Since(start))
	}()
	for i := 0; i < 10; i++ {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, errors.Wrap(err, url)
		}
		req = req.WithContext(ctx)
		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, errors.Wrap(err, url)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, errors.Wrap(err, url)
		}
		// Too many requests, backoff a bit.
		if resp.StatusCode == 429 {
			time.Sleep(time.Second * 5)
			continue
		}
		if resp.StatusCode != 200 {
			return nil, errors.Errorf("%s: %s: %s", url, resp.Status, b)
		}
		return b, nil
	}
	return nil, errors.Errorf("%s: too many retries", url)
}

func (h *hotsContext) ClearCache(ctx context.Context, _ *http.Request) (interface{}, error) {
	if err := h.updateInit(ctx); err != nil {
		return nil, err
	}
	h.mu.cache = make(map[string]cache)
	return nil, nil
}

// txn executes a transaction. If the database returns a retryable error,
// fn is re-invoked. fn should not call Commit or Rollback.
func (h *hotsContext) txn(ctx context.Context, fn func(txn *sqlx.Tx) error) error {
	for {
		txn, err := h.x.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
		err = fn(txn)
		if err == nil {
			return txn.Commit()
		}
		txn.Rollback()

		// Check if this was a retryable error.
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "40001" {
			continue
		}
		return err
	}
}

/*
func (h *hotsContext) GetBuildWinrates(ctx context.Context, r *http.Request) (interface{}, error) {
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
	if prevBuild, ok := h.getBuildBefore(args["build"]); ok {
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
*/

func (h *hotsContext) GetWinrates(ctx context.Context, r *http.Request) (interface{}, error) {
	args := map[string]string{
		"build":      r.FormValue("build"),
		"herolevel":  r.FormValue("herolevel"),
		"map":        r.FormValue("map"),
		"mode":       r.FormValue("mode"),
		"skill_low":  r.FormValue("skill_low"),
		"skill_high": r.FormValue("skill_high"),
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
	if prevBuild, ok := h.getBuildBefore(args["build"]); ok {
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
	hl := args["herolevel"]
	if hl == "" {
		hl = "5"
	}
	wheres = append(wheres, fmt.Sprintf("hero_level >= $%d", len(params)+1))
	params = append(params, hl)
	if m, sl, sh := args["mode"], args["skill_low"], args["skill_high"]; m != "" && (sl != "" || sh != "") {
		i, err := strconv.Atoi(m)
		if err != nil {
			return nil, err
		}
		modes, ok := h.init.buildStats[args["build"]]
		if !ok {
			return nil, errors.Errorf("unknown build: %s", args["build"])
		}
		quantiles := modes[Mode(i)].Quantile
		if sl != "" {
			wheres = append(wheres, fmt.Sprintf("skill >= $%d", len(params)+1))
			i, err = strconv.Atoi(sl)
			if err != nil {
				return nil, err
			}
			params = append(params, quantiles[i])
		}
		if sh != "" {
			wheres = append(wheres, fmt.Sprintf("skill <= $%d", len(params)+1))
			i, err = strconv.Atoi(sh)
			if err != nil {
				return nil, err
			}
			params = append(params, quantiles[i])
		}
	}

	buf := bytes.NewBufferString("SELECT COUNT(*) count, hero, winner FROM players")
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

func (h *hotsContext) getBuildBefore(id string) (build string, ok bool) {
	for i, b := range h.init.Builds {
		if b.ID == id {
			if len(h.init.Builds) == i+1 {
				return "", false
			}
			return h.init.Builds[i+1].ID, true
		}
	}
	return "", false
}

func (h *hotsContext) getBuilds(ctx context.Context) (builds []Build, err error) {
	err = h.x.SelectContext(ctx, &builds, "SELECT * from builds ORDER BY id DESC")
	return builds, err
}

func (h *hotsContext) getNames(ctx context.Context) (maps []string, heroes []Hero, err error) {
	if err := h.x.SelectContext(ctx, &maps, "SELECT * FROM maps ORDER BY name"); err != nil {
		return nil, nil, err
	}
	if err := h.x.SelectContext(ctx, &heroes, "SELECT name, roles, icon FROM heroes"); err != nil {
		return nil, nil, err
	}
	return
}

type Hero struct {
	Name  string
	Roles string
	Icon  string
}

func generateHeroes(db *sql.DB) error {
	var count int
	if err := db.QueryRow("SELECT count(*) FROM heroes").Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		return doGenerateHeroes(db)
	}
	return nil
}

func doGenerateHeroes(db *sql.DB) error {
	resp, err := http.Get("http://us.battle.net/heroes/en/heroes/")
	if err != nil {
		return err
	}
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	matches := regexp.MustCompile("window\\.heroes = (.*);\n//]]>").FindStringSubmatch(string(html))
	if len(matches) != 2 {
		return errors.New("no match")
	}

	type Hero struct {
		icon string

		Name string `json:"name"`
		Slug string `json:"slug"`
		Role struct {
			Slug string `json:"slug"`
		} `json:"role"`
		RoleSecondary struct {
			Slug string `json:"slug"`
		} `json:"roleSecondary"`
	}
	var heroes []*Hero
	if err := json.Unmarshal([]byte(matches[1]), &heroes); err != nil {
		return err
	}
	g, ctx := errgroup.WithContext(context.Background())
	_ = ctx
	thumbnailRE := regexp.MustCompile("http://media.blizzard.com/heroes/(.*)/skins/thumbnails/(.*).jpg")
	for _, h := range heroes {
		h := h
		g.Go(func() error {
			url := fmt.Sprintf("http://us.battle.net/heroes/en/heroes/%s/", h.Slug)
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			html, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			resp.Body.Close()

			match := thumbnailRE.FindString(string(html))
			if match == "" {
				return errors.Errorf("could not find thumbnail for %s", h.Name)
			}

			resp, err = http.Get(match)
			if err != nil {
				return err
			}
			icon, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			if h.icon, err = toDataURI(icon); err != nil {
				return errors.Wrap(err, h.Name)
			}
			resp.Body.Close()
			roles := []string{h.Role.Slug}
			if h.RoleSecondary.Slug != "" {
				roles = append(roles, h.RoleSecondary.Slug)
			}
			_, err = db.Exec("UPSERT INTO heroes (slug, name, roles, icon) VALUES ($1, $2, $3, $4)",
				h.Slug,
				h.Name,
				strings.Join(roles, ","),
				h.icon,
			)
			return err
		})
	}
	return g.Wait()
}

func toDataURI(b []byte) (string, error) {
	img, format, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	img = resize.Resize(40, 40, img, resize.Bicubic)
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return "", err
	}
	return fmt.Sprintf("data:image/%s;base64,%s",
		format,
		base64.StdEncoding.EncodeToString(buf.Bytes()),
	), nil
}
