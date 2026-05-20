# Overview

ASCII-ART is a text-to-ASCII-art conversion tool written in Go. It reads user input from the command line and renders the text using selectable ASCII-art banner styles such as `standard`, `shadow`, and `thinkertoy`.

The application processes user input through a structured rendering pipeline: it normalizes escaped newlines, validates banner selection, loads banner files, converts characters into ASCII-art blocks, and finally renders the generated output to the terminal.

This project demonstrates clean Go architecture, file handling with `os`, string processing with `strings`, rune manipulation, and modular program design. It also handles several edge cases such as:

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
└── main.go
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

---

## Installation

Clone the repository:

```bash
git clone https://acad.learn2earn.ng/git/zeotokpa/ascii-art.git
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

## Program Team

* zeotokpa
* aonminyi
* umusa
