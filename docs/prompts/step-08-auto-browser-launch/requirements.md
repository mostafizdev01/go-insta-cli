# Step 8: Requirements & Acceptance Criteria

## Objective
Implement Cross-Platform Auto Browser Launch and Real-Time Data Sync Engine for `go-insta-cli`, enabling `insta show ui` to automatically launch the default OS web browser pointing to `http://localhost:8080/`, print a non-blocking terminal fallback link, and automatically refresh the dashboard state when data changes.

---

## Detailed Requirements

1. **Cross-Platform Auto Browser Launcher (`pkg/cli/browser.go`)**:
   - Automatically open `http://localhost:8080/` in the OS default browser upon executing `insta show ui`.
   - Support platform-specific browser launch commands:
     - Windows: `cmd /c start http://localhost:8080/`
     - macOS: `open http://localhost:8080/`
     - Linux: `xdg-open http://localhost:8080/`

2. **Non-Blocking Terminal Fallback Link**:
   - If auto browser launch fails or runs in a headless environment, gracefully fall back to printing ANSI colorized clickable terminal link: `Open Dashboard: http://localhost:8080/`.

3. **Real-Time Data Sync Engine**:
   - Integrate automatic periodic background polling (`setInterval`) in `pkg/server/ui.go` JavaScript frontend.
   - Automatically refresh post list state every 10 seconds without interfering with active user interactions (like text entry).

4. **Authentication Protection**:
   - Enforce `auth.RequireAuth(cfg)` check before starting HTTP server or triggering browser launcher.

---

## Acceptance Criteria
- Executing `insta show ui` opens the system default browser to `http://localhost:8080/`.
- Terminal output displays ANSI colorized server running message and clickable URL link.
- Background polling refreshes the web dashboard grid automatically without requiring manual page reload.
- All Go files remain strictly under 300 lines of code.
