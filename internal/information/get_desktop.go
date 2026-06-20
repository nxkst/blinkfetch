package information

import (
	"fmt"
	"os"
)

func GetDesktop() string {
	desktop := os.Getenv("XDG_CURRENT_DESKTOP")
	session := os.Getenv("XDG_SESSION_TYPE")

	return fmt.Sprintf("%s (%s)", desktop, session)
}
