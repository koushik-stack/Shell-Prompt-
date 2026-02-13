package render

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yourusername/shellprompt/internal/config"
	"github.com/yourusername/shellprompt/internal/segments"
	"github.com/yourusername/shellprompt/internal/shell"
)

// ANSI color codes
const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Italic    = "\033[3m"
	Underline = "\033[4m"

	// Foreground colors
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// Bright foreground colors
	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
)

// ColorMap maps color names to ANSI codes
var ColorMap = map[string]string{
	"black":         Black,
	"red":           Red,
	"green":         Green,
	"yellow":        Yellow,
	"blue":          Blue,
	"magenta":       Magenta,
	"cyan":          Cyan,
	"white":         White,
	"brightBlack":   BrightBlack,
	"brightRed":     BrightRed,
	"brightGreen":   BrightGreen,
	"brightYellow":  BrightYellow,
	"brightBlue":    BrightBlue,
	"brightMagenta": BrightMagenta,
	"brightCyan":    BrightCyan,
	"brightWhite":   BrightWhite,
}

// RenderPrompt renders the complete prompt for a given shell
func RenderPrompt(cfg *config.Config, shellType string) (string, error) {
	fmt.Fprintf(os.Stderr, "Rendering prompt for shell: %s\n", shellType)
	fmt.Fprintf(os.Stderr, "Config segments: %d\n", len(cfg.Segments))

	var coloredSegments []string

	for i, segCfg := range cfg.Segments {
		fmt.Fprintf(os.Stderr, "Processing segment %d: %s\n", i, segCfg.Type)
		segment := segments.New(segCfg.Type)
		if segment == nil {
			fmt.Fprintf(os.Stderr, "Segment %s not found\n", segCfg.Type)
			continue
		}

		output, err := segment.Render(segCfg.Props)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error rendering segment %s: %v\n", segCfg.Type, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "Segment %s output: '%s'\n", segCfg.Type, output)

		if output == "" {
			continue
		}

		// Apply colors if specified
		if segCfg.Style.Foreground != "" {
			output = Colorize(output, segCfg.Style.Foreground, segCfg.Style.Bold, shellType)
			fmt.Fprintf(os.Stderr, "Colored output: '%s'\n", output)
		} else if segCfg.Style.Bold {
			output = Colorize(output, "white", true, shellType)
		}

		coloredSegments = append(coloredSegments, output)
	}

	prompt := strings.Join(coloredSegments, " ")
	fmt.Fprintf(os.Stderr, "Final prompt: '%s'\n", prompt)

	// Add prompt symbol based on shell type
	symbol := shell.GetPromptSymbol(shellType)
	prompt = prompt + " " + symbol + " "
	fmt.Fprintf(os.Stderr, "Prompt with symbol: '%s'\n", prompt)

	// Apply shell-specific escaping
	finalPrompt := shell.EscapePrompt(prompt, shellType)
	fmt.Fprintf(os.Stderr, "Final escaped prompt: '%s'\n", finalPrompt)

	return finalPrompt, nil
}

// Colorize applies color and formatting to text
func Colorize(text string, color string, bold bool, shellType string) string {
	if shellType == "pwsh" || shellType == "powershell" {
		return text // PowerShell handles colors differently
	}

	colorCode := ColorMap[color]
	if colorCode == "" {
		colorCode = ColorMap["white"]
	}

	style := colorCode
	if bold {
		style = Bold + colorCode
	}

	// For bash/zsh, wrap non-printing characters
	if shellType == "bash" || shellType == "zsh" {
		return fmt.Sprintf("\\[%s\\]%s\\[%s\\]", style, text, Reset)
	}

	return fmt.Sprintf("%s%s%s", style, text, Reset)
}

// HexToANSI converts hex color to 256-color ANSI code
func HexToANSI(hexColor string) string {
	hexColor = strings.TrimPrefix(hexColor, "#")
	if len(hexColor) != 6 {
		return ""
	}

	r, _ := strconv.ParseInt(hexColor[0:2], 16, 64)
	g, _ := strconv.ParseInt(hexColor[2:4], 16, 64)
	b, _ := strconv.ParseInt(hexColor[4:6], 16, 64)

	// Convert RGB to 256-color
	if r == g && g == b {
		if r < 50 {
			return "\033[30m" // black
		}
		if r > 200 {
			return "\033[97m" // bright white
		}
	}

	// Simple 256-color approximation
	ri := int(r / 51)
	gi := int(g / 51)
	bi := int(b / 51)

	colorIndex := 16 + 36*ri + 6*gi + bi
	return fmt.Sprintf("\033[38;5;%dm", colorIndex)
}
