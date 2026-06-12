package ui

import (
	"fmt"
	"regexp"
	"strings"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func visibleLen(s string) int {
	return len(ansiRegex.ReplaceAllString(s, ""))
}

func FormatOutput(ascii string, header string, info []string) string {
	ascii = strings.TrimRight(ascii, "\n")
	asciiLines := strings.Split(ascii, "\n")

	// calculate ASCII column width (ANSI-safe)
	width := 0
	for _, line := range asciiLines {
		if l := visibleLen(line); l > width {
			width = l
		}
	}

	// build output
	var b strings.Builder

	rows := max(len(asciiLines), len(info)+1) // +1 for header row

	for i := range rows {
		var left, right string

		// LEFT COLUMN (ASCII)
		if i < len(asciiLines) {
			left = asciiLines[i]
		}

		// RIGHT COLUMN
		if i == 0 {
			right = header
		} else if i-1 < len(info) {
			right = info[i-1]
		}

		pad := width + len(left) - visibleLen(left)
		fmt.Fprintf(&b, "%-*s  %s\n", pad, left, right)
	}

	return b.String()
}
