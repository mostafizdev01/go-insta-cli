# Step 3: 8-Phase Architectural Plan & Audit Log

- **Target Step**: Step 3 - Instagram Authentication Module
- **Status**: PLANNING_APPROVED (Ready for Code Implementation)

---

## Phase 2 — Audit of Step 2 Codebase
- **Completed Components**:
  - `pkg/cli/colors.go`: ANSI terminal color formatting.
  - `pkg/config/config.go`: Config struct, Base64 session token obfuscation, JSON file I/O engine (`LoadConfig`, `SaveConfig`) with `0600` permissions.
  - `main.go`: Basic flag parsing and `login` / `status` subcommand skeletons.

---

## Phase 3 — Gap Analysis
- [x] **Step 1 Core CLI Skeleton**: 100% Completed.
- [x] **Step 2 Local Config & Encryption**: 100% Completed.
- [ ] **Interactive Terminal Login Prompt**: 0% (Target for Step 3).
- [ ] **`pkg/auth/auth.go` Module & `RequireAuth()` Guard**: 0% (Target for Step 3).
- [ ] **`insta logout` Subcommand**: 0% (Target for Step 3).
- [ ] **Protecting `posts`, `delete`, `show ui` Subcommands**: 0% (Target for Step 3).

---

## Phase 4 & 8 — Step 3 Implementation Plan

### Task 3.1: Create Auth Package (`pkg/auth/auth.go`)
- Create new file `pkg/auth/auth.go` (under 150 lines).
- Implement `RequireAuth(cfg config.Config) bool` guard.
- Implement `PerformLogin(username, passwordOrCookie string) (config.Config, error)`.
- Implement `PerformLogout() error`.

### Task 3.2: Integrate Auth Guard in `main.go`
- Update `main.go` to import `go-insta-cli/pkg/auth`.
- Add `logout` subcommand to switch case and usage guide.
- Wrap `posts`, `delete`, and `show ui` handlers with `auth.RequireAuth(cfg)`.

### Task 3.3: Interactive Login Prompt
- Enhance `handleLogin()` to prompt interactively for credentials when no argument is provided.

### Task 3.4: Empirical Testing & Rebuild
- Test `insta logout`.
- Test running `insta posts` when logged out (verify red error blocking).
- Test running `insta login` and checking `insta status`.
- Rebuild binary (`go build -o "$env:USERPROFILE\go\bin\insta.exe" main.go`).
