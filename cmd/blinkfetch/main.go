package main

import (
	"fmt"

	"github.com/nxkst/blinkfetch/internal/information"
	"github.com/nxkst/blinkfetch/internal/ui"
)

func main() {
	yellow := "\x1b[33m"
	reset := "\033[0m"

	ASCIIArt := ui.FormatASCII()

	information.InitModules()

	header := fmt.Sprintf(
		reset+"%s@%s",
		information.Username, information.Hostname,
	)

	info := []string{
		"──────────────",
		fmt.Sprintf(yellow+"os:"+reset+" %s",
			information.Distro,
		),
		fmt.Sprintf(yellow+"kernel:"+reset+" %s",
			information.Kernel,
		),
		fmt.Sprintf(yellow+"uptime:"+reset+" %s",
			information.FormattedUptime,
		),
		fmt.Sprintf(yellow+"de:"+reset+" %s",
			information.Desktop,
		),
		fmt.Sprintf(yellow+"shell:"+reset+" %s",
			information.Shell,
		),
		fmt.Sprintf(yellow+"memory:"+reset+" %s MiB / %s MiB (%s%%)",
			information.UsedMemory,
			information.TotalMemory,
			information.UsedMemoryPercent,
		),
	}

	fmt.Print(ui.FormatOutput(ASCIIArt, header, info))
}
