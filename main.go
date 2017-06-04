package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	flagInit = flag.Bool("init", false, "drop database before starting")
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
			if r := recover(); r != nil {
				fmt.Printf("%+v\n", r)
				select {}
			}
		}()

		if _, err := db.Exec(fmt.Sprintf("drop database if exists %s; create database %[1]s", dbName)); err != nil {
			panic(err)
		}
	}

	mustMigrate(db)

	wrap := func(f func(context.Context, *http.Request, httprouter.Params) (interface{}, error)) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
			res, err := f(ctx, r, ps)
			if err != nil {
				log.Println(err)
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
				log.Printf("%+v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}

	router := httprouter.New()
	router.POST("/api/upload-replay", wrap(h.UploadReplay))
	router.GET("/api/get-winrates", wrap(h.GetWinrates))

	const addr = "localhost:4001"

	if *flagInit {
		go mustInitDevData(addr)
	}

	log.Fatal(http.ListenAndServe(addr, router))
}

type hotsContext struct {
	db *sql.DB
	x  *sqlx.DB
}

func (h *hotsContext) GetWinrates(ctx context.Context, r *http.Request, _ httprouter.Params) (interface{}, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	limits := make(map[string][]string)
	for key, values := range r.Form {
		if len(values) == 0 {
			continue
		}
		switch key {
		case "build",
			"map",
			"hero":
		default:
			return nil, errors.Errorf("unknown filter: %s", key)
		}
		limits[key] = values
	}
	var wheres []string
	groups := []string{"hero", "winner"}
	buf := bytes.NewBufferString("SELECT COUNT(*), hero, winner")
	if v := limits["map"]; v != nil {
		buf.WriteString(", map")
		s := fmt.Sprintf(" map IN %s", makeValues(len(v), len(v)))
		wheres = append(wheres, s)
		groups = append(groups, "map")
	}
	if v := limits["hero"]; v != nil {
		s := fmt.Sprintf(" hero IN %s", makeValues(len(v), len(v)))
		wheres = append(wheres, s)
	}
	buf.WriteString(" FROM players")
	if len(wheres) > 0 {
		fmt.Fprintf(buf, " WHERE %s", strings.Join(wheres, ", "))
	}
	fmt.Fprintf(buf, " GROUP BY %s", strings.Join(groups, ", "))
	fmt.Println(buf.String())

	var players []Player
	if err := h.x.Select(&players, buf.String()); err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *hotsContext) UploadReplay(ctx context.Context, r *http.Request, _ httprouter.Params) (interface{}, error) {
	return h.uploadReplay(ctx, r.Body)
}

func (h *hotsContext) uploadReplay(ctx context.Context, r io.Reader) (interface{}, error) {
	dir, err := ioutil.TempDir("", "hots-replay")
	if err != nil {
		return nil, err
	}
	defer func() {
		os.RemoveAll(dir)
	}()
	path := filepath.Join(dir, "replay")
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(file, r)
	file.Close()
	if err != nil {
		return nil, err
	}

	out, err := exec.CommandContext(ctx, "python", "heroprotocol/heroprotocol.py", "--header", "--json", path).CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, string(out))
	}
	var header Header
	if err := json.Unmarshal(out, &header); err != nil {
		return nil, err
	}
	if header.MSignature != "Heroes of the Storm replay\u001b11" {
		return nil, errors.New("invalid signature")
	}

	out, err = exec.CommandContext(ctx, "python", "heroprotocol/heroprotocol.py", "--details", "--json", path).CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, string(out))
	}
	var details Details
	if err := json.Unmarshal(out, &details); err != nil {
		return nil, err
	}

	g := Game{
		Map:   details.MTitle,
		Build: header.MVersion.MBuild,
		Patch: fmt.Sprintf("%d.%d.%d", header.MVersion.MMajor, header.MVersion.MMinor, header.MVersion.MRevision),
	}
	tx, err := h.x.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	if err := tx.Get(&g, `INSERT INTO games (map, patch, build) values ($1, $2, $3) RETURNING id`, g.Map, g.Patch, g.Build); err != nil {
		return nil, errors.Wrap(err, "insert game")
	}

	winners := 0
	playerCount := 0
	players := make(map[int64]Player)
	var data []interface{}
	for i, p := range details.MPlayerList {
		player := Player{
			Build:  g.Build,
			Map:    g.Map,
			Game:   g.ID,
			Hero:   p.MHero,
			Name:   p.MName,
			Team:   p.MTeamID,
			Winner: p.MResult == 1,
		}
		if player.Winner {
			winners++
		}
		playerCount++
		players[int64(i)+1] = player
		data = append(data, player.Build, player.Map, player.Game, player.Hero, player.Name, player.Team, player.Winner)
	}
	if playerCount != 10 || winners != 5 {
		return nil, errors.New("invalid file")
	}
	if _, err := tx.Exec(fmt.Sprintf(
		`INSERT INTO players (build, map, game, hero, name, team, winner) VALUES %s`,
		makeValues(7, len(data))),
		data...); err != nil {
		return nil, errors.Wrap(err, "insert players")
	}

	out, err = exec.CommandContext(ctx, "python", "heroprotocol/heroprotocol.py", "--trackerevents", "--json", path).CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, string(out))
	}
	dec := json.NewDecoder(bytes.NewReader(out))
	data = nil
	talents := make(map[int64]int)
	for {
		var ev Event
		if err := dec.Decode(&ev); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if ev.MEventName == "TalentChosen" {
			pid := ev.MIntData[0].MValue
			p := players[pid]
			talents[pid]++
			talent := &Talent{
				Build:  g.Build,
				Map:    g.Map,
				Game:   g.ID,
				Hero:   p.Hero,
				Tier:   talents[pid],
				Winner: p.Winner,
				Name:   ev.MStringData[0].MValue,
			}
			data = append(data, talent.Build, talent.Map, talent.Game, talent.Hero, talent.Tier, talent.Winner, talent.Name)
		}
	}
	if _, err := tx.Exec(fmt.Sprintf(
		`INSERT INTO talents (build, map, game, hero, tier, winner, name) VALUES %s`,
		makeValues(7, len(data))),
		data...); err != nil {
		return nil, errors.Wrap(err, "insert talents")
	}
	return nil, tx.Commit()
}

type Header struct {
	MDataBuildNum     int64    `json:"m_dataBuildNum"`
	MElapsedGameLoops int64    `json:"m_elapsedGameLoops"`
	MNgdpRootKey      struct{} `json:"m_ngdpRootKey"`
	MSignature        string   `json:"m_signature"`
	MType             int64    `json:"m_type"`
	MUseScaledTime    bool     `json:"m_useScaledTime"`
	MVersion          struct {
		MBaseBuild int64 `json:"m_baseBuild"`
		MBuild     int64 `json:"m_build"`
		MFlags     int64 `json:"m_flags"`
		MMajor     int64 `json:"m_major"`
		MMinor     int64 `json:"m_minor"`
		MRevision  int64 `json:"m_revision"`
	} `json:"m_version"`
}

type Details struct {
	MCacheHandles      []string    `json:"m_cacheHandles"`
	MCampaignIndex     int64       `json:"m_campaignIndex"`
	MDefaultDifficulty int64       `json:"m_defaultDifficulty"`
	MDescription       string      `json:"m_description"`
	MDifficulty        string      `json:"m_difficulty"`
	MGameSpeed         int64       `json:"m_gameSpeed"`
	MImageFilePath     string      `json:"m_imageFilePath"`
	MIsBlizzardMap     bool        `json:"m_isBlizzardMap"`
	MMapFileName       string      `json:"m_mapFileName"`
	MMiniSave          bool        `json:"m_miniSave"`
	MModPaths          interface{} `json:"m_modPaths"`
	MPlayerList        []struct {
		MColor struct {
			MA int64 `json:"m_a"`
			MB int64 `json:"m_b"`
			MG int64 `json:"m_g"`
			MR int64 `json:"m_r"`
		} `json:"m_color"`
		MControl  int64  `json:"m_control"`
		MHandicap int64  `json:"m_handicap"`
		MHero     string `json:"m_hero"`
		MName     string `json:"m_name"`
		MObserve  int64  `json:"m_observe"`
		MRace     string `json:"m_race"`
		MResult   int64  `json:"m_result"`
		MTeamID   int64  `json:"m_teamId"`
		MToon     struct {
			MID        int64  `json:"m_id"`
			MProgramID string `json:"m_programId"`
			MRealm     int64  `json:"m_realm"`
			MRegion    int64  `json:"m_region"`
		} `json:"m_toon"`
		MWorkingSetSlotID int64 `json:"m_workingSetSlotId"`
	} `json:"m_playerList"`
	MRestartAsTransitionMap bool `json:"m_restartAsTransitionMap"`
	MThumbnail              struct {
		MFile string `json:"m_file"`
	} `json:"m_thumbnail"`
	MTimeLocalOffset int64  `json:"m_timeLocalOffset"`
	MTimeUTC         int64  `json:"m_timeUTC"`
	MTitle           string `json:"m_title"`
}

type Event struct {
	Bits       int64  `json:"_bits"`
	Event      string `json:"_event"`
	Eventid    int64  `json:"_eventid"`
	Gameloop   int64  `json:"_gameloop"`
	MEventName string `json:"m_eventName"`
	MFixedData []struct {
		MKey   string `json:"m_key"`
		MValue int64  `json:"m_value"`
	} `json:"m_fixedData"`
	MIntData []struct {
		MKey   string `json:"m_key"`
		MValue int64  `json:"m_value"`
	} `json:"m_intData"`
	MStringData []struct {
		MKey   string `json:"m_key"`
		MValue string `json:"m_value"`
	} `json:"m_stringData"`
	MInstanceList []struct {
		MName   string `json:"m_name"`
		MValues [][]struct {
			MTime  int64 `json:"m_time"`
			MValue int64 `json:"m_value"`
		} `json:"m_values"`
	} `json:"m_instanceList"`
}
