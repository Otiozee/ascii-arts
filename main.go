package main

import (
	"ascii-art/workers"
	"fmt"
	"log"
	"os"
)

func main() {

	var banner, source string

	switch len(os.Args) {
	case 2:
		banner = "standard"
		source = os.Args[1]
	case 3:
		banner = os.Args[1]
		source = os.Args[2]
	default:
		log.Fatal("Usage: go run . [BANNER] [STRING]")
	}

	source = workers.NormalizeInput(source)

	if source == "" {
		fmt.Println()
		return
	}

	if workers.IsOnlyNewlines(source) {
		for range source {
			fmt.Println()
		}
		return
	}

	bannerPath := workers.GetBannerPath(banner)

	workers.Render(source, bannerPath)
}
