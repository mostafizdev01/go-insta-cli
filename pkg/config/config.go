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

// GetConfigFilePath returns the absolute path to F:\Riseup-asia\config.json.
func GetConfigFilePath() (string, error) {
	configDir := `F:\Riseup-asia`
	if err := os.MkdirAll(configDir, 0705); err != nil {
		return "", fmt.Errorf("unable to create config directory: %w", err)
	}

	return filepath.Join(configDir, "config.json"), nil
}

// loadDotEnv parses .env file key-value pairs if present.
func loadDotEnv() map[string]string {
	res := make(map[string]string)
	paths := []string{
		".env",
		filepath.Join("F:", "Riseup-asia", ".env"),
		filepath.Join("F:", "Riseup-asia", "go-insta-cli", ".env"),
	}
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err == nil {
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
					continue
				}
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					k := strings.TrimSpace(parts[0])
					v := strings.TrimSpace(parts[1])
					v = strings.Trim(v, `"'`)
					if k != "" && v != "" {
						res[k] = v
					}
				}
			}
			if len(res) > 0 {
				break
			}
		}
	}
	return res
}

// LoadConfig reads configuration from Environment Variables, .env file, or disk config.
func LoadConfig() (Config, error) {
	var cfg Config

	// 1. Check environment variables & .env file
	dotEnv := loadDotEnv()
	envToken := strings.TrimSpace(os.Getenv("INSTAGRAM_ACCESS_TOKEN"))
	if envToken == "" {
		envToken = dotEnv["INSTAGRAM_ACCESS_TOKEN"]
	}
	envAccountID := strings.TrimSpace(os.Getenv("INSTAGRAM_ACCOUNT_ID"))
	if envAccountID == "" {
		envAccountID = dotEnv["INSTAGRAM_ACCOUNT_ID"]
	}

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

// SaveConfig writes configuration to disk with 0600 permissions and updates .env if present.
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

	// Also sync to .env file if it exists
	dotEnvPaths := []string{
		".env",
		filepath.Join("F:", "Riseup-asia", ".env"),
		filepath.Join("F:", "Riseup-asia", "go-insta-cli", ".env"),
	}
	envContent := fmt.Sprintf("# Meta Instagram Official Graph API Configuration\nINSTAGRAM_ACCESS_TOKEN=%s\nINSTAGRAM_ACCOUNT_ID=%s\n", cfg.AccessToken, cfg.AccountID)
	for _, envPath := range dotEnvPaths {
		if _, err := os.Stat(envPath); err == nil {
			_ = os.WriteFile(envPath, []byte(envContent), 0600)
		}
	}

	return nil
}
