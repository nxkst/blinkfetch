package information

import (
	"os"
	"strings"
)

func GetKernel() (string, error) {
	kernel, err := os.ReadFile("/proc/sys/kernel/osrelease")

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(kernel)), nil
}
