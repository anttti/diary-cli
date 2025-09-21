package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds the application configuration
type Config struct {
	Dir string `json:"dir"`
}

// LoadConfig loads configuration from ~/.config/diary/diary.json if it exists
// Returns a Config with the loaded settings, or an empty Config if file doesn't exist
func LoadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return &Config{}, nil // Return empty config on error
	}

	configPath := filepath.Join(homeDir, ".config", "diary", "diary.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil // Config file doesn't exist, return empty config
		}
		return nil, err // Other errors
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// GetDir returns the directory to use based on priority:
// 1. CLI argument (if not empty)
// 2. Config file setting (if not empty)
// 3. Default value
func GetDir(cliDir string, configDir string, defaultDir string) string {
	if cliDir != "" && cliDir != defaultDir {
		return cliDir
	}
	if configDir != "" {
		return configDir
	}
	return defaultDir
}
