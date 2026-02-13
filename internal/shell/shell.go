package shell

import (
	"regexp"
	"strings"
)

// GetPromptSymbol returns the appropriate prompt symbol for a shell
func GetPromptSymbol(shellType string) string {
	shellType = strings.ToLower(shellType)

	switch shellType {
	case "pwsh", "powershell":
		return "❯"
	case "bash", "zsh", "fish":
		return "❯"
	default:
		return "$"
	}
}

// EscapePrompt escapes the prompt for shell-specific requirements
func EscapePrompt(prompt, shellType string) string {
	shellType = strings.ToLower(shellType)

	switch shellType {
	case "bash", "zsh":
		// Escape special characters for bash/zsh
		return escapeBashZsh(prompt)
	case "fish":
		return escapeFish(prompt)
	case "pwsh", "powershell":
		return escapePowerShell(prompt)
	default:
		return prompt
	}
}

func escapeBashZsh(prompt string) string {
	// Wrap ANSI escape sequences in \[ and \] to tell bash/zsh they are non-printing
	// ANSI escape sequences match the pattern \033[...m
	re := regexp.MustCompile(`\033\[[0-9;]*m`)
	return re.ReplaceAllStringFunc(prompt, func(match string) string {
		return "\\[" + match + "\\]"
	})
}

func escapeFish(prompt string) string {
	// Fish doesn't require escaping of ANSI sequences in the same way
	return prompt
}

func escapePowerShell(prompt string) string {
	// PowerShell doesn't require escaping of ANSI sequences
	return prompt
}
