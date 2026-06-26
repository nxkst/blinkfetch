package information

import (
	"fmt"

	"github.com/Hayao0819/go-distro"
)

func GetDistro() (string, error) {
	distro := distro.GetDetail().FullName()
	if distro == "" {
		return "", fmt.Errorf("couldn't fetch distro name")
	}

	return distro, nil
}
