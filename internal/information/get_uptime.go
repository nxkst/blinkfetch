package information

import "os"

func GetUptime() (string, error) {
	uptime, err := os.ReadFile("/proc/uptime")

	return string(uptime), err
}
