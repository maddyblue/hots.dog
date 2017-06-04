package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func mustInitDevData(addr string) {
	matches, err := filepath.Glob("replays/*.StormReplay")
	if err != nil {
		panic(err)
	}
	for _, m := range matches {
		replay, err := os.Open(m)
		if err != nil {
			panic(err)
		}
		resp, err := http.Post("http://"+addr+"/api/upload-replay", "", replay)
		if err != nil || resp.StatusCode != 200 {
			io.Copy(os.Stderr, resp.Body)
			panic(err)
		}
		fmt.Println(resp.Status)
	}
}
