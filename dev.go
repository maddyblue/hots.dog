package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func mustInitDevData(addr string, db *sql.DB) {
	start := time.Now()
	defer func() { fmt.Println("DEV TOOK", time.Since(start)) }()
	const name = "data.sql"
	if f, err := ioutil.ReadFile(name); err == nil {
		if _, err := db.Exec(string(f)); err != nil {
			panic(err)
		}
		fmt.Println("imported", name)
		return
	}
	matches, err := filepath.Glob("replays/*.StormReplay")
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, m := range matches {
		wg.Add(1)
		go func(m string) {
			defer wg.Done()
			replay, err := os.Open(m)
			if err != nil {
				panic(err)
			}
			resp, err := http.Post("http://"+addr+"/api/upload-replay", "", replay)
			if err != nil || resp.StatusCode != 200 {
				io.Copy(os.Stderr, resp.Body)
				log.Print(err)
			}
			fmt.Println(resp.Status)
		}(m)
	}
	wg.Wait()
}
