package main

import (
	"fmt"
	"github.com/nxkst/blinkfetch/internal/information"
	"github.com/nxkst/blinkfetch/internal/ui"
)

func main() {
	ASCIIArt := ui.FormatASCII()

	header := fmt.Sprintf(
		"%s@%s",
		information.Username, information.Hostname,
	)

	info := []string{
		"──────────────",
		fmt.Sprintf("\033[36mos:\033[0m %s",
			information.Distro,
		),
		fmt.Sprintf("\033[36mkernel:\033[0m %s",
			information.Kernel,
		),
		fmt.Sprintf("\033[36muptime:\033[0m %s",
			information.FormattedUptime,
		),
		fmt.Sprintf("\033[36mde:\033[0m %s",
			information.Desktop,
		),
		fmt.Sprintf("\033[36mshell:\033[0m %s",
			information.Shell,
		),
	}

	fmt.Print(ui.FormatOutput(ASCIIArt, header, info))
}
