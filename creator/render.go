package creator

import (
	"fmt"
	"log"
	"strings"
)

// Render processes each input line and prints generated ASCII art to the terminal.
func Render(input, bannerPath string) {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if line == "" {
			if i != 0 {
				fmt.Println()
			}
			continue
		}

		res, err := StringToArt(line, bannerPath)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(res)

	}
}
