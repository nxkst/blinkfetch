package information

import "os/user"

func GetUsername() (string, error) {
	username, err := user.Current()

	if err != nil {
		return "", err
	}

	return username.Username, nil
}
