package main

import (
	"fmt"
	"os"

	"github.com/koushik-stack/Shell-Prompt-/internal/config"
	"github.com/koushik-stack/Shell-Prompt-/pkg/render"
)

func main() {
	fmt.Fprintf(os.Stderr, "ShellPrompt starting...\n")
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <shell>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Supported shells: bash, zsh, pwsh, fish\n")
		os.Exit(1)
	}

	shell := os.Args[1]
	fmt.Fprintf(os.Stderr, "Shell: %s\n", shell)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Config loaded successfully, segments: %d\n", len(cfg.Segments))
	fmt.Fprintf(os.Stderr, "Config loaded, segments: %d\n", len(cfg.Segments))

	// Render prompt
	prompt, err := render.RenderPrompt(cfg, shell)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering prompt: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Prompt rendered: '%s'\n", prompt)

	fmt.Print(prompt)
}
