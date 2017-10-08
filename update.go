package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func (h *hotsContext) update() error {
	for {
		start := time.Now()
		ctx, _ := context.WithTimeout(context.Background(), time.Minute*20)
		n, err := h.nextBlock(ctx)
		if err != nil {
			log.Println("update error:", err)
		} else {
			log.Println("updated", n)
		}
		log.Println("update took", time.Since(start))
		if n == 0 {
			log.Println("sleeping")
			time.Sleep(time.Minute)
		}
	}
}

func (h *hotsContext) nextBlock(ctx context.Context) (int, error) {
	var lastID int
	const configLastID = "last_id"
	if err := h.x.GetContext(ctx, &lastID, `
		SELECT i
		FROM config
		WHERE key = $1
	`, configLastID); err == sql.ErrNoRows {
		// lastID = 0
	} else if err != nil {
		return 0, err
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

	const HotsApi = "https://hotsapi.net/api/v1"
	url := fmt.Sprintf(HotsApi+"/replays?min_id=%d", lastID)
	b, err := httpGet(ctx, url)
	if err != nil {
		return 0, err
	}
	var replays Replays
	if err := json.Unmarshal(b, &replays); err != nil {
		return 0, errors.Wrapf(err, "JSON decoding replays, url: %s, body: %s", url, b)
	}
	ch := make(chan *Replay, 5)
	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		defer close(ch)
		for _, r := range replays {
			b, err := httpGet(gCtx, fmt.Sprintf(HotsApi+"/replays/%d", r.ID))
			if err != nil {
				return err
			}
			var replay Replay
			if err := json.Unmarshal(b, &replay); err != nil {
				return errors.Wrap(err, "decoding replay")
			}
			select {
			case ch <- &replay:
			case <-gCtx.Done():
				return gCtx.Err()
			}
		}
		return nil
	})
	n := 0
	group.Go(func() error {
		for r := range ch {
			if err := h.getReplay(gCtx, r); err != nil {
				return err
			}
			lastID = r.ID
			n++
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		return n, err
	}
	_, err = h.db.ExecContext(ctx, `UPSERT INTO config (key, i) VALUES ($1, $2)`, configLastID, lastID)
	return n, err
}

func (h *hotsContext) getReplay(ctx context.Context, r *Replay) error {
	mode, ok := gameModes[r.GameType]
	if !ok {
		log.Printf("unknown game type: %s", r.GameType)
		return nil
	}
	if _, err := h.addBuild(ctx, r); err != nil {
		return err
	}
	if _, err := h.db.Exec(`UPSERT INTO maps (name) VALUES ($1) RETURNING NOTHING`, r.GameMap); err != nil {
		return err
	}
	return h.txn(ctx, func(txn *sqlx.Tx) error {
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
		return nil
	})
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

func (h *hotsContext) addBuild(ctx context.Context, replay *Replay) (Build, error) {
	var build Build
	err := h.txn(ctx, func(txn *sqlx.Tx) error {
		id := replay.GameVersion
		start := time.Time(replay.GameDate)
		var err error
		if err = txn.Get(&build, "SELECT * FROM builds WHERE id = $1", id); err == sql.ErrNoRows {
			build.ID = id
			build.Start = start
			build.Finish = start
			_, err = txn.Exec("INSERT INTO builds (id, start, finish) VALUES ($1, $2, $3)", build.ID, build.Start, build.Finish)
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
				_, err = txn.Exec("UPDATE builds SET start = $1, finish = $2 WHERE id = $3", build.Start, build.Finish, id)
			}
		}
		return err
	})
	return build, err
}
