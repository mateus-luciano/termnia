package config

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mateus-luciano/termnia/internal/types"
	"gopkg.in/yaml.v3"
)

//go:embed config.default.yaml
var defaultConfigYAML []byte

type Config struct {
	DefaultShell types.ShellType `yaml:"defaultShell"`
	Theme        string          `yaml:"theme"`
}

var cfg Config

func configDir() string {
	home, _ := os.UserHomeDir()

	return filepath.Join(home, ".termnia")
}

func configPath() string {
	return filepath.Join(configDir(), "config.yaml")
}

func Load() (*Config, error) {
	configDir := configDir()
	configPath := configPath()

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

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &cfg, nil
}

func Get() Config {
	return cfg
}

func Save(c Config) error {
	cfg = c

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to serialize config: %v", err)
	}

	return os.WriteFile(configPath(), data, 0644)
}
