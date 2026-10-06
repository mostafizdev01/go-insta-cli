# Step 2: 8-Phase Architectural Plan & Audit Log

- **Target Step**: Step 2 - Secure Configuration & Local Storage
- **Status**: PLANNING_APPROVED (Ready for Code Implementation)

---

## Phase 2 — Audit of Step 1 Codebase (`main.go`)
- **Reusable Elements**:
  - ANSI Color constants (`ColorReset`, `ColorGreen`, `ColorYellow`, `ColorRed`, `ColorCyan`).
  - Flag parsing (`--version`, `--help`).
  - Subcommand router (`switch command`).
  - Build-time version variable `Version = "v1.0.0"`.

---

## Phase 3 — Gap Analysis
- [x] **Step 1 Core CLI Skeleton**: 100% Completed.
- [ ] **Config Struct & Storage Engine**: 0% (Target for Step 2).
- [ ] **Token Obfuscation & File Permissions (0600)**: 0% (Target for Step 2).
- [ ] **`login <username>` Integration**: Partial (Step 1 placeholder text only; Needs config persistence).
- [ ] **`status` Subcommand**: 0% (Target for Step 2).

---

## Phase 4 & 8 — Step 2 Implementation Plan

### Task 2.1: Define Config Data Structures & Helpers
- Define `Config` struct (`Username`, `SessionToken`, `IsLoggedIn`, `LastLogin`).
- Implement `GetConfigFilePath()`, `encodeToken()`, `decodeToken()`.

### Task 2.2: Implement File I/O Engine (`LoadConfig` & `SaveConfig`)
- Implement `LoadConfig()` (reads JSON from disk, returns default empty config if missing).
- Implement `SaveConfig(cfg Config)` (marshals JSON, writes to file with `0600` permissions).

### Task 2.3: Integrate Subcommands (`login` & `status`)
- Update `login` command: Parse username from args, save mock session token to `config.json`.
- Implement `status` command: Read `config.json` and print formatted ANSI status summary.

### Task 2.4: Empirical Testing & Rebuild
- Test `insta login testuser` and verify `~/.config/insta-cli/config.json` creation.
- Test `insta status` output.
- Rebuild & reinstall CLI (`go install` & `go build`).
