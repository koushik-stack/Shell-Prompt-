# Contributing to ShellPrompt

Thank you for your interest in contributing to ShellPrompt! This document provides guidelines and information for contributors.

## Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/shellprompt.git
   cd shellprompt
   ```

2. **Install Go** (version 1.21 or later)

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Run tests**
   ```bash
   go test ./...
   ```

## Development Workflow

### 1. Choose an Issue
- Check the [Issues](https://github.com/yourusername/shellprompt/issues) page
- Look for issues labeled `good first issue` or `help wanted`

### 2. Create a Branch
```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/issue-number
```

### 3. Make Changes
- Follow the existing code style
- Add tests for new functionality
- Update documentation as needed
- Ensure all tests pass: `go test ./...`

### 4. Commit Changes
```bash
git add .
git commit -m "feat: add new segment type"
```

Use conventional commit format:
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation
- `test:` for tests
- `refactor:` for code refactoring

### 5. Push and Create PR
```bash
git push origin your-branch-name
```
Then create a Pull Request on GitHub.

## Code Guidelines

### Go Code
- Follow standard Go formatting (`gofmt`)
- Use `go vet` to check for common issues
- Write comprehensive tests
- Use meaningful variable and function names
- Add comments for exported functions

### Testing
- Write unit tests for all new code
- Aim for good test coverage
- Use table-driven tests where appropriate
- Test edge cases and error conditions

### Documentation
- Update README.md for user-facing changes
- Add code comments for complex logic
- Update examples and configuration docs

## Adding New Segments

1. **Create segment file** in `internal/segments/`
   ```go
   package segments

   type NewSegment struct{}

   func (s *NewSegment) Render(props map[string]interface{}) (string, error) {
       // Implementation
       return "output", nil
   }
   ```

2. **Register in Registry** (`internal/segments/segment.go`)
   ```go
   Registry["newsegment"] = func() Segment { return &NewSegment{} }
   ```

3. **Add tests** in `internal/segments/segment_test.go`

4. **Update documentation** in README.md

## Project Structure

```
.
├── cmd/
│   ├── prompt/          # Main CLI application
│   └── demo/            # Web demo server
├── internal/
│   ├── config/          # Configuration loading
│   ├── segments/        # Prompt segments
│   └── shell/           # Shell-specific utilities
├── pkg/
│   └── render/          # Prompt rendering
├── themes/              # YAML theme files
└── .github/workflows/   # CI/CD pipelines
```

## Pull Request Process

1. Ensure all tests pass
2. Update documentation if needed
3. Follow conventional commit format
4. Provide a clear description of changes
5. Wait for review and address feedback

## Getting Help

- Open an [Issue](https://github.com/yourusername/shellprompt/issues) for questions
- Check existing documentation
- Review the codebase for examples

Thank you for contributing to ShellPrompt! 🎉