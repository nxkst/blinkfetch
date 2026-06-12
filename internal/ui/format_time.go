package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func FormatTime(timeInSeconds string) string {
	fields := strings.Fields(timeInSeconds)

	uptimeSeconds, _ := strconv.ParseFloat(fields[0], 64)
	d := time.Duration(uptimeSeconds * float64(time.Second))

	hours := int64(d.Hours())
	minutes := int64(d.Minutes()) % 60
	seconds := int64(d.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
