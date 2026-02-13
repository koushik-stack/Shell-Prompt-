package segments

import (
	"os"
	"strings"
)

type DirectorySegment struct{}

func (ds *DirectorySegment) Render(props map[string]interface{}) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	home, err := os.UserHomeDir()
	if err == nil && strings.HasPrefix(cwd, home) {
		cwd = strings.Replace(cwd, home, "~", 1)
	}

	maxDepth := 3
	if d, ok := props["max_depth"].(float64); ok {
		maxDepth = int(d)
	}

	truncate := true
	if t, ok := props["truncate"].(bool); ok {
		truncate = t
	}

	if truncate && strings.Count(cwd, string(os.PathSeparator)) > maxDepth {
		parts := strings.Split(cwd, string(os.PathSeparator))
		if len(parts) > maxDepth {
			cwd = "..." + string(os.PathSeparator) + strings.Join(parts[len(parts)-maxDepth:], string(os.PathSeparator))
		}
	}

	return cwd, nil
}
