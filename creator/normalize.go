package creator

import "strings"

// NormalizeInput converts escaped newline characters (`\n`) into actual new lines.
func NormalizeInput(s string) string {
	return strings.ReplaceAll(s, "\\n", "\n")
}

// IsOnlyNewlines detects inputs made entirely of newline characters.
func IsOnlyNewlines(s string) bool {
	for _, ch := range s {
		if ch != '\n' {
			return false
		}
	}
	return true
}
