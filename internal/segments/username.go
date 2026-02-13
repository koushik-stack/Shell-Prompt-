package segments

import (
	"os"
	"os/user"
)

type UsernameSegment struct{}

func (us *UsernameSegment) Render(props map[string]interface{}) (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	showHost := false
	if s, ok := props["show_host"].(bool); ok {
		showHost = s
	}

	output := currentUser.Username

	if showHost {
		hostname, err := os.Hostname()
		if err == nil {
			output += "@" + hostname
		}
	}

	return output, nil
}
