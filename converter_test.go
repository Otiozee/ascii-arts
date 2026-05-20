package main

import (
	"ascii-arts/creator"
	"strings"
	"testing"
)

const banner = "creator/banners/standard.txt"

func assertEqual(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\nGot:\n%q\n\nExpected:\n%q\n", got, want)
	}
}

// =================================================================
// Test stringToArt
// =================================================================
func Test_StringToArt(t *testing.T) {

	t.Run("Empty string", func(t *testing.T) {
		got, err := creator.StringToArt("", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertEqual(t, got, "")
	})

	t.Run("Single newline", func(t *testing.T) {
		got, err := creator.StringToArt("\n", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertEqual(t, got, "\n")
	})

	t.Run("Multiple newlines", func(t *testing.T) {
		got, err := creator.StringToArt("\n\n", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertEqual(t, got, "\n\n")
	})

	t.Run("Only spaces", func(t *testing.T) {
		got, err := creator.StringToArt("   ", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected output")
		}
	})

	t.Run("Lowercase", func(t *testing.T) {
		got, err := creator.StringToArt("hello", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected output")
		}
	})

	t.Run("Uppercase", func(t *testing.T) {
		got, err := creator.StringToArt("HELLO", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected output")
		}
	})

	t.Run("Mixed case", func(t *testing.T) {
		got, err := creator.StringToArt("HeLlO", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected output")
		}
	})

	t.Run("Numbers", func(t *testing.T) {
		got, err := creator.StringToArt("1234567890", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected output")
		}
	})

	t.Run("Symbols", func(t *testing.T) {
		got, err := creator.StringToArt("!@#$%^&*()", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected output")
		}
	})

	t.Run("Mixed content", func(t *testing.T) {
		got, err := creator.StringToArt("Hello 123!", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected output")
		}
	})

	t.Run("Text with newline", func(t *testing.T) {
		got, err := creator.StringToArt("Hello\nWorld", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if strings.Count(got, "\n") < 16 {
			t.Fatalf("expected multiline ASCII art")
		}
	})

	t.Run("Multiple blank lines", func(t *testing.T) {
		got, err := creator.StringToArt("Hi\n\nThere", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !strings.Contains(got, "\n\n") {
			t.Fatalf("expected blank line")
		}
	})

	t.Run("Leading newlines", func(t *testing.T) {
		got, err := creator.StringToArt("\n\nHello", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !strings.HasPrefix(got, "\n\n") {
			t.Fatalf("expected leading newlines")
		}
	})

	t.Run("Trailing newlines", func(t *testing.T) {
		got, err := creator.StringToArt("Hello\n\n", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !strings.HasSuffix(got, "\n\n") {
			t.Fatalf("expected trailing newlines")
		}
	})

	t.Run("Unsupported character", func(t *testing.T) {
		_, err := creator.StringToArt("Hello☺", banner)

		if err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("Long sentence", func(t *testing.T) {
		got, err := creator.StringToArt("This is a long ASCII ART test 123!", banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected output")
		}
	})
}

//=================================================================
//Test LineToArt
//=================================================================

func Test_LineToArt(t *testing.T) {

	t.Run("Empty line", func(t *testing.T) {
		got, err := creator.LineToArt("")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertEqual(t, got, "\n")
	})

	t.Run("Single character", func(t *testing.T) {
		got, err := creator.LineToArt("A")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if strings.Count(got, "\n") != 8 {
			t.Fatalf("expected 8 lines")
		}
	})

	t.Run("Word", func(t *testing.T) {
		got, err := creator.LineToArt("Hello")

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected output")
		}
	})

	t.Run("Unsupported character", func(t *testing.T) {
		_, err := creator.LineToArt("☺")

		if err == nil {
			t.Fatalf("expected error")
		}
	})
}

//=================================================================
//Test RuneToArt
//=================================================================

func Test_RuneToArt(t *testing.T) {

	data, _ := creator.StringToArt("A", banner)
	_ = data

	t.Run("Valid rune", func(t *testing.T) {
		got, err := creator.RuneToArt('A')

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected ASCII art")
		}
	})

	t.Run("Space rune", func(t *testing.T) {
		got, err := creator.RuneToArt(' ')

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got == "" {
			t.Fatalf("expected ASCII art")
		}
	})

	t.Run("Invalid rune", func(t *testing.T) {
		_, err := creator.RuneToArt('☺')

		if err == nil {
			t.Fatalf("expected error")
		}
	})
}

// =================================================================
// Test NormalizeInput
// =================================================================
func Test_NormalizeInput(t *testing.T) {

	got := creator.NormalizeInput("Hello\\nWorld")

	assertEqual(t, got, "Hello\nWorld")
}

// =================================================================
// Test IsOnlyNewlines
// =================================================================
func Test_IsOnlyNewlines(t *testing.T) {

	t.Run("Only newlines", func(t *testing.T) {
		if !creator.IsOnlyNewlines("\n\n\n") {
			t.Fatalf("expected true")
		}
	})

	t.Run("With text", func(t *testing.T) {
		if creator.IsOnlyNewlines("\nA\n") {
			t.Fatalf("expected false")
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		if !creator.IsOnlyNewlines("") {
			t.Fatalf("expected true")
		}
	})
}

// =================================================================
// Test GetBannerPath
// =================================================================
func Test_GetBannerPath(t *testing.T) {

	t.Run("Standard", func(t *testing.T) {
		got := creator.GetBannerPath("standard")

		assertEqual(t, got, "creator/banners/standard.txt")
	})

	t.Run("Shadow", func(t *testing.T) {
		got := creator.GetBannerPath("shadow")

		assertEqual(t, got, "creator/banners/shadow.txt")
	})

	t.Run("Thinkertoy", func(t *testing.T) {
		got := creator.GetBannerPath("thinkertoy")

		assertEqual(t, got, "creator/banners/thinkertoy.txt")
	})
}
