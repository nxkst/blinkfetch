package information

import "os"

func GetUptime() (string, error) {
	uptime, err := os.ReadFile("/proc/uptime")

	if err != nil {
		return "", err
	}

	return string(uptime), nil
}
