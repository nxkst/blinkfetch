package information

import "os/user"

func GetUsername() (string, error) {
	u, err := user.Current()

	if err != nil {
		return "", err
	}

	return u.Username, nil
}
