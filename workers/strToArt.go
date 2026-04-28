package workers

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var letters []string

func StringToArt(source string, bannerFilePath string) (string, error) {
	bannerBytes, err := os.ReadFile(bannerFilePath)
	if err != nil {
		return "", err
	}

	bannerString := string(bannerBytes)

	bannerString = strings.ReplaceAll(bannerString, "\r", "")

	bannerString = bannerString[1:]

	letters = strings.Split(bannerString, "\n\n")

	lines := strings.Split(source, "\n")
	for i := range lines {
		lines[i], err = LineToArt(lines[i])
		if err != nil {
			return "", err
		}
	}

	if lines[0] == "\n" {
		lines[0] = ""
	}

	return strings.Join(lines, ""), nil
}

func LineToArt(source string) (string, error) {
	if source == "" {
		return "\n", nil
	}

	res := make([]string, 8)

	for _, r := range []rune(source) {
		runeArt, err := RuneToArt(r)
		if err != nil {
			return "", err
		}
		art := strings.Split(runeArt, "\n")
		for i := 0; i < 8; i++ {
			res[i] += art[i]
		}
	}

	return strings.Join(res, "\n") + "\n", nil
}

func RuneToArt(source rune) (string, error) {
	if source == '\n' {
		return "\n", nil
	}
	if int(source-' ') > len(letters)-1 {
		return "", errors.New("ERROR: '" + string(source) + "' (code " + strconv.Itoa(int(source)) + ") is not an ASCII character")
	}
	return letters[source-' '], nil
}
