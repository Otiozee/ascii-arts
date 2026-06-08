> This project is designed as both a standalone CLI tool and a reusable Go package.

# Overview

ASCII-ART is a text-to-ASCII-art conversion tool written in Go. It reads user input from the command line and renders the text using selectable ASCII-art banner styles such as `standard`, `shadow`, and `thinkertoy`.

The application processes user input through a structured rendering pipeline: it normalizes escaped newlines, validates banner selection, loads banner files, converts characters into ASCII-art blocks, and finally renders the generated output to the terminal.

This project demonstrates clean Go architecture, file handling and resource compilation using `embed` and `os`, string processing with `strings`, rune manipulation, and modular program design. It also handles several edge cases such as:

* empty input,
* trailing newlines,
* leading newlines,
* multiple blank lines,
* unsupported ASCII characters,
* and escaped `\n` sequences.

---

## Features

* `main`: controls the entire application flow via a deliberate workflow as shown below:

```text
main()
 ├── NormalizeInput()
 ├── IsOnlyNewlines()
 ├── GetBannerPath()
 └── Render()
      └── StringToArt()
           ├── LineToArt()
           └── RuneToArt()
```

* `NormalizeInput`

  * Converts escaped newline characters (`\n`) into actual new lines.

* `IsOnlyNewlines`

  * Detects inputs made entirely of newline characters.

* `GetBannerPath`

  * Maps banner names (`standard`, `shadow`, `thinkertoy`) to their corresponding banner files.

* `Render`

  * Processes each input line and prints generated ASCII art to the terminal.

* `StringToArt`

  * Converts full text input into ASCII art.

* `LineToArt`

  * Converts one line of text into ASCII-art rows.

* `RuneToArt`

  * Retrieves the ASCII-art representation of a single character from the loaded banner.
 
* Reusable Go package (`creator`)

---

## Pipeline Architecture

```text
User Input (from CLI Arguments)
        │
        │
main()
        │
        ├── parses arguments
        ├── selects banner
        └── reads source text
        │
        │
NormalizeInput()
        │
        ├── converts "\n" -> actual newline
        └── prepares multiline input
        │
        │
IsOnlyNewlines()
        │
        ├── handles newline-only edge cases
        └── prevents extra blank rendering
        │
        │
GetBannerPath()
        │
        ├── validates banner name
        └── returns banner file path
        │
        │
Render()
        │
        ├── splits text into lines
        ├── calls StringToArt()
        └── prints final ASCII art
        │
        │
StringToArt()
        │
        ├── loads banner file
        ├── extracts ASCII blocks
        ├── loops through lines
        └── converts text into ASCII art
        │
        │
LineToArt()
        │
        ├── loops through characters
        ├── gets ASCII pattern per rune
        └── merges rows horizontally
        │
        │
RuneToArt()
        │
        └── returns ASCII-art block on terminal
```

---

## Program Structure

```text
ascii-art/
│
├── creator/
│   ├── banners/
│   │   ├── standard.txt
│   │   ├── shadow.txt
│   │   └── thinkertoy.txt
│   │
│   ├── converter.go
│   ├── render.go
│   ├── normalize.go
│   └── banner.go
│
├── converter_test.go
├── go.mod
├── main.go
└── shell_test.sh
```

---

## Language & Tools

<div align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go Programming Language" width="180"/>
  <br/>
  <strong>Go (Golang)</strong> - Simple, fast, and reliable.
</div>

* Built entirely with Go’s standard library
* No external dependencies
* Uses packages such as:

  * `os`
  * `fmt`
  * `strings`
  * `log`
  * `embed`

---

## Installation

Clone the repository:

```bash
git clone https://github.com/Otiozee/ascii-arts.git
```

Move into the project directory:

```bash
cd ascii-art
```

---

## Testing

The project includes unit tests covering:

* empty inputs,
* newline handling,
* mixed content,
* unsupported characters,
* multiline rendering,
* rune conversion,
* and utility functions.

Run tests:

```bash
go test
```

Verbose mode:

```bash
go test -v
```

---

## Usage

### Default Banner

```bash
go run . "Hello"
```

### Custom Banner

```bash
go run . shadow "Hello"
```

### Multiline Input

```bash
go run . "Hello\nWorld"
```

## Examples

### Standard

```bash
go run . "ASCII"
```

### Shadow

```bash
go run . shadow "ASCII"
```

### Thinkertoy

```bash
go run . thinkertoy "ASCII"
```

---

### Run muiltiple examples from shell file
```bash
chmod +x shell_test.sh
```
```bash
./shell_test.sh
```

## Edge Cases Handled

* Empty string
* Single newline
* Multiple newlines
* Leading blank lines
* Trailing blank lines
* Multiple blank lines between text
* Unsupported Unicode characters
* Escaped newline input (`\n`)
* Spaces-only input

---

## Using as a Package

This project can be imported and reused in other Go applications.
### Public API
These functions can be imported into other Go projects.

```go
creator.NormalizeInput(...)
creator.IsOnlyNewline(...)
creator.GetBannerPath(...)
creator.Render(...)
creator.StringToArt(...)
creator.LineToArt(...)
creator.RuneToArt(...)
```

```go
import "github.com/otiozee/ascii-arts/creator"
```

Example:

```go
package main

import (
"github.com/otiozee/ascii-arts/creator"
"log"
"os"
)

func main() {
var banner, source string

	switch len(os.Args) {
	case 1:
		banner = "standard"
		source = os.Args[1]
	case 2:
		banner = os.Args[1]
		source = os.Args[2]
	default:
		log.Fatal("Usage: go run . [BANNER] [STRING]\nExample: go run . \"Hello\" \"shadow\"\n")
	}

	source = creator.NormalizeInput(source)

	if source == "" {
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
```
---

## Program Team

* [zeotokpa](https://github.com/otiozee)
* [aonminyi](https://github.com/andrewokala)
