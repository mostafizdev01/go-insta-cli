package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"go-insta-cli/pkg/auth"
	"go-insta-cli/pkg/cli"
	"go-insta-cli/pkg/config"
	"go-insta-cli/pkg/server"
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
	fmt.Fprintf(os.Stderr, "  show ui          Start embedded local Web UI server (http://localhost:8080)\n")
	fmt.Fprintf(os.Stderr, "  verify           Execute complete end-to-end Meta Graph API integration test\n")
}

func main() {
	showVersion := flag.Bool("version", false, "Show current version")
	flag.BoolVar(showVersion, "v", false, "Show current version (shorthand)")

	flag.Usage = printUsage
	flag.Parse()

	if *showVersion {
		fmt.Printf("go-insta-cli version %s\n", Version)
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		printUsage()
		return
	}

	command := args[0]
	if command == "show" && len(args) > 1 && args[1] == "ui" {
		command = "show ui"
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("%sWarning: Failed to load config file: %v%s\n", cli.ColorYellow, err, cli.ColorReset)
	}

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
	case "show ui":
		if !auth.RequireAuth(cfg) {
			os.Exit(1)
		}
		if err := server.StartServer(cfg, "8080"); err != nil {
			fmt.Printf("%sServer Error: %v%s\n", cli.ColorRed, err, cli.ColorReset)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "%sError: Unknown command '%s'%s\n\n", cli.ColorRed, strings.Join(args, " "), cli.ColorReset)
		printUsage()
		os.Exit(1)
	}
}
