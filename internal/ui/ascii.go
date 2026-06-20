package ui

import (
	_ "embed"
	"strings"
)

//go:embed assets/tux.txt
var asciiArt string

func FormatASCII() string {
	c1 := "\x1b[90m" // gray
	c2 := "\x1b[97m" // white
	c3 := "\x1b[33m" // yellow

	reset := "\x1b[0m"

	asciiArt = strings.ReplaceAll(asciiArt, "$1", c1)
	asciiArt = strings.ReplaceAll(asciiArt, "$2", c2)
	asciiArt = strings.ReplaceAll(asciiArt, "$3", c3)

	return asciiArt + reset
}
