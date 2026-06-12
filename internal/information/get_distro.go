package information

import "github.com/Hayao0819/go-distro"

func GetDistro() string {
	distro := distro.GetDetail().FullName()

	return distro
}
