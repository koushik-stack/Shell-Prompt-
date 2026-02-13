.PHONY: all build test clean fmt vet lint demo run-demo help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOVET=$(GOCMD) vet

# Main package
MAIN_PATH=./cmd/prompt
DEMO_PATH=./cmd/demo
BINARY_NAME=prompt
DEMO_BINARY=demo

# Build the project
all: clean fmt vet test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)

build-demo:
	$(GOBUILD) -o $(DEMO_BINARY) -v $(DEMO_PATH)

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -race -coverprofile=coverage.out -covermode=atomic ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(DEMO_BINARY)
	rm -f coverage.out coverage.html

# Format code
fmt:
	$(GOFMT) -s -w .

# Run go vet
vet:
	$(GOVET) ./...

# Run golangci-lint if installed
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Run security scan
security:
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "gosec not installed. Install with: go install github.com/securecodewarrior/github-action-gosec@latest"; \
	fi

# Download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Build and run demo
demo: build-demo
	./$(DEMO_BINARY)

run-demo: demo

# Run the prompt for testing
run-bash: build
	./$(BINARY_NAME) bash

run-zsh: build
	./$(BINARY_NAME) zsh

run-pwsh: build
	./$(BINARY_NAME) pwsh

run-fish: build
	./$(BINARY_NAME) fish

# Development setup
dev-setup: deps
	@echo "Development environment set up successfully!"

# CI simulation
ci: fmt vet lint test build build-demo

# Help
help:
	@echo "Available targets:"
	@echo "  all          - Run clean, fmt, vet, test, and build"
	@echo "  build        - Build the main binary"
	@echo "  build-demo   - Build the demo binary"
	@echo "  test         - Run all tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  clean        - Clean build files"
	@echo "  fmt          - Format code"
	@echo "  vet          - Run go vet"
	@echo "  lint         - Run golangci-lint"
	@echo "  security     - Run security scan"
	@echo "  deps         - Download and tidy dependencies"
	@echo "  demo         - Build and run demo"
	@echo "  run-bash     - Build and test bash prompt"
	@echo "  run-zsh      - Build and test zsh prompt"
	@echo "  run-pwsh     - Build and test PowerShell prompt"
	@echo "  run-fish     - Build and test fish prompt"
	@echo "  dev-setup    - Set up development environment"
	@echo "  ci           - Simulate CI pipeline"
	@echo "  help         - Show this help"