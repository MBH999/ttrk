package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config represents the application configuration
type Config struct {
	Debug   bool   `yaml:"debug" json:"debug"`
	LogFile string `yaml:"log_file" json:"log_file"`
	DataDir string `yaml:"data_dir" json:"data_dir"`
	// Add more configuration fields as needed
}

// Default returns a default configuration
func Default() *Config {
	return &Config{
		Debug:   false,
		LogFile: "",
		DataDir: "",
	}
}

// Load reads configuration from ~/.config/ttrk/config.ini if it exists.
func Load() (*Config, error) {
	cfg := Default()

	dir, err := GetConfigDir()
	if err != nil {
		return nil, fmt.Errorf("locate config directory: %w", err)
	}

	path := filepath.Join(dir, "config.ini")
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			resolvedDir, resErr := resolveDataDir("", dir)
			if resErr != nil {
				return nil, fmt.Errorf("resolve default data_dir: %w", resErr)
			}
			cfg.DataDir = resolvedDir
			return cfg, nil
		}
		return nil, fmt.Errorf("open config file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		if idx := strings.IndexRune(line, '='); idx != -1 {
			key := strings.TrimSpace(line[:idx])
			value := strings.TrimSpace(line[idx+1:])
			switch strings.ToLower(key) {
			case "data_dir":
				cfg.DataDir = value
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	resolvedDir, err := resolveDataDir(cfg.DataDir, dir)
	if err != nil {
		return nil, fmt.Errorf("resolve data_dir: %w", err)
	}
	cfg.DataDir = resolvedDir

	return cfg, nil
}

func resolveDataDir(value, defaultDir string) (string, error) {
	if value == "" {
		return defaultDir, nil
	}

	if strings.HasPrefix(value, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		rest := strings.TrimPrefix(value, "~")
		rest = strings.TrimPrefix(rest, "/")
		rest = strings.TrimPrefix(rest, "\\")
		if rest == "" {
			return home, nil
		}
		return filepath.Clean(filepath.Join(home, rest)), nil
	}

	if filepath.IsAbs(value) {
		return filepath.Clean(value), nil
	}

	return filepath.Clean(filepath.Join(defaultDir, value)), nil
}

// GetConfigDir returns the user's config directory for ttrk
func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "ttrk"), nil
}
