package creator

import (
	"embed"
	"fmt"
	"strings"
)

var banner []string

//go:embed banners/*.txt
var bannerFS embed.FS

// the StringToArt func converts full text input into ASCII art.
func StringToArt(text string, file string) (string, error) {

	cleanFile := strings.TrimPrefix(file, "creator/")
	data, err := bannerFS.ReadFile(cleanFile)
	if err != nil {
		return "", fmt.Errorf("failed to read banner style: %w", err)
	}

	if text == "" {
		return "", nil
	}

	if IsOnlyNewlines(text) {
		return text, nil
	}

	content := string(data)

	content = strings.ReplaceAll(content, "\r", "")

	content = content[1:] //removes the very first character "\n" from the banner file.

	banner = strings.Split(content, "\n\n") //this is used because each ASCII character in the
	// banner file is separated by one empty line,
	// so \n\n separates one full character block from the next.

	lines := strings.Split(text, "\n")

	result := ""

	for _, line := range lines {

		art, err := LineToArt(line)
		if err != nil {
			return "", err
		}

		result += art
	}

	return result, nil
}

// LineToArt converts one line of text into ASCII-art rows.
func LineToArt(line string) (string, error) {

	if line == "" {
		return "\n", nil
	}

	rows := make([]string, 8)

	for _, char := range line {

		letter, err := RuneToArt(char)
		if err != nil {
			return "", err
		}

		parts := strings.Split(letter, "\n")

		for i := 0; i < 8; i++ {
			rows[i] += parts[i]
		}
	}

	return strings.Join(rows, "\n") + "\n", nil
}

// RuneToArt retrieves the ASCII-art representation of a single character from the loaded banner.
func RuneToArt(char rune) (string, error) {

	index := int(char) - 32

	if index < 0 || index >= len(banner) {
		return "", fmt.Errorf("'%v' character '%d' is not supported", string(char), char)
	}

	return banner[index], nil
}
