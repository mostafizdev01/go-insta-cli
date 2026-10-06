package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"go-insta-cli/pkg/auth"
	"go-insta-cli/pkg/cli"
	"go-insta-cli/pkg/config"
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
	fmt.Fprintf(os.Stderr, "  login            Initiate Instagram interactive login and save session\n")
	fmt.Fprintf(os.Stderr, "  logout           Clear active session and log out\n")
	fmt.Fprintf(os.Stderr, "  status           Display current login and session status\n")
	fmt.Fprintf(os.Stderr, "  posts            Fetch recent Instagram posts (requires login)\n")
	fmt.Fprintf(os.Stderr, "  delete <id>      Delete a specific Instagram post (requires login)\n")
	fmt.Fprintf(os.Stderr, "  show ui          Start local Web UI server (requires login)\n")
}

func handleLogin(args []string) {
	username := ""
	sessionToken := ""

	if len(args) > 1 {
		username = args[1]
	}
	if len(args) > 2 {
		sessionToken = args[2]
	}

	fmt.Printf("%sInitiating Instagram login process...%s\n", cli.ColorYellow, cli.ColorReset)
	cfg, err := auth.PerformLogin(username, sessionToken)
	if err != nil {
		fmt.Printf("%sError during login: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	fmt.Printf("%sSuccessfully authenticated and saved session for '%s'.%s\n", cli.ColorGreen, cfg.Username, cli.ColorReset)
}

func handleLogout() {
	if err := auth.PerformLogout(); err != nil {
		fmt.Printf("%sError during logout: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}
	fmt.Printf("%sLogged out successfully. Session cleared.%s\n", cli.ColorYellow, cli.ColorReset)
}

func handleStatus() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("%sError loading status: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	fmt.Printf("%sInstagram CLI Session Status:%s\n", cli.ColorCyan, cli.ColorReset)
	if cfg.IsLoggedIn {
		fmt.Printf("  Status:     %sLogged In%s\n", cli.ColorGreen, cli.ColorReset)
		fmt.Printf("  Username:   %s%s%s\n", cli.ColorGreen, cfg.Username, cli.ColorReset)
		fmt.Printf("  Last Login: %s\n", cfg.LastLogin)
	} else {
		fmt.Printf("  Status:     %sLogged Out%s\n", cli.ColorYellow, cli.ColorReset)
		fmt.Printf("  Message:    Run 'insta login' to authenticate.\n")
	}
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
	if command == "show" && len(args) > 1 && strings.ToLower(args[1]) == "ui" {
		command = "show ui"
	}

	cfg, _ := config.LoadConfig()

	switch command {
	case "login":
		handleLogin(args)
	case "logout":
		handleLogout()
	case "status":
		handleStatus()
	case "posts":
		if !auth.RequireAuth(cfg) {
			os.Exit(1)
		}
		fmt.Printf("%sFetching Instagram posts...%s\n", cli.ColorCyan, cli.ColorReset)
	case "delete":
		if !auth.RequireAuth(cfg) {
			os.Exit(1)
		}
		fmt.Printf("%sUsage: delete <post_id>%s\n", cli.ColorYellow, cli.ColorReset)
	case "show ui":
		if !auth.RequireAuth(cfg) {
			os.Exit(1)
		}
		fmt.Printf("%sStarting local web UI server...%s\n", cli.ColorGreen, cli.ColorReset)
	default:
		fmt.Fprintf(os.Stderr, "%sError: Unknown command '%s'%s\n\n", cli.ColorRed, strings.Join(args, " "), cli.ColorReset)
		printUsage()
		os.Exit(1)
	}
}
