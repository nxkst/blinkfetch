package information

import (
	"log"

	"github.com/nxkst/blinkfetch/internal/ui"
)

func InitModules() {
	var err error

	Username, err = GetUsername()
	if err != nil {
		log.Fatal(err)
	}
	Hostname, err = GetHostname()
	if err != nil {
		log.Fatal(err)
	}
	Distro = GetDistro()
	Kernel, err = GetKernel()
	if err != nil {
		log.Fatal(err)
	}
	Uptime, err = GetUptime()
	if err != nil {
		log.Fatal(err)
	}
	FormattedUptime = ui.FormatTime(Uptime)
	Desktop = GetDesktop()
	Shell = GetShell()
	UsedMemory, TotalMemory, UsedMemoryPercent, err = GetMemory()
	if err != nil {
		log.Fatal(err)
	}
}
