package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type LuminaConfig struct {
	DefaultModel string `yaml:"default_model"`
}

func ConfigPath() string {
	usr, _ := user.Current()
	return filepath.Join(usr.HomeDir, ".lumina.yaml")
}

func LoadConfig() (*LuminaConfig, error) {
	path := ConfigPath()
	data, err := os.ReadFile(path)
	if err != nil {
		return &LuminaConfig{}, err
	}

	var cfg LuminaConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return &LuminaConfig{}, err
	}

	return &cfg, nil
}

func SaveConfig(cfg *LuminaConfig) error {
	path := ConfigPath()
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	fmt.Printf("✅ Default model set to '%s'\n", cfg.DefaultModel)

	return nil
}
