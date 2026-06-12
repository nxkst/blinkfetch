package main

import (
	"fmt"
	"github.com/nxkst/blinkfetch/ui"
)

func main() {
	ASCIIArt := ui.FormatASCII()

	header := fmt.Sprintf("%s@%s", Username, Hostname)

	info := []string{
		"──────────────",
		fmt.Sprintf("\033[36mos:\033[0m %s", Distro),
		fmt.Sprintf("\033[36mkernel:\033[0m %s", Kernel),
		fmt.Sprintf("\033[36muptime:\033[0m %s", FormattedUptime),
	}

	fmt.Print(FormatOutput(ASCIIArt, header, info))
}
