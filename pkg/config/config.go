package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Config holds user configuration and active session details.
type Config struct {
	Username     string `json:"username"`
	SessionToken string `json:"session_token"`
	IsLoggedIn   bool   `json:"is_logged_in"`
	LastLogin    string `json:"last_login"`
}

// GetConfigFilePath returns the absolute path to ~/.config/insta-cli/config.json.
// It also ensures that the target directory exists.
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

// encodeToken obfuscates the session token using Base64 encoding.
func encodeToken(token string) string {
	if token == "" {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(token))
}

// decodeToken decodes the Base64 obfuscated session token.
func decodeToken(token string) string {
	if token == "" {
		return ""
	}
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return token
	}
	return string(data)
}

// LoadConfig reads and decodes the configuration file from disk.
func LoadConfig() (Config, error) {
	path, err := GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Config{IsLoggedIn: false}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	cfg.SessionToken = decodeToken(cfg.SessionToken)
	return cfg, nil
}

// SaveConfig obfuscates credentials and writes config to disk with 0600 permissions.
func SaveConfig(cfg Config) error {
	path, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	saveCfg := cfg
	saveCfg.SessionToken = encodeToken(cfg.SessionToken)
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
