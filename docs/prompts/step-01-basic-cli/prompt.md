# Step 1: Master AI Implementation Prompt

Act as a Senior Go Developer. I want to build Step 1 of an Instagram CLI tool in Go (GoLang).

Project Context:
- Module Name: go-insta-cli
- File: main.go

Requirements for Step 1:
1. Use only Go standard libraries (`flag`, `fmt`, `os`). Do not use heavy third-party libraries.
2. Define a global variable `Version = "v1.0.0"` that can be overridden at build time using `-ldflags`.
3. Implement ANSI color constants for terminal output (ColorReset, ColorRed, ColorGreen, ColorYellow, ColorCyan).
4. Support `--version` and `-v` flags to print colored version information.
5. Support `--help` and `-h` flags with a custom usage menu listing available flags and subcommands.
6. Support the skeleton structure for the following subcommands:
   - `login`: Print colored placeholder "Initiating Instagram login process...".
   - `posts`: Print colored placeholder "Fetching Instagram posts...".
   - `delete`: Print colored placeholder "Usage: delete <post_id>".
   - `show ui`: Print colored placeholder "Starting local web UI server...".
7. Gracefully handle unknown commands with a red error message, display usage guide, and exit with code 1.

Please provide:
1. The exact terminal command to initialize the module (`go mod init go-insta-cli`).
2. The complete, production-ready `main.go` code with clean formatting.
3. Step-by-step terminal commands to build and test `--version`, `--help`, `login`, `posts`, `delete`, and `show ui`.
