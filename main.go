package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Version represents the CLI version, overridable at build time via:
// go build -ldflags="-X main.Version=v1.0.0"
var Version = "v1.0.0"

// ANSI Color Constants for Terminal Output
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "%sInstagram Management CLI (%sgo-insta-cli%s)%s\n\n", ColorCyan, ColorGreen, ColorCyan, ColorReset)
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  go-insta-cli [flags]\n")
	fmt.Fprintf(os.Stderr, "  go-insta-cli [command]\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	fmt.Fprintf(os.Stderr, "  -v, --version    Show current version\n")
	fmt.Fprintf(os.Stderr, "  -h, --help       Show help usage guide\n\n")
	fmt.Fprintf(os.Stderr, "Subcommands:\n")
	fmt.Fprintf(os.Stderr, "  login            Initiate Instagram login process\n")
	fmt.Fprintf(os.Stderr, "  posts            Fetch recent Instagram posts\n")
	fmt.Fprintf(os.Stderr, "  delete <id>      Delete a specific Instagram post\n")
	fmt.Fprintf(os.Stderr, "  show ui          Start local Web UI server\n")
}

func main() {
	showVersion := flag.Bool("version", false, "Show current version")
	flag.BoolVar(showVersion, "v", false, "Show current version (shorthand)")

	showHelp := flag.Bool("help", false, "Show help usage guide")
	flag.BoolVar(showHelp, "h", false, "Show help usage guide (shorthand)")

	flag.Usage = printUsage

	flag.Parse()

	if *showVersion {
		fmt.Printf("%sgo-insta-cli version %s%s\n", ColorGreen, Version, ColorReset)
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

	// Handle multi-word subcommand: "show ui"
	if command == "show" && len(args) > 1 && strings.ToLower(args[1]) == "ui" {
		command = "show ui"
	}

	switch command {
	case "login":
		fmt.Printf("%sInitiating Instagram login process...%s\n", ColorYellow, ColorReset)
	case "posts":
		fmt.Printf("%sFetching Instagram posts...%s\n", ColorCyan, ColorReset)
	case "delete":
		fmt.Printf("%sUsage: delete <post_id>%s\n", ColorYellow, ColorReset)
	case "show ui":
		fmt.Printf("%sStarting local web UI server...%s\n", ColorGreen, ColorReset)
	default:
		fmt.Fprintf(os.Stderr, "%sError: Unknown command '%s'%s\n\n", ColorRed, strings.Join(args, " "), ColorReset)
		printUsage()
		os.Exit(1)
	}
}
