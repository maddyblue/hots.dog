package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	servertiming "github.com/mitchellh/go-server-timing"
	"github.com/pkg/errors"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/sync/errgroup"
)

var (
	flagInit      = flag.Bool("init", false, "drop database before starting")
	flagAddr      = flag.String("addr", ":4001", "address to serve; HTTP redirect address if -autocert is set")
	flagAutocert  = flag.String("autocert", "", "domain name to autocert")
	flagAcmedir   = flag.String("acmedir", "", "optional acme directory; can be used to configure dev letsencrypt")
	flagCockroach = flag.String("cockroach", "postgresql://root@localhost:26257/hots?sslmode=disable", "cockroach connection URL")
	flagElo       = flag.Bool("elo", false, "run elo update")
	flagMigrate   = flag.Bool("migrate", false, "run migration")
	flagCron      = flag.Bool("cron", false, "run cronjob")
	flagUpdateNew = flag.String("updatenew", "", "run new update to specified gs bucket")
	flagImport    = flag.String("import", "csv2.hots.dog", "import from bucket")
	flagImportNum = flag.Int("importnum", -1, "max id to import; set to 0 for first block only")
	flagUpdateDB  = flag.Bool("updatedb", false, "update db from import bucket")
	initDB        = false

	popularGameLimit    = 10
	leaderboardMinGames = 50
	daysOld             = 90
)

const dbName = "hots"

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

	if *flagUpdateNew != "" {
		if err := updateNew(*flagUpdateNew); err != nil {
			log.Fatalf("%+v", err)
		}
		return
	}

	dbURL, err := url.Parse(*flagCockroach)
	if err != nil {
		log.Fatal(err)
	}
	dbURL.Path = dbName

	db := mustInitDB(dbURL.String())
	defer db.Close()

	h := &hotsContext{
		dburl: dbURL.String(),
		db:    db,
		x:     sqlx.NewDb(db, "postgres"),
	}

	if *flagImportNum != -1 {
		mustMigrate(db)
		if err := h.Import(*flagImport, *flagImportNum); err != nil {
			log.Fatalf("%+v", err)
		}
		return
	}

	if *flagElo {
		if err := h.elo(); err != nil {
			log.Fatalf("%+v", err)
		}
		return
	}

	/*
	   The database cache has two timestamps: until and last_hit. last_hit is
	   the last time a user request hit the URL. until is the time after which
	   the request should be re-processed. A cron job routinely clears out
	   old cache entries whose last_hit field is older than some threshold
	   (2 days or so?). The same cron job also re-processes entries whose
	   until time has passed, and resets the until time for some small
	   threshold in the future (1 hour)?. Thus, the user code should never
	   consult any timestamps in the cache table.

	   When the in-memory cache is used, the table's last_hit column is not
	   updated. That field is only updated when the cache table is consulted
	   for a hit. This means that writes don't have to happen in most user
	   requests, and at worst the last_hit field will be out-of-date for
	   whatever the until length (1 hour) is.
	*/
	h.cacheTime = time.Hour
	if *flagInit {
		h.cacheTime = time.Second * 5
		popularGameLimit = 2
		leaderboardMinGames = 3
		daysOld = int(time.Since(time.Date(2017, time.April, 1, 0, 0, 0, 0, time.UTC)) / (time.Hour * 24))
	}

	if *flagCron {
		if err := h.cronLoop(); err != nil {
			log.Fatalf("%+v", err)
		}
		return
	}
	if *flagUpdateDB {
		if err := h.updateDB(); err != nil {
			log.Fatalf("%+v", err)
		}
		return
	}

	h.mu.cache = make(map[string]cache)

	if *flagInit {
		// Don't exit on panic; prevents modd from spinning.
		defer func() {
			return
			if r := recover(); r != nil {
				fmt.Printf("%+v\n", r)
				select {}
			}
		}()
		if initDB {
			time.Sleep(time.Second * 2)
			mustExec(db, fmt.Sprintf("drop database if exists %s cascade; create database %[1]s", dbName))
		}
	}

	if *flagMigrate || *flagInit {
		mustMigrate(db)
		if initDB {
			if err := h.Import(*flagImport, *flagImportNum); err != nil {
				log.Fatalf("%+v", err)
			}
		}
		if err := h.syncConfig(*flagImport); err != nil {
			log.Fatalf("%+v", err)
		}
		if !*flagInit {
			return
		}
		h.ClearCache(nil, nil)
	}

	if err := h.updateInit(context.Background()); err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	enableCache := !*flagInit

	wrap := func(f func(context.Context, *http.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), time.Second*60)
			defer cancel()
			var sh servertiming.Header
			ctx = servertiming.NewContext(ctx, &sh)
			if v, err := url.ParseQuery(r.URL.RawQuery); err == nil {
				r.URL.RawQuery = v.Encode()
			}
			url := r.URL.String()
			start := time.Now()
			defer func() { fmt.Printf("%s: %s\n", url, time.Since(start)) }()
			if enableCache && h.CheckCache(ctx, start, w, r, r.URL.Path, url) {
				return
			}
			tm := servertiming.FromContext(ctx).NewMetric("req").Start()
			res, err := f(ctx, r)
			tm.Stop()
			if len(sh.Metrics) > 0 {
				if len(sh.Metrics) > 10 {
					sh.Metrics = sh.Metrics[:10]
				}
				w.Header().Add(servertiming.HeaderKey, sh.String())
			}
			if err != nil {
				log.Printf("%s: %+v", url, err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data, gzip, err := resultToBytes(res)
			if err != nil {
				log.Printf("%s: %v", url, err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			writeDataGzip(w, r, data, gzip)
			if enableCache {
				go h.WriteCache(r.URL.Path, url, start, data, gzip)
			}
		}
	}

	mux := http.NewServeMux()

	mux.Handle("/api/init", wrap(h.Init))
	mux.Handle("/api/get-build-winrates", wrap(h.GetBuildWinrates))
	mux.Handle("/api/get-compare-hero", wrap(h.GetCompareHero))
	mux.Handle("/api/get-game-data", wrap(h.GetGameData))
	mux.Handle("/api/get-hero-data", wrap(h.GetHero))
	mux.Handle("/api/get-leaderboard", wrap(h.GetLeaderboard))
	mux.Handle("/api/get-player-by-name", wrap(h.GetPlayerName))
	mux.Handle("/api/get-player-games", wrap(h.GetPlayerGames))
	mux.Handle("/api/get-player-matchups", wrap(h.GetPlayerMatchups))
	mux.Handle("/api/get-player-profile", wrap(h.GetPlayerProfile))
	mux.Handle("/api/get-player-friends", wrap(h.GetPlayerFriends))
	mux.Handle("/api/get-winrates", wrap(h.GetWinrates))
	if *flagInit {
		mux.HandleFunc("/api/clear-cache", h.ClearCache)
	}

	fileServer := http.FileServer(http.Dir("static"))
	serveFiles := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/service-worker.js" {
			w.Header().Add("Cache-Control", "no-cache")
		} else if strings.HasPrefix(r.URL.Path, "/static") || strings.HasPrefix(r.URL.Path, "/img") {
			// Two week cache for probably static assets.
			w.Header().Add("Cache-Control", "max-age=1209600")
		} else {
			w.Header().Add("Cache-Control", "max-age=3600")
		}
		fileServer.ServeHTTP(w, r)
	}
	serveIndex := func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		serveFiles(w, r)
	}

	talents := make(map[string]bool)
	if err := filepath.Walk(filepath.Join("static", "img", "talent"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		talents[info.Name()] = true
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	mux.HandleFunc("/img/talent/", func(w http.ResponseWriter, r *http.Request) {
		base := filepath.Base(r.URL.Path)
		if !talents[base] {
			makeTalentImg(w, r)
			return
		}
		serveFiles(w, r)
	})

	mux.HandleFunc("/about/", serveIndex)
	mux.HandleFunc("/compare/", serveIndex)
	mux.HandleFunc("/games/", serveIndex)
	mux.HandleFunc("/heroes/", serveIndex)
	mux.HandleFunc("/leaderboard", serveIndex)
	mux.HandleFunc("/players/", serveIndex)
	mux.HandleFunc("/talents/", serveIndex)
	mux.HandleFunc("/", serveFiles)

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {})

	go func() {
		for range time.Tick(time.Minute * 10) {
			if err := h.updateInit(context.Background()); err != nil {
				panic(fmt.Sprintf("%+v", err))
			}
			// Clear the memory cache of old entries.
			h.mu.Lock()
			cutoff := time.Now().Add(-time.Hour).Unix()
			for s, c := range h.mu.cache {
				if c.until < cutoff {
					delete(h.mu.cache, s)
				}
			}
			h.mu.Unlock()
		}
	}()

	if *flagAutocert != "" {
		fmt.Println("AUTOCERT on:", *flagAutocert)
		if *flagAcmedir != "" {
			fmt.Println("ACMEDIR:", *flagAcmedir)
		}
		const cloudflareOrigin = "cloudflare-origin"
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(*flagAutocert),
			Cache:      dbCache{db},
		}
		tlsConfig := &tls.Config{GetCertificate: m.GetCertificate}
		if *flagAcmedir == cloudflareOrigin {
			var certfile, keyfile []byte
			if err := h.x.Get(&certfile, `SELECT s FROM config WHERE key = $1`, cloudflareOrigin+"-cert"); err != nil {
				log.Fatalf("could not get certfile origin: %v", err)
			}
			if err := h.x.Get(&keyfile, `SELECT s FROM config WHERE key = $1`, cloudflareOrigin+"-key"); err != nil {
				log.Fatalf("could not get keyfile origin: %v", err)
			}
			cert, err := tls.X509KeyPair(certfile, keyfile)
			if err != nil {
				log.Fatalf("cert: %v", err)
			}
			tlsConfig = &tls.Config{
				Certificates: []tls.Certificate{cert},
			}
		} else if *flagAcmedir != "" {
			m.Client = &acme.Client{
				DirectoryURL: *flagAcmedir,
			}
		}
		go func() {
			log.Fatal(http.ListenAndServe(":http", m.HTTPHandler(nil)))
		}()
		s := &http.Server{
			Addr:      ":https",
			TLSConfig: tlsConfig,
			Handler:   mux,
		}
		log.Fatal(s.ListenAndServeTLS("", ""))
	} else {
		fmt.Println("HTTP listen on addr:", *flagAddr)
		log.Fatal(http.ListenAndServe(*flagAddr, mux))
	}
}

var (
	enableDBCache = map[string]bool{
		"/api/get-build-winrates": true,
		"/api/get-compare-hero":   true,
		"/api/get-hero-data":      true,
		"/api/get-winrates":       true,
		"/api/get-leaderboard":    true,
	}
	enableMemCache = map[string]bool{
		"/api/init": true,
	}
)

func resultToBytes(res interface{}) (data, gzipped []byte, err error) {
	data, err = json.Marshal(res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "json marshal")
	}
	var gz bytes.Buffer
	gzw, _ := gzip.NewWriterLevel(&gz, gzip.BestCompression)
	if _, err := gzw.Write(data); err != nil {
		return nil, nil, errors.Wrap(err, "gzip")
	}
	if err := gzw.Close(); err != nil {
		return nil, nil, errors.Wrap(err, "gzip close")
	}
	return data, gz.Bytes(), nil
}

func (h *hotsContext) CheckCache(ctx context.Context, start time.Time, w http.ResponseWriter, r *http.Request, path, url string) (done bool) {
	if !enableMemCache[path] && !enableDBCache[path] {
		return false
	}
	h.mu.RLock()
	c, ok := h.mu.cache[url]
	h.mu.RUnlock()
	if ok && c.until > start.Unix() {
		writeDataGzip(w, r, c.data, c.gzip)
		return true
	}
	if !enableDBCache[path] {
		return false
	}
	var data, gz []byte
	if err := h.db.QueryRowContext(ctx,
		"SELECT data, gzip FROM cache WHERE id = $1",
		url,
	).Scan(&data, &gz); err == nil {
		writeDataGzip(w, r, data, gz)
		h.mu.Lock()
		h.mu.cache[url] = cache{
			until: start.Add(h.cacheTime).Unix(),
			data:  data,
			gzip:  gz,
		}
		h.mu.Unlock()
		// Don't block user return on db writes.
		go func() {
			if err := retry(func() error {
				_, err := h.db.Exec(`UPDATE cache SET last_hit = $1 WHERE id = $2`, start, url)
				return err
			}); err != nil {
				log.Printf("couldn't update cache last_hit: %s: %s", url, err)
			}
		}()
		return true
	}
	return false
}

func (h *hotsContext) WriteCache(path, url string, start time.Time, data, gzip []byte) {
	if !enableMemCache[path] && !enableDBCache[path] {
		return
	}
	until := start.Add(h.cacheTime)
	h.mu.Lock()
	h.mu.cache[url] = cache{
		until: until.Unix(),
		data:  data,
		gzip:  gzip,
	}
	h.mu.Unlock()
	if !enableDBCache[path] {
		return
	}
	if err := retry(func() error {
		_, err := h.db.Exec("UPSERT INTO cache (id, data, gzip, last_hit, until) VALUES ($1, $2, $3, $4, NULL)",
			url,
			data,
			gzip,
			start,
		)
		return err
	}); err != nil {
		log.Printf("update cache table: %s: %v", url, err)
	}
}

func writeDataGzip(w http.ResponseWriter, r *http.Request, data, gzip []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Cache-Control", "max-age=3600")
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Add("Content-Encoding", "gzip")
		w.Write(gzip)
	} else {
		w.Write(data)
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
	dburl     string
	db        *sql.DB
	x         *sqlx.DB
	cacheTime time.Duration

	mu struct {
		sync.RWMutex
		cache map[string]cache
		init  initData
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
	BuildStats map[string]map[Mode]Stats
	config     *groupConfig
	lookups    map[string]func(string) string
}

func (i initData) list(name, s string) []string {
	lookup := i.lookups[name]
	var res []string
	for _, v := range strings.Split(s[1:len(s)-1], ",") {
		res = append(res, lookup(v))
	}
	return res
}

func (h *hotsContext) Init(ctx context.Context, _ *http.Request) (interface{}, error) {
	return h.getInit(), nil
}

func (h *hotsContext) updateInit(ctx context.Context) error {
	var maps []byte
	if err := h.x.GetContext(ctx, &maps, "SELECT s FROM config WHERE key = $1", cacheConfig); err != nil {
		return errors.Wrap(err, "get config")
	}
	var c groupConfig
	if err := json.Unmarshal(maps, &c); err != nil {
		return err
	}
	c.readonly = true
	bsRows, err := h.db.QueryContext(ctx, "SELECT build, mode, data FROM skillstats")
	if err != nil {
		return err
	}
	defer bsRows.Close()
	h.mu.Lock()
	h.mu.init = initData{
		Modes:   modeNames,
		Heroes:  heroData,
		config:  &c,
		lookups: make(map[string]func(string) string),
	}
	for m := range c.Map["map"] {
		h.mu.init.Maps = append(h.mu.init.Maps, m)
	}
	sort.Strings(h.mu.init.Maps)
	for n, b := range c.Builds {
		h.mu.init.Builds = append(h.mu.init.Builds, Build{
			ID:     n,
			Start:  b.Start,
			Finish: b.End,
		})
	}
	sort.Slice(h.mu.init.Builds, func(i, j int) bool {
		return h.mu.init.Builds[i].ID > h.mu.init.Builds[j].ID
	})
	for group, m := range c.Map {
		lookup := make(map[string]string)
		for k, v := range m {
			lookup[v] = k
		}
		h.mu.init.lookups[group] = func(name string) string {
			return lookup[name]
		}
	}
	bs := make(map[string]map[Mode]Stats)
	{
		for bsRows.Next() {
			var build string
			var mode Mode
			var data []byte
			if err := bsRows.Scan(&build, &mode, &data); err != nil {
				return err
			}
			var s Stats
			if err := json.Unmarshal(data, &s); err != nil {
				return err
			}
			bid := h.mu.init.lookups["build"](build)
			if bid == "" {
				return errors.Errorf("build not found: %s", build)
			}
			if _, ok := bs[bid]; !ok {
				bs[bid] = make(map[Mode]Stats)
			}
			bs[bid][mode] = s
		}
		if err := bsRows.Err(); err != nil {
			return err
		}
	}
	h.mu.init.BuildStats = bs
	h.mu.Unlock()
	return nil
}

func (h *hotsContext) ClearCache(_ http.ResponseWriter, _ *http.Request) {
	h.mu.Lock()
	h.mu.cache = make(map[string]cache)
	if err := retry(func() error {
		_, err := h.db.Exec("DELETE FROM cache")
		return err
	}); err != nil {
		log.Println(err)
	}
	h.mu.Unlock()
	if err := h.updateInit(context.Background()); err != nil {
		log.Println(err)
	}
}

// txn executes a transaction. If the database returns a retryable error,
// fn is re-invoked. fn should not call Commit or Rollback.
func (h *hotsContext) txn(ctx context.Context, fn func(txn *sqlx.Tx) error) error {
	return retry(func() error {
		txn, err := h.x.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
		err = fn(txn)
		if err == nil {
			return txn.Commit()
		}
		txn.Rollback()
		return err
	})
}

// retry executes fn, but retries it if fn returns a retryable postgres error.
func retry(fn func() error) error {
	for count := 0; count < 10; count++ {
		err := fn()
		if err == nil {
			return nil
		}

		if retryable(err) {
			continue
		}
		return err
	}
	return errors.New("retry limit exhausted")
}

func retryable(err error) bool {
	err = errors.Cause(err)

	pqErr, ok := err.(*pq.Error)
	if ok && pqErr.Code == "40001" {
		return true
	}

	if strings.Contains(err.Error(), "connection reset by peer") {
		return true
	}

	return false
}

func (h *hotsContext) GetBuildWinrates(ctx context.Context, r *http.Request) (interface{}, error) {
	args := map[string]string{
		"build":      r.FormValue("build"),
		"hero":       r.FormValue("hero"),
		"herolevel":  r.FormValue("herolevel"),
		"map":        r.FormValue("map"),
		"mode":       r.FormValue("mode"),
		"skill_low":  r.FormValue("skill_low"),
		"skill_high": r.FormValue("skill_high"),
	}
	init := h.getInit()
	var res struct {
		Current       map[int]map[string]Total
		Previous      map[int]map[string]Total
		PopularBuilds []build
		WinningBuilds []build
		Talents       map[string]talentText
	}
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		res.Current, res.PopularBuilds, res.WinningBuilds, err = h.getBuildWinrates(ctx, init, args)
		return errors.Wrap(err, "getBuildWinrates")
	})
	g.Go(func() error {
		if prevBuild, ok := h.getBuildBefore(init, args["build"]); ok {
			argsPrev := make(map[string]string, len(args))
			for k, v := range args {
				argsPrev[k] = v
			}
			var err error
			argsPrev["build"] = prevBuild
			res.Previous, _, _, err = h.getBuildWinrates(ctx, init, argsPrev)
			return errors.Wrapf(err, "fetch previous build: %v", prevBuild)
		}
		return nil
	})
	err := g.Wait()
	m := make(map[string]talentText)
	for _, talents := range res.Current {
		for id := range talents {
			m[id] = talentData[id]
		}
	}
	res.Talents = m
	return res, err
}

type build struct {
	Build   []string
	Total   int
	Winrate float64
}

func (h *hotsContext) getBuildWinrates(ctx context.Context, init initData, args map[string]string) (map[int]map[string]Total, []build, []build, error) {
	if args["build"] == "" {
		return nil, nil, nil, errors.New("build required")
	}
	if args["hero"] == "" {
		return nil, nil, nil, errors.New("hero required")
	}

	groups := []string{"talents", "winner"}
	var wheres []string
	var params []interface{}
	for _, key := range []string{"build", "hero", "map", "mode"} {
		v := args[key]
		if v == "" {
			continue
		}
		if m, ok := init.config.Map[key]; ok {
			v = m[v]
			if v == "" {
				return nil, nil, nil, errors.Errorf("unrecognized %s: %s", key, args[key])
			}
		}
		wheres = append(wheres, fmt.Sprintf("%s = $%d", key, len(params)+1))
		groups = append(groups, key)
		params = append(params, v)
	}
	hl := args["herolevel"]
	if hl == "" {
		hl = defaultHerolevel
	}
	wheres = append(wheres, fmt.Sprintf("hero_level >= $%d", len(params)+1))
	params = append(params, hl)
	if err := setSkillParams(init, &wheres, &params, args); err != nil {
		return nil, nil, nil, err
	}

	buf := bytes.NewBufferString(`SELECT COUNT(*) count, winner, talents`)
	buf.WriteString(" FROM players")
	wheres = append(wheres, "array_length(talents, 1) > 0")
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
		Count   int
		Talents string
		Winner  bool
	}
	if err := h.x.SelectContext(ctx, &winrates, query, params...); err != nil {
		return nil, nil, nil, errors.Wrap(err, "select")
	}
	total := make(map[string]struct {
		Total int
		Won   int
	})
	for _, wr := range winrates {
		talents := wr.Talents[1 : len(wr.Talents)-1]
		t := total[talents]
		t.Total += wr.Count
		if wr.Winner {
			t.Won += wr.Count
		}
		total[talents] = t
		for tier, talent := range strings.Split(talents, ",") {
			tier += 1
			talent := string(init.lookups["talent"](string(talent)))
			t := tally[tier][talent]
			if wr.Winner {
				t.Wins += wr.Count
			} else {
				t.Losses += wr.Count
			}
			tally[tier][talent] = t
		}
	}
	var builds []build
	for n, t := range total {
		talents := strings.Split(n, ",")
		var talentNames []string
		for _, talent := range talents {
			talentNames = append(talentNames, init.lookups["talent"](talent))
		}
		b := build{
			Build: talentNames,
			Total: t.Total,
		}
		if t.Won > popularGameLimit {
			b.Winrate = float64(t.Won) / float64(t.Total)
		}
		builds = append(builds, b)
	}
	n := 5
	if n > len(builds) {
		n = len(builds)
	}
	var popularBuilds, winningBuilds []build
	sort.Slice(builds, func(i, j int) bool {
		return builds[i].Total > builds[j].Total
	})
	for i := 0; i < n; i++ {
		b := builds[i]
		if b.Total > popularGameLimit {
			popularBuilds = append(popularBuilds, b)
		}
	}
	sort.Slice(builds, func(i, j int) bool {
		return builds[i].Winrate > builds[j].Winrate
	})
	for i := 0; i < n; i++ {
		b := builds[i]
		if b.Winrate > 0 {
			winningBuilds = append(winningBuilds, b)
		}
	}
	return tally, popularBuilds, winningBuilds, nil
}

func (h *hotsContext) GetPlayerName(ctx context.Context, r *http.Request) (interface{}, error) {
	name := strings.ToLower(r.FormValue("name"))
	if name == "" {
		return nil, errors.New("no name parameter")
	}
	region := r.FormValue("region")
	if region == "" {
		return nil, errors.New("no region parameter")
	}
	type entry struct {
		ID     int64
		Region int
		Name   string
		Games  int
	}
	res := make([]entry, 0)
	var last string
	seen := make(map[int64]bool)
	for len(res) < 10 {
		var e entry
		err := h.x.GetContext(ctx, &e, `
			SELECT blizzid id, battletag "name", region
			FROM players
			WHERE battletag >= $1 COLLATE en_u_ks_level1
			AND battletag > $2 COLLATE en_u_ks_level1
			AND region = $3
			ORDER BY region, battletag
			LIMIT 1
		`, name, last, region)
		if err == sql.ErrNoRows {
			break
		} else if err != nil {
			return nil, err
		}
		if !strings.HasPrefix(strings.ToLower(e.Name), name) {
			break
		}
		last = e.Name
		if seen[e.ID] {
			continue
		}
		seen[e.ID] = true
		res = append(res, e)
	}
	for i, e := range res {
		if err := h.x.GetContext(ctx, &res[i], `
			SELECT count(*) games
			FROM players
			WHERE region = $1 AND blizzid = $2
		`, e.Region, e.ID); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (h *hotsContext) GetPlayerProfile(ctx context.Context, r *http.Request) (interface{}, error) {
	blizzid := r.FormValue("blizzid")
	build := r.FormValue("build")
	region := r.FormValue("region")
	if blizzid == "" {
		return nil, errors.New("no blizzid parameter")
	}
	if region == "" {
		return nil, errors.New("no region parameter")
	}
	if build == "" {
		return nil, errors.New("no build parameter")
	}
	init := h.getInit()
	var wheres []string
	var params []interface{}
	for _, key := range []string{"blizzid", "region", "hero"} {
		v := r.FormValue(key)
		if v == "" {
			continue
		}
		if m, ok := init.config.Map[key]; ok {
			v = m[v]
			if v == "" {
				return nil, errors.Errorf("unrecognized %s: %s", key, r.FormValue(key))
			}
		}
		wheres = append(wheres, fmt.Sprintf("%s = $%d", key, len(params)+1))
		params = append(params, v)
	}
	skillBuild := build
	if days, _ := strconv.Atoi(build); days > 0 && days < 100 {
		v := time.Now().UTC().Add(-time.Hour * 24 * time.Duration(days)).Format("2006-01-02")
		wheres = append(wheres, fmt.Sprintf("time >= $%d", len(params)+1))
		params = append(params, v)
		skillBuild = init.Builds[0].ID
	} else {
		key := "build"
		v := init.config.Map[key][build]
		if v == "" {
			return nil, errors.Errorf("unrecognized %s: %s", key, build)
		}
		wheres = append(wheres, fmt.Sprintf("%s = $%d", key, len(params)+1))
		params = append(params, v)
	}

	type buildSkill struct {
		Mode  Mode
		Build string
		Skill float64
		Stats Stats
	}

	var res struct {
		Battletag string
		Profile   struct {
			Heroes map[string]Total
			Maps   map[string]Total
			Modes  map[string]Total
			Roles  map[string]Total
		}
		Skills     map[Mode]buildSkill
		AllSkills  []buildSkill
		BuildStats map[Mode]Stats
	}
	res.Skills = make(map[Mode]buildSkill)
	res.Profile.Heroes = make(map[string]Total)
	res.Profile.Maps = make(map[string]Total)
	res.Profile.Modes = make(map[string]Total)
	res.Profile.Roles = make(map[string]Total)

	{
		v := init.config.Map["build"][skillBuild]
		if v == "" {
			return nil, errors.Errorf("unrecognized %s", skillBuild)
		}
		if err := h.x.SelectContext(ctx, &res.AllSkills, `
				SELECT build, skill, mode
				FROM playerskills
				WHERE region = $1 AND blizzid = $2
				`, region, blizzid); err != nil {
			return nil, err
		}
		for i, s := range res.AllSkills {
			build := init.lookups["build"](s.Build)
			res.AllSkills[i].Build = build
			res.AllSkills[i].Stats = init.BuildStats[build][s.Mode]
			if build <= skillBuild && build > res.Skills[s.Mode].Build {
				res.Skills[s.Mode] = res.AllSkills[i]
			}
		}
		sort.Slice(res.AllSkills, func(i, j int) bool {
			return res.AllSkills[i].Build < res.AllSkills[j].Build
		})
		res.BuildStats = init.BuildStats[skillBuild]
	}

	var err error
	res.Battletag, err = h.getBattletag(ctx, blizzid, region)
	if err != nil {
		return nil, err
	}

	var games []struct {
		Hero   string
		Winner bool
		Map    string
		Mode   string
	}
	if err := h.x.SelectContext(ctx, &games, fmt.Sprintf(`
			SELECT hero, winner, map, mode
			FROM players
			WHERE %s
			`, strings.Join(wheres, " AND ")), params...); err != nil {
		return nil, err
	}
	count := func(winner bool, m map[string]Total, name string) {
		v := m[name]
		if winner {
			v.Wins++
		} else {
			v.Losses++
		}
		m[name] = v
	}
	roles := make(map[string]string)
	for _, h := range init.Heroes {
		roles[h.Name] = h.Role
	}
	for _, g := range games {
		g.Hero = init.lookups["hero"](g.Hero)
		g.Map = init.lookups["map"](g.Map)
		count(g.Winner, res.Profile.Heroes, g.Hero)
		count(g.Winner, res.Profile.Maps, g.Map)
		count(g.Winner, res.Profile.Modes, g.Mode)
		count(g.Winner, res.Profile.Roles, roles[g.Hero])
	}

	return res, nil
}

func (h *hotsContext) GetPlayerFriends(ctx context.Context, r *http.Request) (interface{}, error) {
	blizzid := r.FormValue("blizzid")
	region := r.FormValue("region")
	if blizzid == "" {
		return nil, errors.New("no blizzid parameter")
	}
	if region == "" {
		return nil, errors.New("no region parameter")
	}

	res := struct {
		Battletag string
		Region    string
		Friends   []struct {
			Battletag string
			Games     int
			Blizzid   string
			Winrate   float64
		}
	}{
		Region: region,
	}

	if err := h.x.SelectContext(ctx, &res.Friends, `
		SELECT battletag, games, blizzid, won / games * 100 as winrate FROM (
			SELECT
				least(o.battletag) battletag,
				o.blizzid,
				count(*) games,
				count(nullif(o.winner, false)) won
			FROM players p
			JOIN players o ON
				o.game = p.game AND
				o.blizzid != p.blizzid AND
				NOT (o.team != p.team)
			WHERE
				p.region = $1 AND
				p.blizzid = $2
			GROUP BY o.blizzid, o.battletag
		)
		WHERE games > 5
		ORDER BY games DESC
		LIMIT 40
	`, region, blizzid); err != nil {
		return nil, err
	}

	var err error
	res.Battletag, err = h.getBattletag(ctx, blizzid, region)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *hotsContext) GetPlayerGames(ctx context.Context, r *http.Request) (interface{}, error) {
	blizzid := r.FormValue("blizzid")
	build := r.FormValue("build")
	region := r.FormValue("region")
	if blizzid == "" {
		return nil, errors.New("no blizzid parameter")
	}
	if region == "" {
		return nil, errors.New("no region parameter")
	}
	if build == "" {
		return nil, errors.New("no build parameter")
	}
	var res struct {
		Battletag string
		Games     []struct {
			Game      int
			Hero      string
			HeroLevel int    `db:"hero_level"`
			Date      string `db:"time"`
			Build     string
			Winner    bool
			Length    int
			Map       string
			Mode      Mode
			Skill     *int
		}
	}

	init := h.getInit()
	wheres := []string{"blizzid = $1", "region = $2"}
	params := []interface{}{blizzid, region}
	if days, _ := strconv.Atoi(build); days > 0 && days < 100 {
		v := time.Now().UTC().Add(-time.Hour * 24 * time.Duration(days)).Format("2006-01-02")
		wheres = append(wheres, fmt.Sprintf("time >= $%d", len(params)+1))
		params = append(params, v)
	} else {
		key := "build"
		v := init.config.Map[key][build]
		if v == "" {
			return nil, errors.Errorf("unrecognized %s: %s", key, build)
		}
		wheres = append(wheres, fmt.Sprintf("%s = $%d", key, len(params)+1))
		params = append(params, v)
	}

	var err error
	res.Battletag, err = h.getBattletag(ctx, blizzid, region)
	if err != nil {
		return nil, err
	}

	if err := h.x.SelectContext(ctx, &res.Games, fmt.Sprintf(`
			SELECT game, hero, hero_level, build, winner, length, map, mode, time
			FROM players
			WHERE %s
			ORDER BY time DESC
			`, strings.Join(wheres, " AND ")), params...); err != nil {
		return nil, err
	}
	for i, g := range res.Games {
		g.Hero = init.lookups["hero"](g.Hero)
		g.Map = init.lookups["map"](g.Map)
		g.Build = init.lookups["build"](g.Build)
		res.Games[i] = g
	}

	return res, nil
}

func (h *hotsContext) getBattletag(ctx context.Context, blizzid, region string) (string, error) {
	var battletag string
	err := h.x.GetContext(ctx, &battletag, `
		SELECT battletag
		FROM players
		WHERE blizzid = $1 and region = $2
		ORDER BY time DESC
		LIMIT 1
		`, blizzid, region)
	return battletag, err
}

func (h *hotsContext) GetPlayerMatchups(ctx context.Context, r *http.Request) (interface{}, error) {
	blizzid := r.FormValue("blizzid")
	build := r.FormValue("build")
	region := r.FormValue("region")
	if blizzid == "" {
		return nil, errors.New("no blizzid parameter")
	}
	if region == "" {
		return nil, errors.New("no region parameter")
	}
	if build == "" {
		return nil, errors.New("no build parameter")
	}
	res := struct {
		Battletag string
		Same      map[string]Total
		Opposing  map[string]Total
	}{
		Same:     make(map[string]Total),
		Opposing: make(map[string]Total),
	}

	init := h.getInit()
	wheres := []string{"blizzid = $1", "region = $2"}
	params := []interface{}{blizzid, region}
	if days, _ := strconv.Atoi(build); days > 0 && days < 100 {
		v := time.Now().UTC().Add(-time.Hour * 24 * time.Duration(days)).Format("2006-01-02")
		wheres = append(wheres, fmt.Sprintf("time >= $%d", len(params)+1))
		params = append(params, v)
	} else {
		key := "build"
		v := init.config.Map[key][build]
		if v == "" {
			return nil, errors.Errorf("unrecognized %s: %s", key, build)
		}
		wheres = append(wheres, fmt.Sprintf("%s = $%d", key, len(params)+1))
		params = append(params, v)
	}

	var err error
	res.Battletag, err = h.getBattletag(ctx, blizzid, region)
	if err != nil {
		return nil, err
	}

	type Game struct {
		Game   int
		Winner bool
		Hero   string
	}
	var games, all []Game
	if err := h.x.SelectContext(ctx, &games, fmt.Sprintf(`
			SELECT game, winner, hero
			FROM players
			WHERE %s
			`, strings.Join(wheres, " AND ")), params...); err != nil {
		return nil, err
	}
	if len(games) == 0 {
		return res, nil
	}
	ids := make([]interface{}, len(games))
	gs := make(map[int]Game)
	for i, g := range games {
		ids[i] = g.Game
		gs[g.Game] = g
	}
	if err := h.x.SelectContext(ctx, &all, fmt.Sprintf(`
		SELECT game, winner, hero
		FROM players
		WHERE game IN %s
		`, makeValues(len(games), 1)), ids...); err != nil {
		return nil, err
	}
	count := func(winner bool, m map[string]Total, name string) {
		v := m[name]
		if winner {
			v.Wins++
		} else {
			v.Losses++
		}
		m[name] = v
	}
	for _, g := range all {
		pg := gs[g.Game]
		if pg == g {
			continue
		}
		if pg.Winner == g.Winner {
			count(pg.Winner, res.Same, init.lookups["hero"](g.Hero))
		} else {
			count(pg.Winner, res.Opposing, init.lookups["hero"](g.Hero))
		}
	}

	return res, nil
}

func (h *hotsContext) GetGameData(ctx context.Context, r *http.Request) (interface{}, error) {
	id := r.FormValue("id")
	init := h.getInit()
	if id == "" {
		return nil, errors.New("no id parameter")
	}
	res := struct {
		Game struct {
			Mode    Mode
			Date    string `db:"time"`
			Map     string
			Length  int
			Build   string
			Bans    string `json:"-"`
			BanList []string
		}
		Players []*struct {
			Hero       string
			HeroLevel  int `db:"hero_level"`
			Winner     bool
			Blizzid    int
			Battletag  string
			Talents    string `json:"-"`
			TalentList []string
			Data       json.RawMessage
			Region     int
		}
		Talents map[string]talentText
	}{
		Talents: make(map[string]talentText),
	}

	if err := h.x.GetContext(ctx, &res.Game, `
		SELECT mode, time, map, length, build, bans
		FROM games
		WHERE id = $1
		`, id); err != nil {
		return nil, err
	}

	if err := h.x.SelectContext(ctx, &res.Players, `
		SELECT hero, hero_level, winner, blizzid, battletag, talents, data, region
		FROM players
		WHERE game = $1
		`, id); err != nil {
		return nil, err
	}
	res.Game.BanList = init.list("hero", res.Game.Bans)
	res.Game.Map = init.lookups["map"](res.Game.Map)
	res.Game.Build = init.lookups["build"](res.Game.Build)
	for _, p := range res.Players {
		p.Hero = init.lookups["hero"](p.Hero)
		p.TalentList = init.list("talent", p.Talents)
		for _, t := range p.TalentList {
			res.Talents[t] = talentData[t]
		}
	}

	return res, nil
}

type heroRelativeData struct {
	Base    map[string]Total
	Lengths map[string]Total
	Levels  map[string]Total
	Maps    map[string]Total
	Modes   map[string]Total
}

func (h *hotsContext) GetHero(ctx context.Context, r *http.Request) (interface{}, error) {
	init := h.getInit()
	build := init.config.build(r.FormValue("build"))
	hero := init.config.hero(r.FormValue("hero"))
	var res struct {
		Current  heroRelativeData
		Previous heroRelativeData
	}
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		res.Current, err = h.getHero(ctx, init, build, hero)
		return errors.Wrap(err, "getHero current build")
	})
	g.Go(func() error {
		if prevBuild, ok := h.getBuildBefore(init, r.FormValue("build")); ok {
			var err error
			res.Previous, err = h.getHero(ctx, init, init.config.build(prevBuild), hero)
			return errors.Wrap(err, "getHero previous build")
		}
		return nil
	})
	err := g.Wait()
	return res, err
}

func (h *hotsContext) getHero(ctx context.Context, init initData, build, hero string) (heroRelativeData, error) {
	params := []interface{}{
		build,
		hero,
	}
	var res heroRelativeData
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		res.Base, err = h.countWins(ctx, nil, `
			SELECT COUNT(*) count, winner, '' counter
			FROM players
			WHERE build = $1 AND hero = $2
			GROUP BY winner
		`, params)
		if res.Base != nil && len(res.Base) == 0 {
			res.Base[""] = Total{}
		}
		return errors.Wrap(err, "base")
	})
	g.Go(func() error {
		var err error
		res.Maps, err = h.countWins(ctx, init.lookups["map"], `
			SELECT COUNT(*) count, winner, map counter
			FROM players
			WHERE build = $1 AND hero = $2
			GROUP BY winner, map
		`, params)
		return errors.Wrap(err, "maps")
	})
	g.Go(func() error {
		var err error
		res.Modes, err = h.countWins(ctx, nil, `
			SELECT COUNT(*) count, winner, mode counter
			FROM players
			WHERE build = $1 AND hero = $2
			GROUP BY winner, mode
		`, params)
		return errors.Wrap(err, "modes")
	})
	g.Go(func() error {
		var err error
		// Group hero levels by 5s.
		res.Levels, err = h.countWins(ctx, nil, `
			SELECT count(*) count, winner, greatest(1, counter * 5) as counter
			FROM (
				SELECT winner, hero_level // 5 as counter
				FROM players
				WHERE build = $1 AND hero = $2
			)
			GROUP BY winner, counter
		`, params)
		return errors.Wrap(err, "hero level")
	})
	g.Go(func() error {
		var err error
		// Group game lengths in 5 minute blocks.
		res.Lengths, err = h.countWins(ctx, nil, `
			SELECT count(*) count, winner, counter * 60 * 5 as counter
			FROM (
				SELECT winner, round(length / 60 / 5) as counter
				FROM players
				WHERE build = $1 AND hero = $2
			)
			GROUP BY winner, counter
		`, params)
		return errors.Wrap(err, "length")
	})
	err := g.Wait()
	return res, err
}

func (h *hotsContext) countWins(ctx context.Context, nameFn func(string) string, query string, params []interface{}) (map[string]Total, error) {
	tally := make(map[string]Total)
	var winrates []struct {
		Counter string
		Count   int
		Winner  bool
	}
	if err := h.x.SelectContext(ctx, &winrates, query, params...); err != nil {
		return nil, errors.Wrap(err, "select wins")
	}
	for _, wr := range winrates {
		n := wr.Counter
		if nameFn != nil {
			n = nameFn(n)
		}
		t := tally[n]
		if wr.Winner {
			t.Wins += wr.Count
		} else {
			t.Losses += wr.Count
		}
		tally[n] = t
	}
	return tally, nil
}

func (h *hotsContext) getInit() initData {
	h.mu.RLock()
	init := h.mu.init
	h.mu.RUnlock()
	return init
}

func (h *hotsContext) GetWinrates(ctx context.Context, r *http.Request) (interface{}, error) {
	args := map[string]string{
		"build":      r.FormValue("build"),
		"herolevel":  r.FormValue("herolevel"),
		"map":        r.FormValue("map"),
		"mode":       r.FormValue("mode"),
		"skill_low":  r.FormValue("skill_low"),
		"skill_high": r.FormValue("skill_high"),
	}
	init := h.getInit()
	var res struct {
		Current  map[string]Total
		Previous map[string]Total
	}
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		res.Current, err = h.getWinrates(ctx, init, args)
		return errors.Wrap(err, "getWinrates current build")
	})
	g.Go(func() error {
		if prevBuild, ok := h.getBuildBefore(init, args["build"]); ok {
			argsPrev := make(map[string]string, len(args))
			for k, v := range args {
				argsPrev[k] = v
			}
			var err error
			argsPrev["build"] = prevBuild
			res.Previous, err = h.getWinrates(ctx, init, argsPrev)
			return errors.Wrap(err, "getWinrates previous build")
		}
		return nil
	})
	err := g.Wait()
	return res, err
}

const defaultHerolevel = "5"

func (h *hotsContext) getWinrates(ctx context.Context, init initData, args map[string]string) (map[string]Total, error) {
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
		if m, ok := init.config.Map[key]; ok {
			v = m[v]
			if v == "" {
				return nil, errors.Errorf("unrecognized %s: %s", key, args[key])
			}
		}
		wheres = append(wheres, fmt.Sprintf("%s = $%d", key, len(params)+1))
		params = append(params, v)
	}
	hl := args["herolevel"]
	if hl == "" {
		hl = defaultHerolevel
	}
	wheres = append(wheres, fmt.Sprintf("hero_level >= $%d", len(params)+1))
	params = append(params, hl)
	if err := setSkillParams(init, &wheres, &params, args); err != nil {
		return nil, err
	}

	buf := bytes.NewBufferString("SELECT COUNT(*) count, hero counter, winner FROM players")
	if len(wheres) > 0 {
		fmt.Fprintf(buf, " WHERE %s", strings.Join(wheres, " AND "))
	}
	buf.WriteString(" GROUP BY ")
	buf.WriteString(strings.Join(groups, ", "))

	return h.countWins(ctx, init.lookups["hero"], buf.String(), params)
}

func setSkillParams(init initData, wheres *[]string, params *[]interface{}, args map[string]string) error {
	if m, sl, sh := args["mode"], args["skill_low"], args["skill_high"]; m != "" && (sl != "" || sh != "") {
		i, err := strconv.Atoi(m)
		if err != nil {
			return err
		}
		modes, ok := init.BuildStats[args["build"]]
		if !ok {
			return errors.Errorf("unknown build: %s", args["build"])
		}
		quantiles := modes[Mode(i)].Quantile
		if sl != "" {
			*wheres = append(*wheres, fmt.Sprintf("skill >= $%d", len(*params)+1))
			i, err = strconv.Atoi(sl)
			if err != nil {
				return err
			}

			*params = append(*params, quantiles[skillQuantiles[i]])
		}
		if sh != "" {
			*wheres = append(*wheres, fmt.Sprintf("skill <= $%d", len(*params)+1))
			i, err = strconv.Atoi(sh)
			if err != nil {
				return err
			}
			*params = append(*params, quantiles[skillQuantiles[i+1]])
		}
	}
	return nil
}

type Total struct {
	Wins, Losses int
}

func (h *hotsContext) GetCompareHero(ctx context.Context, r *http.Request) (interface{}, error) {
	init := h.getInit()
	args := map[string]string{
		"build":     init.config.build(r.FormValue("build")),
		"hero":      init.config.hero(r.FormValue("hero")),
		"herolevel": r.FormValue("herolevel"),
		"map":       init.config.gamemap(r.FormValue("map")),
		"mode":      r.FormValue("mode"),
	}
	if args["build"] == "" {
		return nil, errors.New("build required")
	}
	if args["hero"] == "" {
		return nil, errors.New("hero required")
	}

	var wheres []string
	var params []interface{}
	for _, key := range []string{"build", "hero", "map", "mode"} {
		v := args[key]
		if v == "" {
			continue
		}
		wheres = append(wheres, fmt.Sprintf("%s = $%d", key, len(params)+1))
		params = append(params, v)
	}
	hl := args["herolevel"]
	if hl == "" {
		hl = defaultHerolevel
	}
	wheres = append(wheres, fmt.Sprintf("hero_level >= $%d", len(params)+1))
	params = append(params, hl)
	if err := setSkillParams(init, &wheres, &params, map[string]string{
		"mode":       r.FormValue("mode"),
		"build":      r.FormValue("build"),
		"skill_low":  r.FormValue("skill_low"),
		"skill_high": r.FormValue("skill_high"),
	}); err != nil {
		return nil, err
	}
	var games []struct {
		Game   int64
		Team   int
		Winner bool
	}
	if err := h.x.SelectContext(ctx, &games, fmt.Sprintf(`
		SELECT game, team, winner
		FROM players
		WHERE %s
		`, strings.Join(wheres, " AND "),
	), params...); err != nil {
		return nil, err
	}
	var total Total
	team0 := make([]interface{}, 0, len(games))
	team1 := make([]interface{}, 0, len(games))
	for _, g := range games {
		if g.Team == 0 {
			team0 = append(team0, g.Game)
		} else {
			team1 = append(team1, g.Game)
		}
		if g.Winner {
			total.Wins++
		} else {
			total.Losses++
		}
	}
	sameTeam := make(map[string]Total)
	otherTeam := make(map[string]Total)
	getGames := func(team int, ids []interface{}) error {
		for len(ids) > 0 {
			limit := 1000
			if limit > len(ids) {
				limit = len(ids)
			}
			params := append([]interface{}{team, args["hero"]}, ids[:limit]...)
			query := fmt.Sprintf(`
					SELECT hero, team = $1 as sameteam, winner, count(*) as count
					FROM players
					WHERE
						game IN %s
						AND hero != $2
					GROUP BY hero, team, winner
					`, makeValues(limit, 3),
			)
			var res []struct {
				Hero     string
				Sameteam bool
				Winner   bool
				Count    int
			}
			if err := h.x.SelectContext(ctx, &res, query, params...); err != nil {
				return errors.Wrap(err, "select")
			}
			for _, r := range res {
				team := sameTeam
				if !r.Sameteam {
					team = otherTeam
				}
				hero := init.lookups["hero"](r.Hero)
				t := team[hero]
				if r.Winner == r.Sameteam {
					t.Wins += r.Count
				} else {
					t.Losses += r.Count
				}
				team[hero] = t
			}
			ids = ids[limit:]
		}
		return nil
	}
	if err := getGames(0, team0); err != nil {
		return nil, err
	}
	if err := getGames(1, team1); err != nil {
		return nil, err
	}
	return struct {
		SameTeam  map[string]Total
		OtherTeam map[string]Total
		Total     Total
	}{
		SameTeam:  sameTeam,
		OtherTeam: otherTeam,
		Total:     total,
	}, nil
}

func (h *hotsContext) GetLeaderboard(ctx context.Context, r *http.Request) (interface{}, error) {
	init := h.getInit()
	mode := r.FormValue("mode")
	region := r.FormValue("region")
	if mode == "" {
		return nil, errors.New("mode required")
	}
	if region == "" {
		return nil, errors.New("region required")
	}
	const (
		maxPlayers           = 100
		fetchPlayersPerPatch = 1000
	)
	var daysTime = time.Hour * 24 * time.Duration(daysOld)
	since := time.Now().Add(-daysTime)
	type playerSkill struct {
		Blizzid string
		Skill   float64
	}
	// blizzid -> playerSkill
	players := make(map[string]playerSkill)
	for _, b := range init.Builds {
		if b.Finish.Before(since) {
			continue
		}

		var ps []playerSkill
		if err := h.x.SelectContext(ctx, &ps, `
		SELECT blizzid, skill
		FROM playerskills
		WHERE
			region = $1 AND
			build = $2 AND
			mode = $3
		ORDER BY skill DESC
		LIMIT $4
	`, region, init.config.build(b.ID), mode, fetchPlayersPerPatch); err != nil {
			return nil, err
		}
		for _, p := range ps {
			if _, ok := players[p.Blizzid]; !ok {
				players[p.Blizzid] = p
			}
		}
	}
	skills := make([]playerSkill, 0, len(players))
	for _, p := range players {
		skills = append(skills, p)
	}
	sort.Slice(skills, func(i, j int) bool {
		return skills[i].Skill > skills[j].Skill
	})

	type rankPlayer struct {
		Rank      int
		Battletag string
		Blizzid   string
		Skill     float64
		Games     int
	}

	res := struct {
		Players  []*rankPlayer
		Attempts int
		Days     int
		MinGames int
	}{
		Players:  make([]*rankPlayer, 0, maxPlayers),
		Days:     daysOld,
		MinGames: leaderboardMinGames,
	}

	g, gCtx := errgroup.WithContext(ctx)
	for _, s := range skills {
		if len(res.Players) >= maxPlayers {
			break
		}
		var games int
		res.Attempts++
		if err := h.x.GetContext(ctx, &games, `
			SELECT count(*)
			FROM players
			WHERE
				region = $1 AND
				blizzid = $2 AND
				"time" > $3 AND
				mode = $4
		`, region, s.Blizzid, since, mode); err != nil {
			return nil, err
		}
		if games < leaderboardMinGames {
			continue
		}
		p := &rankPlayer{
			Rank:    len(res.Players) + 1,
			Blizzid: s.Blizzid,
			Skill:   s.Skill,
			Games:   games,
		}
		res.Players = append(res.Players, p)
		g.Go(func() error {
			var err error
			p.Battletag, err = h.getBattletag(gCtx, p.Blizzid, region)
			return err
		})
	}
	err := g.Wait()
	return res, err
}

func (h *hotsContext) getBuildBefore(init initData, id string) (build string, ok bool) {
	for i, b := range init.Builds {
		if b.ID == id {
			if len(init.Builds) == i+1 {
				return "", false
			}
			return init.Builds[i+1].ID, true
		}
	}
	return "", false
}

var capsRE = regexp.MustCompile(`[A-Z][a-z]+`)
var font *truetype.Font

func init() {
	var err error
	font, err = freetype.ParseFont(goregular.TTF)
	if err != nil {
		panic(err)
	}
}

func makeTalentImg(w http.ResponseWriter, r *http.Request) {
	idx := strings.LastIndexByte(r.URL.Path, '/')
	name := r.URL.Path[idx+1:]
	words := capsRE.FindAllStringSubmatch(name, 4)
	i := image.NewRGBA(image.Rect(0, 0, 40, 40))
	draw.Draw(i, i.Bounds(), &image.Uniform{image.White}, image.ZP, draw.Src)

	const size = 10
	c := freetype.NewContext()
	c.SetFont(font)
	c.SetFontSize(size)
	c.SetClip(i.Bounds())
	c.SetDst(i)
	c.SetSrc(image.Black)

	for i, word := range words {
		if _, err := c.DrawString(word[0], freetype.Pt(1, (i+1)*size-2)); err != nil {
			log.Printf("%s: %+v", r.URL.Path, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Add("Cache-Control", "max-age=3600")
	if err := png.Encode(w, i); err != nil {
		log.Printf("%s: %+v", r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
