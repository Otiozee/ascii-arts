package workers

import "strings"

func NormalizeInput(s string) string {
	return strings.ReplaceAll(s, "\\n", "\n")
}

func IsOnlyNewlines(s string) bool {
	for _, ch := range s {
		if ch != '\n' {
			return false
		}
	}
	return true
}
