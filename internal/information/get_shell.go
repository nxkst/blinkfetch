package information

import "os"

func GetShell() string {
	shell := os.Getenv("SHELL")

	return shell
}
