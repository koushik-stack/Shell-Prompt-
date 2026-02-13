package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Segments []SegmentConfig   `yaml:"segments"`
	Colors   map[string]string `yaml:"colors,omitempty"`
	Theme    string            `yaml:"theme,omitempty"`
}

type SegmentConfig struct {
	Type  string                 `yaml:"type"`
	Props map[string]interface{} `yaml:"properties,omitempty"`
	Style SegmentStyle           `yaml:"style,omitempty"`
}

type SegmentStyle struct {
	Foreground string `yaml:"foreground,omitempty"`
	Background string `yaml:"background,omitempty"`
	Bold       bool   `yaml:"bold,omitempty"`
	Italic     bool   `yaml:"italic,omitempty"`
}

// Load reads the configuration file
func Load() (*Config, error) {
	configPath := getConfigPath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		// Return default config if file doesn't exist
		return defaultConfig(), nil
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// getConfigPath returns the path to the config file
func getConfigPath() string {
	if path := os.Getenv("PROMPT_CONFIG"); path != "" {
		return path
	}

	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "shellprompt", "config.yaml")
}

// defaultConfig returns a default configuration
func defaultConfig() *Config {
	return &Config{
		Segments: []SegmentConfig{
			{
				Type: "directory",
				Props: map[string]interface{}{
					"max_depth": 3,
					"truncate":  true,
				},
			},
			{
				Type: "git",
				Props: map[string]interface{}{
					"show_status": true,
				},
			},
			{
				Type: "time",
				Props: map[string]interface{}{
					"format": "15:04",
				},
			},
		},
	}
}
