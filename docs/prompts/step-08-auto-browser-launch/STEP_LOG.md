# Step 8: 8-Phase Architectural Plan & Audit Log

- **Target Step**: Step 8 - Auto Browser Launch & CLI-Web Sync Engine
- **Status**: PLAN_APPROVED_READY_FOR_IMPLEMENTATION
- **Timestamp**: 2026-10-07

---

## Phase 1 — Understand
- Cross-platform default browser launch (`Windows`, `macOS`, `Linux`) when executing `insta show ui`.
- Non-blocking ANSI colored fallback terminal link output.
- Frontend periodic background polling (`setInterval(loadPosts, 10000)`) in `pkg/server/ui.go` for real-time CLI-Web data sync.
- `RequireAuth()` guard enforcement.

---

## Phase 2 — Audit of Existing Work
- `main.go`: Subcommand `show ui` entry point.
- `pkg/server/server.go`: Embedded HTTP server and REST endpoints (`/api/status`, `/api/posts`, `/api/delete`, `/api/edit`).
- `pkg/server/ui.go`: Embedded single-page Web UI (`IndexHTML`).
- `pkg/cli/colors.go`: ANSI color constants.
- `pkg/auth/auth.go`: `RequireAuth(cfg)` protection guard.

---

## Phase 3 — Gap Analysis
- [ ] Task 1: `pkg/cli/browser.go` cross-platform helper (`OpenBrowser`).
- [ ] Task 2: Server startup integration with browser launch Goroutine & ANSI fallback link in `pkg/server/server.go`.
- [ ] Task 3: Background JS polling in `pkg/server/ui.go`.

---

## Phase 4 — Task Breakdown
1. **Task 1: `pkg/cli/browser.go` Helper**: Implement `OpenBrowser(url string) error` using `exec.Command` for `windows`, `darwin`, and `linux`.
2. **Task 2: Server Integration**: Call `cli.OpenBrowser` in `server.go` non-blocking goroutine upon HTTP server startup.
3. **Task 3: JS Auto-Refresh Polling**: Add `setInterval(loadPosts, 10000)` in `ui.go`.

---

## Phase 5 — Requirement Mapping
- Requirement 1 & 2 -> Task 1 & Task 2
- Requirement 3 -> Task 3
- Requirement 4 -> Task 2 (`RequireAuth` reuse)

---

## Phase 6 — Reuse & Duplication Check
- 100% reuse of existing HTTP server, REST endpoints, ANSI colors, and auth guards.

---

## Phase 7 — Validation Plan
- Test `insta show ui` auto-browser opening.
- Verify fallback ANSI link output in terminal.
- Test background polling auto-refresh when post data is modified.
