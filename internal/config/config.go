package config

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"termnia/internal/core"

	"gopkg.in/yaml.v3"
)

//go:embed config.default.yaml
var defaultConfigYAML []byte

type Config struct {
	DefaultShell core.ShellType `yaml:"defaultShell"`
}

func Load() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(home, ".termnia")
	configPath := filepath.Join(configDir, "config.yaml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return nil, err
		}
		if err := os.WriteFile(configPath, defaultConfigYAML, 0644); err != nil {
			return nil, fmt.Errorf("failed to write default config: %v", err)
		}
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}
	return &cfg, nil
}
