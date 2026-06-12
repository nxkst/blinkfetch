package information

import "os/user"

func GetUsername() (string, error) {
	username, err := user.Current()

	return username.Username, err
}
