package shell

import (
	"strings"
	"testing"
)

func TestGetPromptSymbol(t *testing.T) {
	tests := []struct {
		shellType string
		expected  string
	}{
		{"bash", "❯"},
		{"zsh", "❯"},
		{"fish", "❯"},
		{"pwsh", "❯"},
		{"powershell", "❯"},
		{"unknown", "$"},
		{"BASH", "❯"},       // case insensitive
		{"Zsh", "❯"},        // case insensitive
		{"PowerShell", "❯"}, // case insensitive
		{"", "$"},           // empty string
	}

	for _, tt := range tests {
		t.Run(tt.shellType, func(t *testing.T) {
			got := GetPromptSymbol(tt.shellType)
			if got != tt.expected {
				t.Errorf("GetPromptSymbol(%q) = %q, want %q", tt.shellType, got, tt.expected)
			}
		})
	}
}

func TestEscapePrompt(t *testing.T) {
	tests := []struct {
		name      string
		prompt    string
		shellType string
		checkFunc func(string) bool
	}{
		{
			name:      "bash basic",
			prompt:    "test prompt",
			shellType: "bash",
			checkFunc: func(output string) bool {
				return output == "test prompt"
			},
		},
		{
			name:      "zsh basic",
			prompt:    "test prompt",
			shellType: "zsh",
			checkFunc: func(output string) bool {
				return output == "test prompt"
			},
		},
		{
			name:      "fish basic",
			prompt:    "test prompt",
			shellType: "fish",
			checkFunc: func(output string) bool {
				return output == "test prompt"
			},
		},
		{
			name:      "powershell basic",
			prompt:    "test prompt",
			shellType: "pwsh",
			checkFunc: func(output string) bool {
				return output == "test prompt"
			},
		},
		{
			name:      "unknown shell",
			prompt:    "test prompt",
			shellType: "unknown",
			checkFunc: func(output string) bool {
				return output == "test prompt"
			},
		},
		{
			name:      "bash with ANSI codes",
			prompt:    "\033[31mred\033[0m",
			shellType: "bash",
			checkFunc: func(output string) bool {
				return strings.Contains(output, "\\[") && strings.Contains(output, "\\]")
			},
		},
		{
			name:      "zsh with ANSI codes",
			prompt:    "\033[32mgreen\033[0m",
			shellType: "zsh",
			checkFunc: func(output string) bool {
				return strings.Contains(output, "\\[") && strings.Contains(output, "\\]")
			},
		},
		{
			name:      "fish with ANSI codes",
			prompt:    "\033[34mblue\033[0m",
			shellType: "fish",
			checkFunc: func(output string) bool {
				return strings.Contains(output, "\033[34m") && strings.Contains(output, "\033[0m")
			},
		},
		{
			name:      "powershell with ANSI codes",
			prompt:    "\033[35mmagenta\033[0m",
			shellType: "pwsh",
			checkFunc: func(output string) bool {
				return output == "\033[35mmagenta\033[0m"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EscapePrompt(tt.prompt, tt.shellType)
			if !tt.checkFunc(got) {
				t.Errorf("EscapePrompt(%q, %q) = %q, failed check", tt.prompt, tt.shellType, got)
			}
		})
	}
}
