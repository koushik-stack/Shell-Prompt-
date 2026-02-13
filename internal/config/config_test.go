package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	// Test default config loading
	cfg, err := Load()
	if err != nil {
		t.Errorf("Load() error = %v", err)
		return
	}

	if cfg == nil {
		t.Fatal("Load() returned nil config")
	}

	if len(cfg.Segments) == 0 {
		t.Error("Expected default segments to be loaded")
	}

	// Check that we have expected segment types
	expectedTypes := map[string]bool{
		"directory": false,
		"git":       false,
		"language":  false,
		"exit_code": false,
	}

	for _, seg := range cfg.Segments {
		if _, exists := expectedTypes[seg.Type]; exists {
			expectedTypes[seg.Type] = true
		}
	}

	for segType, found := range expectedTypes {
		if !found {
			t.Errorf("Expected segment type %q not found in default config", segType)
		}
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := defaultConfig()

	if cfg == nil {
		t.Fatal("defaultConfig() returned nil")
	}

	if len(cfg.Segments) != 3 {
		t.Errorf("Expected 3 default segments, got %d", len(cfg.Segments))
	}

	expectedTypes := []string{"directory", "git", "time"}
	for i, seg := range cfg.Segments {
		if seg.Type != expectedTypes[i] {
			t.Errorf("Segment %d: expected type %q, got %q", i, expectedTypes[i], seg.Type)
		}
	}
}

func TestGetConfigPath(t *testing.T) {
	// Test with PROMPT_CONFIG env var
	originalEnv := os.Getenv("PROMPT_CONFIG")
	defer func() {
		if originalEnv == "" {
			os.Unsetenv("PROMPT_CONFIG")
		} else {
			os.Setenv("PROMPT_CONFIG", originalEnv)
		}
	}()

	testPath := "/tmp/test-config.yaml"
	os.Setenv("PROMPT_CONFIG", testPath)

	if got := getConfigPath(); got != testPath {
		t.Errorf("getConfigPath() = %q, want %q", got, testPath)
	}

	// Test default path
	os.Unsetenv("PROMPT_CONFIG")
	got := getConfigPath()
	home, _ := os.UserHomeDir()
	expected := filepath.Join(home, ".config", "shellprompt", "config.yaml")

	if got != expected {
		t.Errorf("getConfigPath() = %q, want %q", got, expected)
	}
}

func TestConfigWithFile(t *testing.T) {
	// Create a temporary config file
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")

	configContent := `
segments:
  - type: directory
    properties:
      max_depth: 5
  - type: git
    properties:
      show_status: false
`

	err := os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Temporarily change the config path
	originalEnv := os.Getenv("PROMPT_CONFIG")
	os.Setenv("PROMPT_CONFIG", configPath)
	defer func() {
		if originalEnv == "" {
			os.Unsetenv("PROMPT_CONFIG")
		} else {
			os.Setenv("PROMPT_CONFIG", originalEnv)
		}
	}()

	cfg, err := Load()
	if err != nil {
		t.Errorf("Load() error = %v", err)
		return
	}

	if len(cfg.Segments) != 2 {
		t.Errorf("Expected 2 segments, got %d", len(cfg.Segments))
	}

	if cfg.Segments[0].Type != "directory" {
		t.Errorf("First segment type = %q, want %q", cfg.Segments[0].Type, "directory")
	}

	if cfg.Segments[1].Type != "git" {
		t.Errorf("Second segment type = %q, want %q", cfg.Segments[1].Type, "git")
	}
}
