package information

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/mem"
	"math"
)

const MiB = 1024 * 1024

func GetMemory() (string, string, string, error) {
	vm, err := mem.VirtualMemory()

	if err != nil {
		return "", "", "", err
	}

	used := fmt.Sprintf("%.0f", math.Round(float64(vm.Used)/MiB))
	total := fmt.Sprintf("%.0f", math.Round(float64(vm.Total)/MiB))
	usedPercent := fmt.Sprintf("%.0f", vm.UsedPercent)

	return used, total, usedPercent, nil
}
