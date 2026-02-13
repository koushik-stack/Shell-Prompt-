package segments

import (
	"os"
	"path/filepath"
)

type LanguageSegment struct{}

func (ls *LanguageSegment) Render(props map[string]interface{}) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Check for language indicators
	langIndicators := map[string][]string{
		" go":   {"go.mod", "go.sum"},
		" py":   {"main.py", "requirements.txt", "setup.py"},
		" node": {"package.json", "node_modules"},
		" rs":   {"Cargo.toml", "Cargo.lock"},
		" ts":   {"tsconfig.json"},
		" java": {"pom.xml", "build.gradle"},
		" rb":   {"Gemfile", "Rakefile"},
	}

	for lang, files := range langIndicators {
		for _, file := range files {
			if _, err := os.Stat(filepath.Join(cwd, file)); err == nil {
				return lang, nil
			}
		}
	}

	return "", nil
}
