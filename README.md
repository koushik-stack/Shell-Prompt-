# ShellPrompt

A cross-platform shell prompt customization tool written in Go, inspired by [oh-my-posh](https://github.com/JanDeDobbeleer/oh-my-posh).

## Features

- **Cross-platform support**: bash, zsh, PowerShell, and fish shells
- **Customizable segments**: Directory, Git status, programming language detection, time, username, and more
- **YAML/JSON configuration**: Easy-to-read configuration files
- **Fast rendering**: Written in Go for performance
- **Theme support**: Pre-built and custom themes

## Segments

- `directory` - Current working directory with truncation support
- `git` - Git branch and status indicators
- `language` - Programming language detection
- `time` - Current time display
- `exit_code` - Last command exit code
- `username` - Current user name
- `hostname` - Machine hostname

## Installation

### Build from Source

```bash
go build -o prompt ./cmd/prompt
```

### Add to Your Shell

#### Bash/Zsh
Add to your `.bashrc` or `.zshrc`:

```bash
eval "$(~/path/to/prompt bash)"
```

#### Fish
Add to your `config.fish`:

```fish
~/path/to/prompt fish | source
```

#### PowerShell
Add to your `$PROFILE`:

```powershell
~/path/to/prompt pwsh | Out-String | Invoke-Expression
```

## Configuration

Configuration file should be placed at:
- `~/.config/shellprompt/config.yaml` (Unix-like systems)
- `%APPDATA%\shellprompt\config.yaml` (Windows)

You can also set the `PROMPT_CONFIG` environment variable to use a custom location.

### Example Configuration

```yaml
segments:
  - type: directory
    properties:
      max_depth: 3
      truncate: true
    style:
      foreground: "blue"
      bold: true

  - type: git
    properties:
      show_status: true
    style:
      foreground: "green"

  - type: time
    properties:
      format: "15:04"
    style:
      foreground: "yellow"

colors:
  primary: "#00FF00"
  secondary: "#0000FF"
```

## Themes

ShellPrompt comes with several pre-built themes:

### Available Themes

- **`default`** - Balanced theme with blue directories, green git status, and yellow language indicators
- **`minimal`** - Clean theme with essential segments only
- **`dark`** - Dark theme with cyan, green, yellow, magenta, and red colors
- **`light`** - Light theme with blue, green, purple, cyan, and red colors
- **`colorful`** - Bright theme using bright colors for all segments
- **`monochrome`** - Minimalist theme using white and gray tones
- **`powerline`** - Powerline-style theme with background colors and more segments
- **`simple`** - Very basic theme with just directory and git status
- **`developer`** - Developer-focused theme with all segments and bright colors
- **`futuristic`** - Sci-fi theme with cyan, green, magenta, yellow, and blue
- **`professional`** - Business-appropriate theme with conservative colors
- **`hacker`** - Matrix-style theme with green-on-black aesthetic
- **`warm`** - Cozy theme with orange, brown, and warm colors
- **`rainbow`** - Colorful theme using all rainbow colors
- **`ocean`** - Aquatic theme with blue and aqua tones

### Using Themes

To use a theme, copy it to your config location:

```bash
# Copy a theme to your config directory
cp themes/default.yaml ~/.config/shellprompt/config.yaml
```

Or set the `PROMPT_CONFIG` environment variable:

```bash
export PROMPT_CONFIG=~/themes/colorful.yaml
```

## Development

### Project Structure

```
.
├── cmd/prompt/          # Main application
├── internal/
│   ├── config/         # Configuration loading
│   ├── segments/       # Prompt segments
│   └── shell/          # Shell-specific helpers
├── pkg/render/         # Rendering logic
└── themes/             # Theme files
```

### Building

```bash
go build -o prompt ./cmd/prompt
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific package tests
go test ./internal/config
```

### Code Quality

```bash
# Format code
gofmt -s -w .

# Run static analysis
go vet ./...

# Run linter (if installed)
golangci-lint run
```

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
