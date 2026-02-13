package segments

import (
	"os/exec"
	"strings"
)

type GitSegment struct{}

func (gs *GitSegment) Render(props map[string]interface{}) (string, error) {
	// Check if we're in a git repository
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		return "", nil
	}

	showStatus := true
	if s, ok := props["show_status"].(bool); ok {
		showStatus = s
	}

	// Get current branch
	branch, err := getBranch()
	if err != nil {
		return "", nil
	}

	output := " on  " + branch

	if showStatus {
		status, err := getGitStatus()
		if err == nil && status != "" {
			output += " " + status
		}
	}

	return output, nil
}

func getBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func getGitStatus() (string, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	status := strings.TrimSpace(string(out))
	if status == "" {
		return "✓", nil
	}

	// Count changes
	lines := strings.Split(status, "\n")
	return "✗ (" + string(rune(len(lines))) + ")", nil
}
