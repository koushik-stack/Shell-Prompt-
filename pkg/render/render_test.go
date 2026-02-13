package render

import (
	"strings"
	"testing"

	"github.com/koushik-stack/Shell-Prompt-/internal/config"
)

func TestRenderPrompt(t *testing.T) {
	tests := []struct {
		name      string
		cfg       *config.Config
		shellType string
		checkFunc func(string) bool
	}{
		{
			name: "basic bash prompt",
			cfg: &config.Config{
				Segments: []config.SegmentConfig{
					{Type: "directory"},
				},
			},
			shellType: "bash",
			checkFunc: func(output string) bool {
				return strings.Contains(output, "❯") && output != ""
			},
		},
		{
			name: "powershell prompt",
			cfg: &config.Config{
				Segments: []config.SegmentConfig{
					{Type: "directory"},
				},
			},
			shellType: "pwsh",
			checkFunc: func(output string) bool {
				return strings.Contains(output, "❯") && output != ""
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := RenderPrompt(tt.cfg, tt.shellType)
			if err != nil {
				t.Errorf("RenderPrompt() error = %v", err)
				return
			}
			if !tt.checkFunc(output) {
				t.Errorf("RenderPrompt() = %q, failed check", output)
			}
		})
	}
}

func TestColorize(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		color     string
		bold      bool
		shellType string
		checkFunc func(string) bool
	}{
		{
			name:      "bash blue color",
			text:      "test",
			color:     "blue",
			bold:      false,
			shellType: "bash",
			checkFunc: func(output string) bool {
				return strings.Contains(output, "\033[34m") && strings.Contains(output, "\033[0m")
			},
		},
		{
			name:      "powershell no color",
			text:      "test",
			color:     "blue",
			bold:      false,
			shellType: "pwsh",
			checkFunc: func(output string) bool {
				return output == "test"
			},
		},
		{
			name:      "bold text",
			text:      "test",
			color:     "green",
			bold:      true,
			shellType: "bash",
			checkFunc: func(output string) bool {
				return strings.Contains(output, "\033[1m") && strings.Contains(output, "\033[32m")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := Colorize(tt.text, tt.color, tt.bold, tt.shellType)
			if !tt.checkFunc(output) {
				t.Errorf("Colorize(%q, %q, %v, %q) = %q", tt.text, tt.color, tt.bold, tt.shellType, output)
			}
		})
	}
}

func TestHexToANSI(t *testing.T) {
	tests := []struct {
		name     string
		hexColor string
		expected string
	}{
		{
			name:     "valid hex",
			hexColor: "#FF0000",
			expected: "\033[38;5;196m", // Red
		},
		{
			name:     "invalid hex",
			hexColor: "invalid",
			expected: "",
		},
		{
			name:     "short hex",
			hexColor: "#FFF",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HexToANSI(tt.hexColor)
			if got != tt.expected {
				t.Errorf("HexToANSI(%q) = %q, want %q", tt.hexColor, got, tt.expected)
			}
		})
	}
}
