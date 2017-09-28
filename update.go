package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func update(addr string) error {
	url := fmt.Sprintf("%s/api/next-block", addr)
	log.Printf("running update for: %s", addr)

	for {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		resp.Body.Close()
		if resp.StatusCode != 200 {
			return errors.Errorf("return: %s", resp.Status)
		}
		log.Println("update took", time.Since(start))
	}
}
