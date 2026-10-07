# Instagram Management CLI with Local Web UI

A lightweight, hybrid Go-based Command Line Interface (CLI) and Local Web Dashboard Application designed to manage Instagram accounts, view posts, delete content, and perform bulk operations efficiently.

---

## Table of Contents
- [What is CLI?](#what-is-cli)
- [What is this Project?](#what-is-this-project)
- [Goal](#goal)
- [Problem & Solution](#problem--solution)
- [Core Features](#core-features)
- [Key Architectural Questions](#key-architectural-questions)
- [Trade-offs & UX Guidelines](#trade-offs--ux-guidelines)
- [Success Criteria](#success-criteria)
- [Final Purpose](#final-purpose)
- [10-Step Implementation Roadmap](#10-step-implementation-roadmap)
- [Getting Started & Local Usage](#getting-started--local-usage)

---

## What is CLI?
A Command Line Interface (CLI) is a text-based user interface used to run programs, manage files, and interact with computer systems directly via terminal commands. CLIs are fast, scriptable, resource-efficient, and favored by power users and developers.

---

## What is this Project?
Instagram Management CLI with Local Web UI (`go-insta-cli`) is a hybrid productivity tool built in Go (GoLang). It allows users to interact with Instagram directly from their terminal using simple commands, while also providing an on-demand, lightweight local web dashboard (`insta show ui`) for interactive visual management.

---

## Goal
Make Instagram management simple, fast, safe, and effortless—giving users full control via terminal keyboard shortcuts or a local browser interface without heavy web app bloat.

---

## Problem & Solution

### The Problem
* Official Instagram web and mobile apps are heavy, resource-intensive, and lack fast bulk-management tools.
* Power users and developers working in terminals have to constantly switch context between command-line tools and web browsers.
* Removing multiple posts manually is tedious, slow, and error-prone.

### The Solution
* **Hybrid Dual Interface**: Full terminal command control for quick tasks, plus a local browser web UI when visual grid selection is needed.
* **Automation & Safety**: Built-in interactive confirmation prompts, local session encryption, and rate-limiting to protect account health.

---

## Core Features
- `login`: Interactive login to authenticate and securely store session credentials.
- `logout`: Clear active session credentials and log out securely.
- `status`: View current active login state, username, and session timestamp.
- `posts [options]`: View recent Instagram posts with likes, comments, captions, `--limit <n>`, and `--json` flags (requires authentication).
- `delete <post_id>`: Remove specific posts safely with interactive `[y/N]` confirmation (requires authentication).
- `show ui`: Launch an embedded local web dashboard (`http://localhost:8080`) on demand (requires authentication).
- **Colorized Output**: High-visibility ANSI terminal formatting (Green for success, Yellow for warnings, Red for errors).
- **Zero Heavy Dependencies**: Built using Go standard libraries for high execution speed and minimal memory footprint.

---

## Key Architectural Questions
1. **How is authentication stored?** Encrypted locally at `~/.config/insta-cli/config.json` with strict `0600` file permissions.
2. **How does real-time sync work?** Two-way synchronization between CLI commands and Web UI state.
3. **How are rate limits handled?** Integrated request throttling and exponential backoff to prevent Instagram account flags.
4. **Is cross-platform supported?** Yes, compiles into a single binary for Windows, Linux, and macOS.

---

## Trade-offs & UX Guidelines

| Area | Good UX (Implemented) | Bad UX (Avoided) |
| :--- | :--- | :--- |
| **Authentication** | Encrypted local session storage & auth guard | Prompting password on every single command |
| **Post Deletion** | Interactive confirmation prompt `[y/N]` & `--force` flag | Instant unconfirmed deletion |
| **Local Web UI** | Auto-opening default browser upon `show ui` | Forcing user to manually type URL |
| **API Safety** | Built-in request delay & throttling | Rapid request spamming triggering account bans |

---

## Success Criteria
- Connects and authenticates with Instagram safely.
- Displays accurate post metrics in terminal and web UI.
- Deletes posts securely with proper user confirmation.
- Launches local web UI within seconds upon `show ui`.
- Protects Instagram account health via rate-limit management.

---

## Final Purpose
> Connect once, control anywhere—manage Instagram effortlessly via terminal or local web dashboard.

---

## 10-Step Implementation Roadmap

- [x] **Step 1: Basic CLI & Subcommand Skeleton** *(Completed)*
- [x] **Step 2: Secure Local Configuration & Storage** *(Completed)*
- [x] **Step 3: Instagram Authentication Module** *(Completed)*
- [x] **Step 4: Post Retrieval Engine** *(Completed)*
- [ ] **Step 5: Post Deletion Engine & Safety Mechanism**
- [ ] **Step 6: Embedded Local Web Server Core**
- [ ] **Step 7: Interactive Web Dashboard UI**
- [ ] **Step 8: Auto Browser Launch & CLI-Web Sync Engine**
- [ ] **Step 9: API Rate Limiting & Account Protection Controls**
- [ ] **Step 10: Packaging, Cross-Platform Distribution & CI/CD**

---

## Getting Started & Local Usage

### Prerequisites
- Go 1.20 or higher installed.

### Installation & Execution
```powershell
# 1. Clone the repository
git clone https://github.com/mostafizdev01/go-insta-cli.git
cd go-insta-cli

# 2. Check CLI Version
insta --version

# 3. Display Usage Guide
insta --help

# 4. Check Session Status
insta status

# 5. Authenticate & Save Session
insta login

# 6. Fetch Posts (with limit and json flags)
insta posts
insta posts --limit 2
insta posts --json

# 7. Test Logout
insta logout
```

---

## Verification

### Setup Requirements
- Meta for Developers Account ([developers.facebook.com](https://developers.facebook.com/))
- Configured Instagram Graph API Access Token (`INSTAGRAM_ACCESS_TOKEN`)
- Connected Instagram Account ID (`INSTAGRAM_ACCOUNT_ID`)

### Required Environment Variables
See `.env.example` for reference:
```bash
INSTAGRAM_ACCESS_TOKEN=EAA...
INSTAGRAM_ACCOUNT_ID=17841...
```

### Exact Command to Run
```powershell
insta verify
```

### Expected Successful Output
```
Executing Meta Graph API End-to-End Integration Verification...

[PASSED] Reached Meta Graph API successfully.
  User ID:   17841458817135842
  Username:  mostafizdev01

[PASSED] Retrieved 2 real Instagram posts from account.

[1] Post ID: 1792348719234 (2026-10-07 14:00)
    Caption:  Building a self-updating CLI with Meta Graph API!
    Type:     IMAGE
    Metrics:  12 Likes | 3 Comments
    --------------------------------------------------

[VERIFICATION SUCCESSFUL] Real Instagram API integration verified successfully!
```

### Common Failure Messages and Meaning
1. `INSTAGRAM_ACCESS_TOKEN is missing`: The `INSTAGRAM_ACCESS_TOKEN` environment variable or config file is not configured.
2. `Meta API Error (190): Invalid OAuth access token`: The access token is invalid, expired, or revoked.
3. `Meta API Error (10): Permission Error`: The Meta access token lacks required `instagram_basic` or `pages_show_list` permissions.
4. `Implementation is not fully verified because real Instagram API access is not configured.`: Printed when credentials/permissions are missing or incomplete.

---

Maintained by [@mostafizdev01](https://github.com/mostafizdev01)

