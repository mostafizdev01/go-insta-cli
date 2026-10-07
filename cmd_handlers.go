package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-insta-cli/pkg/auth"
	"go-insta-cli/pkg/cli"
	"go-insta-cli/pkg/config"
	"go-insta-cli/pkg/instagram"
	"go-insta-cli/pkg/posts"
)

func handleLogin(args []string) {
	token := ""
	accountID := ""

	if len(args) > 1 {
		token = args[1]
	}
	if len(args) > 2 {
		accountID = args[2]
	}

	fmt.Printf("%sInitiating Meta Graph API authentication...%s\n", cli.ColorYellow, cli.ColorReset)
	_, err := auth.PerformLogin(token, accountID)
	if err != nil {
		fmt.Printf("%sError during login: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	fmt.Printf("%sSuccessfully authenticated and saved Meta Graph credentials.%s\n", cli.ColorGreen, cli.ColorReset)
}

func handleLogout() {
	if err := auth.PerformLogout(); err != nil {
		fmt.Printf("%sError during logout: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}
	fmt.Printf("%sLogged out successfully. Credentials cleared.%s\n", cli.ColorYellow, cli.ColorReset)
}

func handleStatus(cfg config.Config) {
	fmt.Printf("%sMeta Instagram Graph API Integration Status:%s\n", cli.ColorCyan, cli.ColorReset)

	if !cfg.IsLoggedIn || strings.TrimSpace(cfg.AccessToken) == "" {
		fmt.Printf("  Status:     %sNot Configured%s\n", cli.ColorYellow, cli.ColorReset)
		fmt.Printf("  Message:    Set INSTAGRAM_ACCESS_TOKEN environment variable or run 'insta login'.\n")
		return
	}

	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	user, err := client.FetchUserProfile()
	if err != nil {
		fmt.Printf("  Status:     %sInvalid Credentials / Expired Token%s\n", cli.ColorRed, cli.ColorReset)
		fmt.Printf("  Error:      %v\n", err)
		return
	}

	limiterSummary := instagram.GlobalLimiter.GetStatusSummary()
	fmt.Printf("  Status:       %sConnected (Verified via Meta Graph API)%s\n", cli.ColorGreen, cli.ColorReset)
	fmt.Printf("  Username:     %s%s%s\n", cli.ColorGreen, user.Username, cli.ColorReset)
	fmt.Printf("  Account ID:   %s%s%s\n", cli.ColorGreen, user.ID, cli.ColorReset)
	fmt.Printf("  Rate Limiter: %s%v%s (Safety Pause: %vms)\n", cli.ColorYellow, limiterSummary["status"], cli.ColorReset, limiterSummary["safety_interval_ms"])
}

func handlePosts(args []string, cfg config.Config) {
	limit := 10
	isJSON := false

	for i := 1; i < len(args); i++ {
		arg := args[i]
		if arg == "--json" {
			isJSON = true
		} else if strings.HasPrefix(arg, "--limit=") {
			parts := strings.Split(arg, "=")
			if len(parts) == 2 {
				if l, err := strconv.Atoi(parts[1]); err == nil && l > 0 {
					limit = l
				}
			}
		} else if arg == "--limit" && i+1 < len(args) {
			if l, err := strconv.Atoi(args[i+1]); err == nil && l > 0 {
				limit = l
				i++
			}
		}
	}

	fmt.Printf("%sFetching Instagram posts via Meta Graph API (limit: %d)...%s\n", cli.ColorCyan, limit, cli.ColorReset)
	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	postsList, err := client.FetchUserPosts(limit)
	if err != nil {
		fmt.Printf("%sError fetching posts: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	if isJSON {
		jsonOutput, err := posts.RenderPostJSON(postsList)
		if err != nil {
			fmt.Printf("%sError rendering JSON: %v%s\n", cli.ColorRed, err, cli.ColorReset)
			os.Exit(1)
		}
		fmt.Println(jsonOutput)
	} else {
		posts.RenderPostTable(postsList)
	}
}

func handleCreate(args []string, cfg config.Config) {
	if len(args) < 2 {
		fmt.Printf("%sError: Media URL required. Usage: insta create <media_url> [caption]%s\n", cli.ColorRed, cli.ColorReset)
		os.Exit(1)
	}

	mediaURL := args[1]
	caption := ""
	if len(args) > 2 {
		caption = strings.Join(args[2:], " ")
	}

	fmt.Printf("%sPublishing post to Instagram via Meta Graph API...%s\n", cli.ColorCyan, cli.ColorReset)
	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	postID, err := client.PublishMedia(mediaURL, caption)
	if err != nil {
		fmt.Printf("%sError publishing post: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	fmt.Printf("%s✓ Post published successfully! Post ID: %s%s\n", cli.ColorGreen, postID, cli.ColorReset)
}

func handleDelete(args []string, cfg config.Config) {
	if len(args) < 2 {
		fmt.Printf("%sError: Post ID is required. Usage: insta delete <post_id> [--force / -f]%s\n", cli.ColorRed, cli.ColorReset)
		os.Exit(1)
	}

	postID := ""
	force := false

	for i := 1; i < len(args); i++ {
		arg := args[i]
		if arg == "--force" || arg == "-f" {
			force = true
		} else if postID == "" {
			postID = arg
		}
	}

	if postID == "" {
		fmt.Printf("%sError: Post ID is required.%s\n", cli.ColorRed, cli.ColorReset)
		os.Exit(1)
	}

	if !force {
		fmt.Printf("%sAre you sure you want to delete post '%s'? [y/N]: %s", cli.ColorYellow, postID, cli.ColorReset)
		var response string
		fmt.Scanln(&response)
		response = strings.TrimSpace(strings.ToLower(response))
		if response != "y" && response != "yes" {
			fmt.Printf("%s⚠ Deletion cancelled by user.%s\n", cli.ColorYellow, cli.ColorReset)
			return
		}
	}

	fmt.Printf("%sInitiating Meta Graph API deletion for post ID '%s'...%s\n", cli.ColorCyan, postID, cli.ColorReset)
	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	if err := client.DeleteMedia(postID); err != nil {
		fmt.Printf("%sMeta Graph API Deletion Error: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	fmt.Printf("%s✓ Post '%s' deleted successfully via Meta Graph API.%s\n", cli.ColorGreen, postID, cli.ColorReset)
}

func handleVerify(cfg config.Config) {
	fmt.Printf("%s--- Executing Meta Graph API Integration Verification ---%s\n\n", cli.ColorCyan, cli.ColorReset)

	if !cfg.IsLoggedIn || strings.TrimSpace(cfg.AccessToken) == "" {
		fmt.Printf("%sFAIL: Meta Access Token not configured.%s\n", cli.ColorRed, cli.ColorReset)
		os.Exit(1)
	}

	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	user, err := client.FetchUserProfile()
	if err != nil {
		fmt.Printf("%sFAIL: User profile fetch failed: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	fmt.Printf("%s✓ Verified User Profile: %s (ID: %s)%s\n", cli.ColorGreen, user.Username, user.ID, cli.ColorReset)

	postsList, err := client.FetchUserPosts(5)
	if err != nil {
		fmt.Printf("%sFAIL: User posts fetch failed: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	fmt.Printf("%s✓ Verified User Posts Retrieval (%d posts fetched)%s\n", cli.ColorGreen, len(postsList), cli.ColorReset)
	fmt.Printf("\n%s✓ Meta Graph API Integration 100%% Operational.%s\n", cli.ColorGreen, cli.ColorReset)
}
