package segments

import (
	"os"
	"strconv"
)

type ExitCodeSegment struct{}

func (ecs *ExitCodeSegment) Render(props map[string]interface{}) (string, error) {
	exitCodeStr := os.Getenv("EXIT_CODE")
	if exitCodeStr == "" {
		return "", nil
	}

	exitCode, err := strconv.Atoi(exitCodeStr)
	if err != nil {
		return "", nil
	}

	if exitCode != 0 {
		return " ✗ " + exitCodeStr, nil
	}

	return "", nil
}
