package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"log"
)

const statsfile string = ".config/watson/state"

type Stats struct {
	Project string `json:"project"`
	Start   int64  `json:"start"`
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if err != nil {
		return false
	}

	return !info.IsDir()
}

func main() {
	statsfile := filepath.Join(os.Getenv("HOME"), statsfile)

	if !fileExists(statsfile) {
		return
	}

	data, err := os.ReadFile(statsfile)
	if err != nil {
		log.Fatalf("Could not read watson stats file: %w", err)
	}

	var stats Stats
	if err = json.Unmarshal(data, &stats); err != nil {
		log.Fatalf("Could not unmarshal JSON: %w", err)
	}

	start := time.Unix(stats.Start, 0)
	elapsed := time.Since(start)

	fmt.Printf("%.02fh\n", elapsed.Hours())
}
