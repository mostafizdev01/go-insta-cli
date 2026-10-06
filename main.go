package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

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
	fmt.Fprintf(os.Stderr, "  login <user>     Initiate Instagram login and save session\n")
	fmt.Fprintf(os.Stderr, "  status           Display current login and session status\n")
	fmt.Fprintf(os.Stderr, "  posts            Fetch recent Instagram posts\n")
	fmt.Fprintf(os.Stderr, "  delete <id>      Delete a specific Instagram post\n")
	fmt.Fprintf(os.Stderr, "  show ui          Start local Web UI server\n")
}

func handleLogin(args []string) {
	username := "mostafizdev01"
	if len(args) > 1 && args[1] != "" {
		username = args[1]
	}

	cfg := config.Config{
		Username:     username,
		SessionToken: "mock_session_token_" + username,
		IsLoggedIn:   true,
		LastLogin:    time.Now().Format(time.RFC3339),
	}

	if err := config.SaveConfig(cfg); err != nil {
		fmt.Printf("%sError saving session: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		os.Exit(1)
	}

	fmt.Printf("%sInitiating Instagram login process...%s\n", cli.ColorYellow, cli.ColorReset)
	fmt.Printf("%sSuccessfully logged in and saved session for '%s'.%s\n", cli.ColorGreen, username, cli.ColorReset)
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
		fmt.Printf("  Message:    Run 'insta login <username>' to authenticate.\n")
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

	switch command {
	case "login":
		handleLogin(args)
	case "status":
		handleStatus()
	case "posts":
		fmt.Printf("%sFetching Instagram posts...%s\n", cli.ColorCyan, cli.ColorReset)
	case "delete":
		fmt.Printf("%sUsage: delete <post_id>%s\n", cli.ColorYellow, cli.ColorReset)
	case "show ui":
		fmt.Printf("%sStarting local web UI server...%s\n", cli.ColorGreen, cli.ColorReset)
	default:
		fmt.Fprintf(os.Stderr, "%sError: Unknown command '%s'%s\n\n", cli.ColorRed, strings.Join(args, " "), cli.ColorReset)
		printUsage()
		os.Exit(1)
	}
}
