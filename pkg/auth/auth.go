package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"go-insta-cli/pkg/cli"
	"go-insta-cli/pkg/config"
)

// RequireAuth checks if a valid Meta Graph API Access Token is present.
func RequireAuth(cfg config.Config) bool {
	if !cfg.IsLoggedIn || strings.TrimSpace(cfg.AccessToken) == "" {
		fmt.Printf("%sError: Meta Graph API Access Token is missing.%s\n", cli.ColorRed, cli.ColorReset)
		fmt.Printf("Please set INSTAGRAM_ACCESS_TOKEN environment variable or run 'insta login'.\n")
		return false
	}
	return true
}

// PerformLogin prompts for Meta Graph Access Token and Account ID, saving config.
func PerformLogin(accessToken, accountID string) (config.Config, error) {
	reader := bufio.NewReader(os.Stdin)

	if accessToken == "" {
		fmt.Printf("%sEnter Meta Graph API Access Token (EAA... or IG...): %s", cli.ColorCyan, cli.ColorReset)
		input, err := reader.ReadString('\n')
		if err != nil {
			return config.Config{}, fmt.Errorf("failed to read access token: %w", err)
		}
		accessToken = strings.TrimSpace(input)
	}

	if strings.Contains(accessToken, "sessionid=") || strings.Contains(accessToken, "csrftoken=") {
		return config.Config{}, fmt.Errorf("browser cookies are NOT supported. Please provide an official Meta Graph API Access Token")
	}

	if accountID == "" {
		fmt.Printf("%sEnter Instagram Account ID (optional, press Enter to skip): %s", cli.ColorCyan, cli.ColorReset)
		input, _ := reader.ReadString('\n')
		accountID = strings.TrimSpace(input)
	}

	cfg := config.Config{
		AccessToken: accessToken,
		AccountID:   accountID,
		IsLoggedIn:  true,
		LastLogin:   time.Now().Format(time.RFC3339),
	}

	if err := config.SaveConfig(cfg); err != nil {
		return config.Config{}, fmt.Errorf("failed to save config: %w", err)
	}

	return cfg, nil
}

// PerformLogout clears active Meta credentials.
func PerformLogout() error {
	cfg := config.Config{
		IsLoggedIn: false,
	}
	return config.SaveConfig(cfg)
}
