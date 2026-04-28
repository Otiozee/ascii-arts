package workers

import (
	"fmt"
	"log"
	"strings"
)

func Render(source, bannerPath string) {
	lines := strings.Split(source, "\n")

	for i, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}

		res, err := StringToArt(line, bannerPath)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(res)

		if i < len(lines)-1 {
			fmt.Println()
		}
	}
}
