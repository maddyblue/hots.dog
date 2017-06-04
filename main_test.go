package main

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestUploadReplay(t *testing.T) {
	var h hotsContext

	ctx := context.Background()
	f, err := os.Open("8518_R1_1458527788.StormReplay")
	if err != nil {
		t.Fatal(err)
	}
	v, err := h.uploadReplay(ctx, f)
	f.Close()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)
}
