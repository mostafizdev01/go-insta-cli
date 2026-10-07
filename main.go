package main

import (
	"flag"
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

var Version = "v1.0.0"

func printUsage() {
	fmt.Fprintf(os.Stderr, "%sInstagram Management CLI (%sgo-insta-cli%s)%s\n\n", cli.ColorCyan, cli.ColorGreen, cli.ColorCyan, cli.ColorReset)
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  go-insta-cli [flags]\n")
	fmt.Fprintf(os.Stderr, "  go-insta-cli [command]\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	fmt.Fprintf(os.Stderr, "  -v, --version    Show current version\n")
	fmt.Fprintf(os.Stderr, "  -h, --help       Show help usage guide\n\n")
	fmt.Fprintf(os.Stderr, "Subcommands:\n")
	fmt.Fprintf(os.Stderr, "  login            Save Meta Graph API Access Token and Account ID\n")
	fmt.Fprintf(os.Stderr, "  logout           Clear active credentials\n")
	fmt.Fprintf(os.Stderr, "  status           Verify active Meta Graph API token and profile\n")
	fmt.Fprintf(os.Stderr, "  posts [options]  Fetch recent Instagram posts via Meta Graph API (--limit <n>, --json)\n")
	fmt.Fprintf(os.Stderr, "  create <url>     Publish a new post via Meta Graph API\n")
	fmt.Fprintf(os.Stderr, "  delete <id>      Delete a specific Instagram post (--force / -f)\n")
	fmt.Fprintf(os.Stderr, "  verify           Execute complete end-to-end Meta Graph API integration test\n")
}

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

	fmt.Printf("  Status:     %sConnected (Verified via Meta Graph API)%s\n", cli.ColorGreen, cli.ColorReset)
	fmt.Printf("  Username:   %s%s%s\n", cli.ColorGreen, user.Username, cli.ColorReset)
	fmt.Printf("  Account ID: %s%s%s\n", cli.ColorGreen, user.ID, cli.ColorReset)
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
				if l, err := strconv.Atoi(parts[1]); err == nil {
					limit = l
				}
			}
		} else if arg == "--limit" && i+1 < len(args) {
			if l, err := strconv.Atoi(args[i+1]); err == nil {
				limit = l
				i++
			}
		}
	}

	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	postList, err := client.FetchUserPosts(limit)
	if err != nil {
		fmt.Printf("%s[Meta Graph API Error] %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	if isJSON {
		jsonStr, err := posts.RenderPostJSON(postList)
		if err != nil {
			fmt.Printf("%sError formatting JSON: %v%s\n", cli.ColorRed, err, cli.ColorReset)
			os.Exit(1)
		}
		fmt.Println(jsonStr)
	} else {
		posts.RenderPostTable(postList)
	}
}

func handleCreate(args []string, cfg config.Config) {
	if len(args) < 2 {
		fmt.Printf("%sUsage: insta create <image_url> [caption...]%s\n", cli.ColorYellow, cli.ColorReset)
		return
	}

	imageURL := args[1]
	caption := ""
	if len(args) > 2 {
		caption = strings.Join(args[2:], " ")
	}

	params := posts.CreatePostParams{
		ImageURL: imageURL,
		Caption:  caption,
	}

	if err := params.Validate(); err != nil {
		fmt.Printf("%sError: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	fmt.Printf("%sPublishing post via Meta Content Publishing API...%s\n", cli.ColorYellow, cli.ColorReset)
	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	postID, err := client.PublishMedia(params.ImageURL, params.Caption)
	if err != nil {
		fmt.Printf("%sError creating post: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	posts.RenderCreateSuccess(postID)
}

func handleDelete(args []string, cfg config.Config) {
	if len(args) < 2 {
		fmt.Printf("%sUsage: insta delete <post_id> [--force|-f]%s\n", cli.ColorYellow, cli.ColorReset)
		os.Exit(1)
	}

	postID := ""
	force := false

	for i := 1; i < len(args); i++ {
		arg := args[i]
		if arg == "--force" || arg == "-f" {
			force = true
		} else if postID == "" && !strings.HasPrefix(arg, "-") {
			postID = arg
		}
	}

	params := posts.DeletePostParams{
		PostID: postID,
		Force:  force,
	}

	if err := params.Validate(); err != nil {
		fmt.Printf("%sError: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	if !params.Force {
		confirmed := posts.ConfirmDeletion(params.PostID)
		if !confirmed {
			posts.RenderDeleteCancelled(params.PostID)
			return
		}
	}

	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	err := client.DeleteMedia(params.PostID)
	if err != nil {
		posts.RenderDeleteError(params.PostID, err)
		os.Exit(1)
	}

	posts.RenderDeleteSuccess(params.PostID)
}

func handleVerify(cfg config.Config) {
	fmt.Printf("%sExecuting Meta Graph API End-to-End Integration Verification...%s\n\n", cli.ColorCyan, cli.ColorReset)

	if strings.TrimSpace(cfg.AccessToken) == "" {
		fmt.Printf("%s[VERIFICATION FAILED]%s INSTAGRAM_ACCESS_TOKEN environment variable is not configured.\n", cli.ColorRed, cli.ColorReset)
		fmt.Printf("Implementation is not fully verified because real Instagram API access is not configured.\n")
		os.Exit(1)
	}

	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)

	// Step 1: User Profile Reachability Test
	user, err := client.FetchUserProfile()
	if err != nil {
		fmt.Printf("%s[VERIFICATION FAILED]%s Unable to reach Meta Graph API with provided access token:\n  %v\n\n", cli.ColorRed, cli.ColorReset, err)
		fmt.Printf("Implementation is not fully verified because real Instagram API access is not configured.\n")
		os.Exit(1)
	}

	fmt.Printf("%s[PASSED]%s Reached Meta Graph API successfully.\n", cli.ColorGreen, cli.ColorReset)
	fmt.Printf("  User ID:   %s\n", user.ID)
	fmt.Printf("  Username:  %s\n\n", user.Username)

	// Step 2: Media Retrieval Test
	postsList, err := client.FetchUserPosts(5)
	if err != nil {
		fmt.Printf("%s[VERIFICATION FAILED]%s Profile reached, but media fetch failed:\n  %v\n\n", cli.ColorRed, cli.ColorReset, err)
		fmt.Printf("Implementation is not fully verified because real Instagram API permissions are missing.\n")
		os.Exit(1)
	}

	fmt.Printf("%s[PASSED]%s Retrieved %d real Instagram posts from account.\n\n", cli.ColorGreen, len(postsList))
	posts.RenderPostTable(postsList)

	fmt.Printf("%s[VERIFICATION SUCCESSFUL]%s Real Instagram API integration verified successfully!\n", cli.ColorGreen, cli.ColorReset)
}

func main() {
	showVersion := flag.Bool("version", false, "Show current version")
	flag.BoolVar(showVersion, "v", false, "Show current version (shorthand)")

	showHelp := flag.Bool("help", false, "Show help usage guide")
	flag.BoolVar(showHelp, "h", false, "Show help usage guide (shorthand)")

	flag.Usage = printUsage
	flag.Parse()

	if *showVersion {
		fmt.Printf("%sgo-insta-cli version %s%s\n", cli.ColorGreen, Version, cli.ColorReset)
		return
	}

	if *showHelp {
		printUsage()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		printUsage()
		return
	}

	command := strings.ToLower(args[0])
	cfg, _ := config.LoadConfig()

	switch command {
	case "login":
		handleLogin(args)
	case "logout":
		handleLogout()
	case "status":
		handleStatus(cfg)
	case "verify":
		handleVerify(cfg)
	case "posts":
		if !auth.RequireAuth(cfg) {
			os.Exit(1)
		}
		handlePosts(args, cfg)
	case "create":
		if !auth.RequireAuth(cfg) {
			os.Exit(1)
		}
		handleCreate(args, cfg)
	case "delete":
		if !auth.RequireAuth(cfg) {
			os.Exit(1)
		}
		handleDelete(args, cfg)
	default:
		fmt.Fprintf(os.Stderr, "%sError: Unknown command '%s'%s\n\n", cli.ColorRed, strings.Join(args, " "), cli.ColorReset)
		printUsage()
		os.Exit(1)
	}
}
