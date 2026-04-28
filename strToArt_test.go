package main

import (
	"ascii-art/workers"
	"strings"
	"testing"
)

const banner = "workers/banners/standard.txt"

func assertEqual(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("\nGot:\n%q\n\nExpected:\n%q\n", got, want)
	}
}

func Test_StringToArt_EdgeCases(t *testing.T) {

	t.Run("Empty string", func(t *testing.T) {
		got, _ := workers.StringToArt("", banner)
		assertEqual(t, got, "")
	})

	t.Run("Single newline", func(t *testing.T) {
		got, _ := workers.StringToArt("\n", banner)
		assertEqual(t, got, "\n")
	})

	t.Run("Multiple newlines", func(t *testing.T) {
		got, _ := workers.StringToArt("\n\n", banner)
		assertEqual(t, got, "\n\n")
	})

	t.Run("Only spaces", func(t *testing.T) {
		got, _ := workers.StringToArt("   ", banner)

		lines := strings.Split(got, "\n")
		if len(lines) < 8 {
			t.Errorf("Expected at least 8 lines, got %d", len(lines))
		}
	})

	t.Run("Lowercase", func(t *testing.T) {
		got, _ := workers.StringToArt("hello", banner)
		if got == "" {
			t.Errorf("Expected non-empty output")
		}
	})

	t.Run("Uppercase", func(t *testing.T) {
		got, _ := workers.StringToArt("HELLO", banner)
		if got == "" {
			t.Errorf("Expected non-empty output")
		}
	})

	t.Run("Mixed case", func(t *testing.T) {
		got, _ := workers.StringToArt("HeLlO", banner)
		if got == "" {
			t.Errorf("Expected non-empty output")
		}
	})

	t.Run("Numbers", func(t *testing.T) {
		got, _ := workers.StringToArt("123", banner)
		if got == "" {
			t.Errorf("Expected non-empty output")
		}
	})

	t.Run("Symbols", func(t *testing.T) {
		got, _ := workers.StringToArt("!@#$%^&*", banner)
		if got == "" {
			t.Errorf("Expected non-empty output")
		}
	})

	t.Run("Mixed content", func(t *testing.T) {
		got, _ := workers.StringToArt("Hello 123!", banner)
		if got == "" {
			t.Errorf("Expected non-empty output")
		}
	})

	t.Run("With newline split", func(t *testing.T) {
		got, _ := workers.StringToArt("Hello\nWorld", banner)

		if strings.Count(got, "\n") < 16 {
			t.Errorf("Expected multi-line output, got:\n%s", got)
		}
	})

	t.Run("Multiple empty lines between text", func(t *testing.T) {
		got, _ := workers.StringToArt("Hi\n\nThere", banner)

		if !strings.Contains(got, "\n\n") {
			t.Errorf("Expected blank line separation")
		}
	})

	t.Run("Unsupported character", func(t *testing.T) {
		_, err := workers.StringToArt("Hello☺", banner)

		if err == nil {
			t.Errorf("Expected error for non-ASCII character")
		}
	})
}
