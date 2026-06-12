package information

import (
	"os"
	"strings"
)

func GetKernel() (string, error) {
	kernel, err := os.ReadFile("/proc/sys/kernel/osrelease")

	return strings.TrimSpace(string(kernel)), err
}
