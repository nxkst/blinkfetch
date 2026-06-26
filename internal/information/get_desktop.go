package information

import (
	"fmt"
	"os"
)

func GetDesktop() (string, error) {
	desktop := os.Getenv("XDG_CURRENT_DESKTOP")
	if desktop == "" {
		return "", fmt.Errorf("couldn't fetch current desktop")
	}
	session := os.Getenv("XDG_SESSION_TYPE")
	if session == "" {
		return "", fmt.Errorf("couldn't fetch current session type")
	}

	return fmt.Sprintf("%s (%s)", desktop, session), nil
}
