# Instagram Management CLI (`go-insta-cli`)

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey.svg)]()

`go-insta-cli` is a high-performance, self-updating, cross-platform CLI tool and embedded local Web Dashboard for managing Instagram posts, metrics, and publishing using official Meta Graph API v19.0 endpoints.

---

## Key Features

- **Official Meta Graph API v19.0 Integration**: 100% compliant with Meta Business/Professional endpoints with zero browser cookie scraping.
- **Embedded Web Dashboard UI (`insta show ui`)**: Single-page dark theme web interface served locally (`http://localhost:8080/`) with interactive post cards, selection checkboxes, and live toast notifications.
- **Auto Browser Launcher & CLI-Web Sync**: Automatically launches system default browser and auto-refreshes data in real-time.
- **API Rate Limiting & Account Protection**: Central request throttling (1.5s interval), sequential bulk deletion pauses (3s interval), exponential backoff, and circuit breaker controls to keep accounts safe.
- **Single Static Executable**: Embedded static assets (`//go:embed`) with zero external runtime dependencies.
- **Cross-Platform Distribution**: Pre-compiled binaries for Windows, Linux, and macOS (Intel & Apple Silicon).

---

## Installation

### Option 1: Download Pre-Compiled Binary (Recommended)
Download the latest standalone executable from [GitHub Releases](https://github.com/mostafizdev01/go-insta-cli/releases):

- **Windows**: `insta.exe`
- **Linux**: `insta-linux`
- **macOS (Apple Silicon M1/M2/M3)**: `insta-darwin-arm64`
- **macOS (Intel)**: `insta-darwin-amd64`

### Option 2: Install via Go CLI
```bash
go install github.com/mostafizdev01/go-insta-cli@latest
```

---

## Quick Start & Authentication Setup

### 1. Configure Credentials
Obtain your Meta Graph API Access Token (`INSTAGRAM_ACCESS_TOKEN`) and connected Instagram Account ID (`INSTAGRAM_ACCOUNT_ID`), then run:

```bash
insta login <access_token> <account_id>
```

Alternatively, set environment variables:
```powershell
$env:INSTAGRAM_ACCESS_TOKEN="EAAP86..."
$env:INSTAGRAM_ACCOUNT_ID="122093749821512653"
```

### 2. Verify Connection
```bash
insta status
```

---

## Command Reference

| Command | Description |
| :--- | :--- |
| `insta login <token> [id]` | Save Meta Graph API credentials securely |
| `insta logout` | Clear active stored credentials |
| `insta status` | Display connection, profile, and rate limit status |
| `insta posts [options]` | Fetch recent posts (`--limit <n>`, `--json`) |
| `insta create <url> [caption]` | Publish a new image post via Meta Graph API |
| `insta delete <id> [--force]` | Delete an Instagram post (`-f` / `--force` bypasses prompt) |
| `insta show ui` | Start local Web Server and launch Interactive Web Dashboard |
| `insta verify` | Execute end-to-end Meta Graph API integration test |
| `insta version` | Display version, build date, and git commit metadata |

---

## Local Web Dashboard (`insta show ui`)

Execute `insta show ui` in your terminal to start the embedded web server on `http://localhost:8080/`. The CLI will automatically launch your system default browser.

- **Interactive Post Cards**: View images, captions, timestamps, like counts, and comment counts.
- **Bulk Selection & Deletion**: Select multiple post cards using checkboxes and delete them sequentially with forced 3s safety intervals.
- **Live Caption Editing**: Click "Edit" on any card to update post captions.

---

## Rate Limiting & Account Protection

`go-insta-cli` includes a built-in rate limiter (`pkg/instagram/ratelimit.go`) that enforces human-like request pacing:

- **1.5s Safety Pause**: Inserted between consecutive API calls.
- **3s Sequential Bulk Interval**: Forced pause between bulk deletion requests.
- **Circuit Breaker**: Automatically trips for 60 seconds if HTTP 429 is encountered repeatedly.

---

## Building from Source & Cross-Compilation

To cross-compile binaries for all supported operating systems:

```powershell
# Run build automation script
powershell -ExecutionPolicy Bypass -File scripts/build.ps1 -Version "v1.0.0"
```

Output binaries and SHA-256 `checksums.txt` will be generated in the `dist/` directory.

---

## License

Distributed under the MIT License. See `LICENSE` for more information.
