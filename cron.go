package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// k8s cron jobs on GKE don't work, so do it here instead.
func (h *hotsContext) cronLoop() error {
	start := time.Now()
	fmt.Println("starting cron")
	defer func() { fmt.Println("cron finished in", time.Since(start)) }()
	return h.cron()
}

func (h *hotsContext) cron() error {
	if err := h.updateInit(context.Background()); err != nil {
		return errors.Wrap(err, "update init")
	}
	done := make(map[string]bool)
	update := func(u string, upsert bool) error {
		url, err := url.Parse(u)
		if err != nil {
			return err
		}
		if done[u] {
			log.Printf("cron: ignoring dup: %s", u)
			return nil
		}
		done[u] = true
		var start time.Time
		req := &http.Request{URL: url}
		if err := retry(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
			defer cancel()
			var res interface{}
			start = time.Now()
			switch url.Path {
			case "/api/get-winrates":
				res, err = h.GetWinrates(ctx, req)
			case "/api/get-hero-data":
				res, err = h.GetHero(ctx, req)
			case "/api/get-build-winrates":
				res, err = h.GetBuildWinrates(ctx, req)
			case "/api/get-compare-hero":
				res, err = h.GetCompareHero(ctx, req)
			default:
				log.Printf("cron: unknown path: %s", u)
				return nil
			}
			if err != nil {
				return errors.Wrap(err, u)
			}
			data, gzip, err := resultToBytes(res)
			if err != nil {
				return err
			}
			if upsert {
				_, err = h.db.Exec(`UPSERT INTO cache (id, last_hit, data, gzip) VALUES ($1, now(), $2, $3)`,
					u, data, gzip,
				)

			} else {
				_, err = h.db.Exec(`UPDATE cache SET data = $1, gzip = $2, until = now() WHERE id = $3`,
					data, gzip, u,
				)
			}
			return err
		}); err != nil {
			return errors.Wrap(err, "update cache")
		}
		fmt.Println("recached", url, "in", time.Since(start))
		return nil
	}
	{
		init := h.getInit()
		v := make(url.Values)
		v.Add("build", init.Builds[0].ID)
		var u url.URL
		for _, h := range init.Heroes {
			v.Set("hero", h.Name)
			for _, s := range []string{
				"/api/get-build-winrates",
				"/api/get-compare-hero",
				"/api/get-hero-data",
			} {
				u.Path = s
				u.RawQuery = v.Encode()
				log.Printf("force update: %s", u.String())
				if err := update(u.String(), true); err != nil {
					return err
				}
			}
		}
	}
	if err := retry(func() error {
		_, err := h.db.Exec(`DELETE FROM cache WHERE last_hit < (now() - '48h'::interval)`)
		return err
	}); err != nil {
		return errors.Wrap(err, "empty cache")
	}
	var urls []string
	if err := retry(func() error {
		urls = urls[:0]
		return h.x.Select(&urls, `SELECT id FROM cache`)
	}); err != nil {
		return errors.Wrap(err, "fetch urls")
	}
	for _, u := range urls {
		if err := update(u, false); err != nil {
			return err
		}
	}
	return nil
}
