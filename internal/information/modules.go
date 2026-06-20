package information

import "github.com/nxkst/blinkfetch/internal/ui"

var (
	Username, _                                   = GetUsername()
	Hostname, _                                   = GetHostname()
	Distro                                        = GetDistro()
	Kernel, _                                     = GetKernel()
	Uptime, _                                     = GetUptime()
	FormattedUptime                               = ui.FormatTime(Uptime)
	Desktop                                       = GetDesktop()
	Shell                                         = GetShell()
	UsedMemory, TotalMemory, UsedMemoryPercent, _ = GetMemory()
)
