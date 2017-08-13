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
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	flagInit = flag.Bool("init", false, "drop database before starting")
	initDB   = false
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
			if _, err := db.Exec(fmt.Sprintf("drop database if exists %s; create database %[1]s", dbName)); err != nil {
				panic(err)
			}
		}
	}

	mustMigrate(db)

	wrap := func(f func(context.Context, *http.Request, httprouter.Params) (interface{}, error)) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	router.POST("/api/upload-replay", wrap(h.UploadReplay))
	router.GET("/api/init", wrap(h.Init))
	router.GET("/api/get-winrates", wrap(h.GetWinrates))
	router.GET("/api/get-build-winrates/:hero", wrap(h.GetBuildWinrates))

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

func (h *hotsContext) GetBuildWinrates(ctx context.Context, r *http.Request, ps httprouter.Params) (interface{}, error) {
	args := map[string]string{
		"build": r.FormValue("build"),
		"hero":  ps.ByName("hero"),
		"map":   r.FormValue("map"),
		"mode":  r.FormValue("mode"),
	}
	wrs, err := h.getBuildWinrates(ctx, args)
	if err != nil {
		return nil, err
	}
	prevBuild, err := h.getBuildBefore(ctx, args["build"])
	if err != nil {
		return nil, err
	}
	args["build"] = prevBuild
	prevWrs, err := h.getBuildWinrates(ctx, args)
	if err != nil {
		return nil, errors.Wrapf(err, "fetch previous build: %v", prevBuild)
	}
	return struct {
		Current  map[int]map[string]Total
		Previous map[int]map[string]Total
	}{
		Current:  wrs,
		Previous: prevWrs,
	}, err
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
		return nil, err
	}
	prevBuild, err := h.getBuildBefore(ctx, args["build"])
	if err != nil {
		return nil, err
	}
	args["build"] = prevBuild
	prevWrs, err := h.getWinrates(ctx, args)
	if err != nil {
		return nil, errors.Wrapf(err, "fetch previous build: %v", prevBuild)
	}
	return struct {
		Current  map[string]Total
		Previous map[string]Total
	}{
		Current:  wrs,
		Previous: prevWrs,
	}, err
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
		return nil, err
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

func (h *hotsContext) UploadReplay(ctx context.Context, r *http.Request, _ httprouter.Params) (interface{}, error) {
	return h.uploadReplay(ctx, r.Body)
}

func (h *hotsContext) getBuildBefore(ctx context.Context, id string) (string, error) {
	var build string
	err := h.x.GetContext(ctx, &build, `
		SELECT id
		FROM builds
		WHERE id < $1
		ORDER BY id DESC
		LIMIT 1
	`, id)
	return build, err
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

func (h *hotsContext) unknownName(ctx context.Context, name, typ string) {
	fmt.Printf("UNKNOWN %q\n", name)
	h.db.ExecContext(ctx, `INSERT INTO translations (orig, type) VALUES ($1, $2)`, name, typ)
}

func (h *hotsContext) getBuild(ctx context.Context, header Header, start time.Time) (Build, error) {
	var build Build
	tx, err := h.x.BeginTxx(ctx, nil)
	if err != nil {
		return build, err
	}

	id := header.MVersion.MBuild
	if err = tx.Get(&build, "SELECT * FROM builds WHERE id = $1", id); err == sql.ErrNoRows {
		build.ID = id
		build.Patch = fmt.Sprintf("%d.%d.%d", header.MVersion.MMajor, header.MVersion.MMinor, header.MVersion.MRevision)
		build.Start = start
		build.Finish = start
		_, err = tx.Exec("INSERT INTO builds (id, patch, start, finish) VALUES ($1, $2, $3, $4)", build.ID, build.Patch, build.Start, build.Finish)
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

var uploadLimiter = make(chan struct{}, runtime.NumCPU())

func (h *hotsContext) uploadReplay(ctx context.Context, r io.Reader) (interface{}, error) {
	uploadLimiter <- struct{}{}
	defer func() { <-uploadLimiter }()

	maps, heroes, err := h.getNamesAsMaps(ctx)
	if err != nil {
		return nil, err
	}

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

	out, err = exec.CommandContext(ctx, "python", "heroprotocol/heroprotocol.py", "--init", "--json", path).CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, string(out))
	}
	var init Init
	dec := json.NewDecoder(bytes.NewReader(out))
	var empty interface{}
	if err := dec.Decode(&empty); err != nil {
		return nil, errors.Wrap(err, "decode init ignore")
	}
	if err := dec.Decode(&init); err != nil {
		return nil, errors.Wrap(err, "decode init")
	}

	out, err = exec.CommandContext(ctx, "python", "heroprotocol/heroprotocol.py", "--details", "--json", path).CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, string(out))
	}
	var details Details
	if err := json.Unmarshal(out, &details); err != nil {
		return nil, err
	}

	unknown := false
	if !maps[details.MTitle] {
		unknown = true
		h.unknownName(ctx, details.MTitle, "map")
	}

	desc := init.MSyncLobbyState.MGameDescription
	g := Game{
		Seed: desc.MRandomValue,
		Time: fromFileTimeUTC(details.MTimeUTC),

		Build:  header.MVersion.MBuild,
		Length: header.MElapsedGameLoops / 16,
		Map:    details.MTitle,
		Mode:   modeMap[desc.MGameOptions.MAmmID],
	}
	if g.Mode == 0 {
		return nil, errors.Errorf("unsupported game mode: %d", desc.MGameOptions.MAmmID)
	}

	if _, err := h.getBuild(ctx, header, g.Time); err != nil {
		return nil, err
	}

	tx, err := h.x.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	if err := tx.Get(&g, `INSERT INTO games (map, build, mode, seed, time, length) values ($1, $2, $3, $4, $5, $6) RETURNING id`, g.Map, g.Build, g.Mode, g.Seed, g.Time, g.Length); err != nil {
		return nil, errors.Wrap(err, "insert game")
	}

	winners := 0
	playerCount := 0
	players := make(map[int64]Player)
	var data []interface{}
	for i, p := range details.MPlayerList {
		if !heroes[p.MHero] {
			unknown = true
			h.unknownName(ctx, p.MHero, "hero")
		}
		player := Player{
			Hero:   p.MHero,
			Winner: p.MResult == 1,
		}
		if player.Winner {
			winners++
		}
		playerCount++
		players[int64(i)+1] = player
		data = append(data,
			g.ID,

			p.MHero,
			p.MName,
			p.MTeamID,
			player.Winner,

			g.Build,
			g.Length,
			g.Map,
			g.Mode,
		)
	}
	if unknown {
		return nil, errors.New("unknown strings in replay; skipping")
	}
	if playerCount != 10 || winners != 5 {
		return nil, errors.New("invalid file")
	}
	var pids []int64
	if err := tx.Select(&pids, fmt.Sprintf(
		`INSERT INTO players (game, hero, name, team, winner, build, length, map, mode) VALUES %s RETURNING id`,
		makeValues(9, len(data), 0)),
		data...); err != nil {
		return nil, errors.Wrap(err, "insert players")
	}

	out, err = exec.CommandContext(ctx, "python", "heroprotocol/heroprotocol.py", "--trackerevents", "--json", path).CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, string(out))
	}
	dec = json.NewDecoder(bytes.NewReader(out))
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
			data = append(data,
				g.ID,
				pids[pid-1],

				p.Hero,
				ev.MStringData[0].MValue,
				talents[pid],
				p.Winner,

				g.Build,
				g.Length,
				g.Map,
				g.Mode,
			)
		}
	}
	if _, err := tx.Exec(fmt.Sprintf(
		`INSERT INTO talents (game, player, hero, name, tier, winner, build, length, map, mode) VALUES %s`,
		makeValues(10, len(data), 0)),
		data...); err != nil {
		return nil, errors.Wrap(err, "insert talents")
	}
	return nil, tx.Commit()
}

type Init struct {
	MSyncLobbyState struct {
		MGameDescription struct {
			MCacheHandles      []string `json:"m_cacheHandles"`
			MDefaultAIBuild    int64    `json:"m_defaultAIBuild"`
			MDefaultDifficulty int64    `json:"m_defaultDifficulty"`
			MGameCacheName     string   `json:"m_gameCacheName"`
			MGameOptions       struct {
				MAdvancedSharedControl bool  `json:"m_advancedSharedControl"`
				MAmm                   bool  `json:"m_amm"`
				MAmmID                 int64 `json:"m_ammId"`
				MBattleNet             bool  `json:"m_battleNet"`
				MClientDebugFlags      int64 `json:"m_clientDebugFlags"`
				MCompetitive           bool  `json:"m_competitive"`
				MCooperative           bool  `json:"m_cooperative"`
				MFog                   int64 `json:"m_fog"`
				MHeroDuplicatesAllowed bool  `json:"m_heroDuplicatesAllowed"`
				MLockTeams             bool  `json:"m_lockTeams"`
				MNoVictoryOrDefeat     bool  `json:"m_noVictoryOrDefeat"`
				MObservers             int64 `json:"m_observers"`
				MPractice              bool  `json:"m_practice"`
				MRandomRaces           bool  `json:"m_randomRaces"`
				MTeamsTogether         bool  `json:"m_teamsTogether"`
				MUserDifficulty        int64 `json:"m_userDifficulty"`
			} `json:"m_gameOptions"`
			MGameSpeed           int64  `json:"m_gameSpeed"`
			MGameType            int64  `json:"m_gameType"`
			MHasExtensionMod     bool   `json:"m_hasExtensionMod"`
			MIsBlizzardMap       bool   `json:"m_isBlizzardMap"`
			MIsCoopMode          bool   `json:"m_isCoopMode"`
			MIsPremadeFFA        bool   `json:"m_isPremadeFFA"`
			MMapAuthorName       string `json:"m_mapAuthorName"`
			MMapFileName         string `json:"m_mapFileName"`
			MMapFileSyncChecksum int64  `json:"m_mapFileSyncChecksum"`
			MMapSizeX            int64  `json:"m_mapSizeX"`
			MMapSizeY            int64  `json:"m_mapSizeY"`
			MMaxColors           int64  `json:"m_maxColors"`
			MMaxControls         int64  `json:"m_maxControls"`
			MMaxObservers        int64  `json:"m_maxObservers"`
			MMaxPlayers          int64  `json:"m_maxPlayers"`
			MMaxRaces            int64  `json:"m_maxRaces"`
			MMaxTeams            int64  `json:"m_maxTeams"`
			MMaxUsers            int64  `json:"m_maxUsers"`
			MModFileSyncChecksum int64  `json:"m_modFileSyncChecksum"`
			MRandomValue         int64  `json:"m_randomValue"`
			MSlotDescriptions    []struct {
				MAllowedAIBuilds     []int64       `json:"m_allowedAIBuilds"`
				MAllowedColors       []int64       `json:"m_allowedColors"`
				MAllowedControls     []json.Number `json:"m_allowedControls"`
				MAllowedDifficulty   []int64       `json:"m_allowedDifficulty"`
				MAllowedObserveTypes []int64       `json:"m_allowedObserveTypes"`
				MAllowedRaces        []int64       `json:"m_allowedRaces"`
			} `json:"m_slotDescriptions"`
		} `json:"m_gameDescription"`
		MLobbyState struct {
			MDefaultAIBuild    int64       `json:"m_defaultAIBuild"`
			MDefaultDifficulty int64       `json:"m_defaultDifficulty"`
			MGameDuration      int64       `json:"m_gameDuration"`
			MHostUserID        interface{} `json:"m_hostUserId"`
			MIsSinglePlayer    bool        `json:"m_isSinglePlayer"`
			MMaxObservers      int64       `json:"m_maxObservers"`
			MMaxUsers          int64       `json:"m_maxUsers"`
			MPhase             int64       `json:"m_phase"`
			MPickedMapTag      int64       `json:"m_pickedMapTag"`
			MRandomSeed        int64       `json:"m_randomSeed"`
			MSlots             []struct {
				MAiBuild       int64    `json:"m_aiBuild"`
				MAnnouncerPack string   `json:"m_announcerPack"`
				MArtifacts     []string `json:"m_artifacts"`
				MBanner        string   `json:"m_banner"`
				MColorPref     struct {
					MColor int64 `json:"m_color"`
				} `json:"m_colorPref"`
				MControl           int64         `json:"m_control"`
				MDifficulty        int64         `json:"m_difficulty"`
				MHandicap          int64         `json:"m_handicap"`
				MHasSilencePenalty bool          `json:"m_hasSilencePenalty"`
				MHero              string        `json:"m_hero"`
				MHeroMasteryTiers  []interface{} `json:"m_heroMasteryTiers"`
				MLogoIndex         int64         `json:"m_logoIndex"`
				MMount             string        `json:"m_mount"`
				MObserve           int64         `json:"m_observe"`
				MRacePref          struct {
					MRace interface{} `json:"m_race"`
				} `json:"m_racePref"`
				MRewards            []int64     `json:"m_rewards"`
				MSkin               string      `json:"m_skin"`
				MSpray              string      `json:"m_spray"`
				MTandemLeaderUserID interface{} `json:"m_tandemLeaderUserId"`
				MTeamID             int64       `json:"m_teamId"`
				MToonHandle         string      `json:"m_toonHandle"`
				MUserID             int64       `json:"m_userId"`
				MVoiceLine          string      `json:"m_voiceLine"`
				MWorkingSetSlotID   int64       `json:"m_workingSetSlotId"`
			} `json:"m_slots"`
		} `json:"m_lobbyState"`
		MUserInitialData []struct {
			MBanner             string      `json:"m_banner"`
			MClanLogo           interface{} `json:"m_clanLogo"`
			MClanTag            string      `json:"m_clanTag"`
			MCombinedRaceLevels int64       `json:"m_combinedRaceLevels"`
			MCustomInterface    bool        `json:"m_customInterface"`
			MExamine            bool        `json:"m_examine"`
			MHero               string      `json:"m_hero"`
			MHighestLeague      int64       `json:"m_highestLeague"`
			MMount              string      `json:"m_mount"`
			MName               string      `json:"m_name"`
			MObserve            int64       `json:"m_observe"`
			MRacePreference     struct {
				MRace interface{} `json:"m_race"`
			} `json:"m_racePreference"`
			MRandomSeed     int64  `json:"m_randomSeed"`
			MSkin           string `json:"m_skin"`
			MSpray          string `json:"m_spray"`
			MTeamPreference struct {
				MTeam interface{} `json:"m_team"`
			} `json:"m_teamPreference"`
			MTestAuto   bool   `json:"m_testAuto"`
			MTestMap    bool   `json:"m_testMap"`
			MTestType   int64  `json:"m_testType"`
			MToonHandle string `json:"m_toonHandle"`
		} `json:"m_userInitialData"`
	} `json:"m_syncLobbyState"`
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

func fromFileTimeUTC(i int64) time.Time {
	const (
		ticksPerDay = time.Hour * 24 / (time.Nanosecond * 100)

		daysPerYear        = 365
		daysPer4Years      = daysPerYear*4 + 1
		daysPer100Years    = daysPer4Years*25 - 1
		daysFrom1601To1970 = daysPer100Years*3 + daysPer4Years*17 + daysPerYear
		offset             = int64(ticksPerDay * daysFrom1601To1970)
	)
	return time.Unix((i-offset)/1e7, 0)
}
