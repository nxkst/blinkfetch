package information

import (
	"fmt"
	"os"
)

func GetShell() (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "", fmt.Errorf("couldn't fetch shell")
	}

	return shell, nil
}
