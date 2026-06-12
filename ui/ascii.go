package ui

import (
	_ "embed"
	"fmt"
	"strings"
)

const (
	reset = "\033[0m"
	cyan  = "\033[36m"
)

//go:embed assets/tux.txt
var asciiArt string

func FormatASCII() string {
	lines := strings.Split(asciiArt, "\n")

	var b strings.Builder

	for _, line := range lines {
		fmt.Fprintf(&b, "%s%s%s\n", cyan, line, reset)
	}

	return b.String()
}
