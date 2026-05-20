package main

import (
	"ascii-arts/creator"
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
		log.Fatal("Usage: go run . [BANNER] [STRING]\nExample: go run . \"Hello\" \"shadow\"\n")
	}

	source = creator.NormalizeInput(source)
	
	if source == "" {
		fmt.Println()
		return
	}

	if creator.IsOnlyNewlines(source) {
		for range source {
			fmt.Println()
		}
		return
	}

	bannerPath := creator.GetBannerPath(banner)

	creator.Render(source, bannerPath)
}
