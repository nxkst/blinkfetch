package information

import (
	"log"

	"github.com/nxkst/blinkfetch/internal/ui"
)

func must[T any](val T, err error) T {
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func mustMemory() (used, total, usedPercent string) {
	used, total, usedPercent, err := GetMemory()

	if err != nil {
		log.Fatal(err)
	}

	return used, total, usedPercent
}

func InitModules() {
	Username = must(GetUsername())
	Hostname = must(GetHostname())
	Distro = must(GetDistro())
	Kernel = must(GetKernel())
	Uptime = must(GetUptime())
	Desktop = must(GetDesktop())
	Shell = must(GetShell())

	UsedMemory, TotalMemory, UsedMemoryPercent = mustMemory()

	FormattedUptime = ui.FormatTime(Uptime)
}
