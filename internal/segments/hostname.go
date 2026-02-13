package segments

import (
	"os"
)

type HostnameSegment struct{}

func (hs *HostnameSegment) Render(props map[string]interface{}) (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	return hostname, nil
}
