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
	Distro, err = GetDistro()
	if err != nil {
		log.Fatal(err)
	}
	Kernel, err = GetKernel()
	if err != nil {
		log.Fatal(err)
	}
	Uptime, err = GetUptime()
	if err != nil {
		log.Fatal(err)
	}
	FormattedUptime = ui.FormatTime(Uptime)
	Desktop, err = GetDesktop()
	if err != nil {
		log.Fatal(err)
	}
	Shell, err = GetShell()
	if err != nil {
		log.Fatal(err)
	}
	UsedMemory, TotalMemory, UsedMemoryPercent, err = GetMemory()
	if err != nil {
		log.Fatal(err)
	}
}
