package creator

import (
	"fmt"
	"os"
	"strings"
)

var banner []string

func StringToArt(text string, file string) (string, error) {

	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	if text == "" {
		return "", nil
	}

	if IsOnlyNewlines(text) {
		return text, nil
	}

	content := string(data)

	content = strings.ReplaceAll(content, "\r", "")

	banner = strings.Split(content, "\n\n")

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

func RuneToArt(char rune) (string, error) {

	index := int(char) - 32

	if index < 0 || index >= len(banner) {
		return "", fmt.Errorf("'%v' character '%d' is not supported", string(char), char)
	}

	return banner[index], nil
}
