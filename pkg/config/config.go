package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Config holds user configuration and Meta Graph API credentials.
type Config struct {
	AccessToken string `json:"access_token"`
	AccountID   string `json:"account_id"`
	Username    string `json:"username"`
	IsLoggedIn  bool   `json:"is_logged_in"`
	LastLogin   string `json:"last_login"`
}

// GetConfigFilePath returns the absolute path to ~/.config/insta-cli/config.json.
func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to determine user home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "insta-cli")
	if err := os.MkdirAll(configDir, 0705); err != nil {
		return "", fmt.Errorf("unable to create config directory: %w", err)
	}

	return filepath.Join(configDir, "config.json"), nil
}

// LoadConfig reads configuration from Environment Variables or disk config.
func LoadConfig() (Config, error) {
	var cfg Config

	// 1. Check environment variables first
	envToken := strings.TrimSpace(os.Getenv("INSTAGRAM_ACCESS_TOKEN"))
	envAccountID := strings.TrimSpace(os.Getenv("INSTAGRAM_ACCOUNT_ID"))

	if envToken != "" {
		cfg.AccessToken = envToken
		cfg.AccountID = envAccountID
		cfg.IsLoggedIn = true
		return cfg, nil
	}

	// 2. Read from config file if env variables are empty
	path, err := GetConfigFilePath()
	if err != nil {
		return cfg, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Config{IsLoggedIn: false}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	if cfg.AccessToken != "" {
		cfg.IsLoggedIn = true
	}

	return cfg, nil
}

// SaveConfig writes configuration to disk with 0600 permissions.
func SaveConfig(cfg Config) error {
	path, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	saveCfg := cfg
	if saveCfg.LastLogin == "" {
		saveCfg.LastLogin = time.Now().Format(time.RFC3339)
	}

	data, err := json.MarshalIndent(saveCfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format config JSON: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
