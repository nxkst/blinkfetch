package main

import "github.com/nxkst/blinkfetch/information"

var (
	Username, _     = information.GetUsername()
	Hostname, _     = information.GetHostname()
	Distro          = information.GetDistro()
	Kernel, _       = information.GetKernel()
	Uptime, _       = information.GetUptime()
	FormattedUptime = FormatTime(Uptime)
)
