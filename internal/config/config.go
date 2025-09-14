package config

import (
	"os"
	"path/filepath"
)

// Config represents the application configuration
type Config struct {
	Debug   bool   `yaml:"debug" json:"debug"`
	LogFile string `yaml:"log_file" json:"log_file"`
	// Add more configuration fields as needed
}

// Default returns a default configuration
func Default() *Config {
	return &Config{
		Debug:   false,
		LogFile: "",
	}
}

// GetConfigDir returns the user's config directory for ttrk
func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "ttrk"), nil
}