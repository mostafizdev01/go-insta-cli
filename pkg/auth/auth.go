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

// RequireAuth checks if the user is logged in and has a valid session token.
// If not authenticated, it prints a red error message and returns false.
func RequireAuth(cfg config.Config) bool {
	if !cfg.IsLoggedIn || strings.TrimSpace(cfg.SessionToken) == "" {
		fmt.Printf("%sError: You must be logged in. Run 'insta login' first.%s\n", cli.ColorRed, cli.ColorReset)
		return false
	}
	return true
}

// PerformLogin prompts for credentials if necessary, validates them, and saves session.
func PerformLogin(username, sessionToken string) (config.Config, error) {
	reader := bufio.NewReader(os.Stdin)

	if username == "" {
		fmt.Printf("%sEnter Instagram Username or App Name: %s", cli.ColorCyan, cli.ColorReset)
		input, err := reader.ReadString('\n')
		if err != nil {
			return config.Config{}, fmt.Errorf("failed to read username: %w", err)
		}
		username = strings.TrimSpace(input)
	}

	if sessionToken == "" {
		fmt.Printf("%sEnter Meta Access Token / Instagram Session Cookie: %s", cli.ColorCyan, cli.ColorReset)
		input, err := reader.ReadString('\n')
		if err != nil {
			return config.Config{}, fmt.Errorf("failed to read session token: %w", err)
		}
		sessionToken = strings.TrimSpace(input)
	}

	if username == "" || sessionToken == "" {
		return config.Config{}, fmt.Errorf("username and session token cannot be empty")
	}

	cfg := config.Config{
		Username:     username,
		SessionToken: sessionToken,
		IsLoggedIn:   true,
		LastLogin:    time.Now().Format(time.RFC3339),
	}

	if err := config.SaveConfig(cfg); err != nil {
		return config.Config{}, fmt.Errorf("failed to persist session: %w", err)
	}

	return cfg, nil
}

// PerformLogout clears active session data from config.json.
func PerformLogout() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load session: %w", err)
	}

	cfg.IsLoggedIn = false
	cfg.SessionToken = ""

	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("failed to clear session: %w", err)
	}

	return nil
}
