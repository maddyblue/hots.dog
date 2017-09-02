package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	resp, err := http.Get("http://" + addr + "/api/next-block")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil || resp.StatusCode != 200 {
		log.Printf("DEV IMPORT ERROR: %v", err)
	}
}
