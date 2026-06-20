package main

import (
	"fmt"
	"github.com/nxkst/blinkfetch/internal/information"
	"github.com/nxkst/blinkfetch/internal/ui"
)

func main() {
	ASCIIArt := ui.FormatASCII()
	yellow := "\x1b[33m"

	header := fmt.Sprintf(
		"\033[0m%s@%s",
		information.Username, information.Hostname,
	)

	info := []string{
		"──────────────",
		fmt.Sprintf(yellow+"os:\033[0m %s",
			information.Distro,
		),
		fmt.Sprintf(yellow+"kernel:\033[0m %s",
			information.Kernel,
		),
		fmt.Sprintf(yellow+"uptime:\033[0m %s",
			information.FormattedUptime,
		),
		fmt.Sprintf(yellow+"de:\033[0m %s",
			information.Desktop,
		),
		fmt.Sprintf(yellow+"shell:\033[0m %s",
			information.Shell,
		),
		fmt.Sprintf(yellow+"memory:\033[0m %s MiB / %s MiB (%s%%)",
			information.UsedMemory,
			information.TotalMemory,
			information.UsedMemoryPercent,
		),
	}

	fmt.Print(ui.FormatOutput(ASCIIArt, header, info))
}
