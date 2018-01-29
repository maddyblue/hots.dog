package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/sync/errgroup"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/lib/pq"
	"github.com/pkg/errors"

	"cloud.google.com/go/storage"
)

const (
	HotsApi    = "https://hotsapi.net/api/v1"
	configJSON = "config.json"
	dirGame    = "game"
	dirPlayer  = "player"
	perFile    = 10000
	perRequest = 100
	configBase = "%09d.csv"
)

func (h *hotsContext) updateDB() error {
	ctx := context.Background()
	cl, err := storage.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "new client")
	}
	// Possibly we are on a new image, so update cache if the underlying
	// implementation changed.
	updated := true
	bucket := cl.Bucket(*flagImport)
	for {
		err := h.updateDBNext(bucket)
		if err == storage.ErrObjectNotExist {
			fmt.Println("no new data; sleeping")
			if updated {
				updated = false
				if err := h.cronLoop(); err != nil {
					return errors.Wrap(err, "cronLoop")
				}
			}
			time.Sleep(time.Minute * 10)
		} else if err != nil {
			return errors.Wrap(err, "updateDBNext")
		} else if err == nil {
			updated = true
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
			log.Printf("update next group error: %+v", err)
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
	{
		// Make sure the current block is present by fetching the next one.
		url := fmt.Sprintf(HotsApi+"/replays?min_id=%d", until)
		b, err := httpGet(ctx, url)
		if err != nil {
			return 0, errors.Wrapf(err, "http get: %s", b)
		}
		var replays Replays
		if err := json.Unmarshal(b, &replays); err != nil {
			return 0, errors.Wrap(err, "json decode")
		}
		if len(replays) == 0 {
			return 0, ErrNotEnough
		}
	}
	base := fmt.Sprintf(configBase, start)
	{
		obj := bucket.Object(path.Join("test", base))
		gw := obj.NewWriter(ctx)
		if _, err := gw.Write([]byte("blah")); err != nil {
			return 0, errors.Wrap(err, "gb write test")
		}
		if err := gw.Close(); err != nil {
			return 0, errors.Wrap(err, "gw close test")
		}
		if err := obj.Delete(ctx); err != nil {
			return 0, errors.Wrap(err, "gw delete test")
		}
		fmt.Println("write test passed")
	}
	var gb, pb bytes.Buffer
	gc := csv.NewWriter(&gb)
	pc := csv.NewWriter(&pb)
	{
		idCh := make(chan int)
		replaysCh := make(chan []byte, 1)
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
			sub, subCtx := errgroup.WithContext(ctx)
			for b := range replaysCh {
				var replays Replays
				if err := json.Unmarshal(b, &replays); err != nil {
					return errors.Wrap(err, "json decode")
				}
				for _, r := range replays {
					if r.ID >= until {
						continue
					}
					r := r
					sub.Go(func() error {
						if !r.Processed {
							if err := processReplay(subCtx, &r, config); err != nil {
								return errors.Wrapf(err, "processing error %d", r.ID)
							}
						}
						select {
						case <-done:
							return ctx.Err()
						case replayCh <- r:
						}
						return nil
					})
				}
			}
			return sub.Wait()
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
				var bans []string
				for _, bs := range r.Bans {
					for _, b := range bs {
						if b != "" {
							bans = append(bans, config.hero(b))
						}
					}
				}
				version := strings.Join(strings.Split(r.GameVersion, ".")[:3], ".")
				config.addBuild(version, time.Time(r.GameDate))
				common := []string{
					fmt.Sprint(r.ID),
					fmt.Sprint(mode),
					time.Time(r.GameDate).Format(time.RFC3339Nano),
					config.gamemap(r.GameMap),
					fmt.Sprint(r.GameLength),
					config.build(version),
					fmt.Sprint(r.Region),
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
					data, err := json.Marshal(p.Score)
					if err != nil {
						return errors.Wrap(err, "marshal score")
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
						string(data),
					)); err != nil {
						return errors.Wrap(err, "gc write")
					}
				}
			}
			return nil
		})
		if err := g.Wait(); err != nil {
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
	gw := bucket.Object(path.Join(dirGame, base)).NewWriter(ctx)
	if _, err := gb.WriteTo(gw); err != nil {
		return 0, errors.Wrap(err, "gb write")
	}
	if err := gw.Close(); err != nil {
		return 0, errors.Wrap(err, "gw close")
	}
	pw := bucket.Object(path.Join(dirPlayer, base)).NewWriter(ctx)
	if _, err := pb.WriteTo(pw); err != nil {
		return 0, errors.Wrap(err, "pb write")
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

var (
	awsSess   *session.Session
	lambdaSvc *lambda.Lambda
	accessKey string
	secretKey string
	banMap    = map[string]string{}

	processSema = make(chan struct{}, 100)
)

func init() {
	creds := credentials.NewSharedCredentials("", "")
	value, err := creds.Get()
	if err != nil {
		panic(errors.Wrap(err, "get session creds"))
	}
	accessKey = value.AccessKeyID
	secretKey = value.SecretAccessKey
	awsSess = session.New(&aws.Config{
		Credentials: creds,
		Region:      aws.String("eu-west-1"),
	})
	lambdaSvc = lambda.New(awsSess)

	for _, h := range heroData {
		short := h.ID
		if len(short) > 4 {
			short = short[:4]
		}
		banMap[short] = h.Name
	}
}

func processReplay(ctx context.Context, r *Replay, g *groupConfig) error {
	select {
	case processSema <- struct{}{}:
		defer func() { <-processSema }()
	case <-ctx.Done():
		return ctx.Err()
	}
	fmt.Println("process", r.ID)

	b, err := json.Marshal(struct {
		Input  string `json:"input"`
		Access string `json:"access"`
		Secret string `json:"secret"`
	}{
		Input:  r.URL,
		Access: accessKey,
		Secret: secretKey,
	})
	if err != nil {
		return err
	}

	input := &lambda.InvokeInput{
		FunctionName: aws.String("parse-hots"),
		Payload:      b,
	}

	var result *lambda.InvokeOutput
	for retries := 0; true; retries++ {
		result, err = lambdaSvc.InvokeWithContext(ctx, input)
		if err != nil {
			fmt.Printf("invoke error: %s: %s", err, b)
			continue
		}
		if result.FunctionError != nil {
			if retries >= 10 {
				fmt.Printf("process %d giving up\n", r.ID)
				return nil
			}
			fmt.Printf("retry process (%d) %d: %d: %s: %s\n", retries, r.ID, *result.StatusCode, *result.FunctionError, result.Payload)
			if *result.FunctionError == "Unhandled" && strings.Contains(string(result.Payload), "Process exited before completing request") {
				continue
			}
			return nil
		}
		if result.StatusCode == nil || *result.StatusCode != 200 {
			return errors.New("bad status code")
		}
		break
	}
	var pr ProcessedReplay
	if err := json.Unmarshal(result.Payload, &pr); err != nil {
		fmt.Printf("processing %d json error: %s: %s\n", r.ID, err, result.Payload)
		return nil
	}

	players := make(map[int]ProcessedPlayer)
	for _, p := range pr.Players {
		players[p.BlizzID] = p
	}

	for _, p := range r.Players {
		np, ok := players[p.BlizzID]
		if !ok {
			return errors.Errorf("unfound player %d", p.BlizzID)
		}
		p.Battletag = fmt.Sprintf("%s#%d", np.BattletagName, np.BattletagID)
		switch len(np.Talents) {
		case 7:
			p.Talents.Num20 = np.Talents[6]
			fallthrough
		case 6:
			p.Talents.Num16 = np.Talents[5]
			fallthrough
		case 5:
			p.Talents.Num13 = np.Talents[4]
			fallthrough
		case 4:
			p.Talents.Num10 = np.Talents[3]
			fallthrough
		case 3:
			p.Talents.Num7 = np.Talents[2]
			fallthrough
		case 2:
			p.Talents.Num4 = np.Talents[1]
			fallthrough
		case 1:
			p.Talents.Num1 = np.Talents[0]
		}

		p.Score.Level = np.Score.Level
		p.Score.Takedowns = np.Score.Takedowns
		p.Score.Kills = np.Score.SoloKills
		p.Score.Assists = np.Score.Assists
		p.Score.Deaths = np.Score.Deaths
		p.Score.HighestKillStreak = np.Score.HighestKillStreak
		p.Score.HeroDamage = np.Score.HeroDamage
		p.Score.SiegeDamage = np.Score.SiegeDamage
		p.Score.StructureDamage = np.Score.StructureDamage
		p.Score.MinionDamage = np.Score.MinionDamage
		p.Score.CreepDamage = np.Score.CreepDamage
		p.Score.SummonDamage = np.Score.SummonDamage
		p.Score.TimeCcEnemyHeroes = np.Score.TimeCCdEnemyHeroes
		p.Score.Healing = np.Score.Healing
		p.Score.SelfHealing = np.Score.SelfHealing
		p.Score.DamageTaken = np.Score.DamageTaken
		p.Score.ExperienceContribution = np.Score.ExperienceContribution
		p.Score.TownKills = np.Score.TownKills
		p.Score.TimeSpentDead = np.Score.TimeSpentDead
		p.Score.MercCampCaptures = np.Score.MercCampCaptures
		p.Score.WatchTowerCaptures = np.Score.WatchTowerCaptures
		p.Score.MetaExperience = np.Score.MetaExperience
	}

	var bans []string
	for _, teamBans := range pr.Bans {
		for _, b := range teamBans {
			b, ok := banMap[b]
			if ok {
				bans = append(bans, b)
			}
		}
	}
	r.Bans = [][]string{bans}

	return nil
}

type ProcessedPlayer struct {
	BattletagName string   `json:"battletag_name"`
	BattletagID   int      `json:"battletag_id"`
	BlizzID       int      `json:"blizz_id"`
	Hero          string   `json:"hero"`
	HeroLevel     int      `json:"hero_level"`
	Team          int      `json:"team"`
	Winner        bool     `json:"winner"`
	Silenced      bool     `json:"silenced"`
	Party         int      `json:"party"`
	Talents       []string `json:"talents"`
	Score         struct {
		Level                  int   `json:"Level"`
		Takedowns              int   `json:"Takedowns"`
		SoloKills              int   `json:"SoloKills"`
		Assists                int   `json:"Assists"`
		Deaths                 int   `json:"Deaths"`
		HighestKillStreak      int   `json:"HighestKillStreak"`
		HeroDamage             int   `json:"HeroDamage"`
		SiegeDamage            int   `json:"SiegeDamage"`
		StructureDamage        int   `json:"StructureDamage"`
		MinionDamage           int   `json:"MinionDamage"`
		CreepDamage            int   `json:"CreepDamage"`
		SummonDamage           int   `json:"SummonDamage"`
		TimeCCdEnemyHeroes     int   `json:"TimeCCdEnemyHeroes"`
		Healing                int   `json:"Healing"`
		SelfHealing            int   `json:"SelfHealing"`
		DamageTaken            int   `json:"DamageTaken"`
		ExperienceContribution int   `json:"ExperienceContribution"`
		TownKills              int   `json:"TownKills"`
		TimeSpentDead          int   `json:"TimeSpentDead"`
		MercCampCaptures       int   `json:"MercCampCaptures"`
		WatchTowerCaptures     int   `json:"WatchTowerCaptures"`
		MetaExperience         int   `json:"MetaExperience"`
		MatchAwards            []int `json:"MatchAwards"`
	} `json:"score"`
}

type ProcessedReplay struct {
	Mode    string            `json:"mode"`
	Date    time.Time         `json:"date"`
	Length  string            `json:"length"`
	Map     string            `json:"map"`
	Version string            `json:"version"`
	Bans    [][]string        `json:"bans"`
	Players []ProcessedPlayer `json:"players"`
}

var httpClient = &http.Client{
	Timeout: time.Minute * 2,
}

func httpGet(ctx context.Context, url string) ([]byte, error) {
	start := time.Now()
	defer func() {
		fmt.Println("GET", url, "took", time.Since(start))
	}()
	for i := 0; i < 100; i++ {
		if i > 0 {
			fmt.Printf("retry %d: %s\n", i, url)
		}
		ctx, cancel := context.WithTimeout(ctx, time.Minute*2)
		defer cancel()
		resp, err := ctxhttp.Get(ctx, httpClient, url)
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
			log.Println(resp.Status, url)
			time.Sleep(time.Second * 30)
			continue
		}
		if resp.StatusCode != 200 {
			return nil, errors.Errorf("%s: %s: %s", url, resp.Status, b)
		}
		return b, nil
	}
	return nil, errors.Errorf("%s: too many retries", url)
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
	Players     []*struct {
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
			Level                  int `json:"level"`
			Kills                  int `json:"kills"`
			Assists                int `json:"assists"`
			Takedowns              int `json:"takedowns"`
			Deaths                 int `json:"deaths"`
			HighestKillStreak      int `json:"highest_kill_streak"`
			HeroDamage             int `json:"hero_damage"`
			SiegeDamage            int `json:"siege_damage"`
			StructureDamage        int `json:"structure_damage"`
			MinionDamage           int `json:"minion_damage"`
			CreepDamage            int `json:"creep_damage"`
			SummonDamage           int `json:"summon_damage"`
			TimeCcEnemyHeroes      int `json:"time_cc_enemy_heroes"`
			Healing                int `json:"healing"`
			SelfHealing            int `json:"self_healing"`
			DamageTaken            int `json:"damage_taken"`
			ExperienceContribution int `json:"experience_contribution"`
			TownKills              int `json:"town_kills"`
			TimeSpentDead          int `json:"time_spent_dead"`
			MercCampCaptures       int `json:"merc_camp_captures"`
			WatchTowerCaptures     int `json:"watch_tower_captures"`
			MetaExperience         int `json:"meta_experience"`
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
