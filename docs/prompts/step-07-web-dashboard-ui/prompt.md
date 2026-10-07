# Step 7 Prompt: Interactive Web Dashboard UI

## System Role & Objective
Act as a Principal Software Architect and Lead Developer. Implement and verify Step 7 (**Interactive Web Dashboard UI**) for `go-insta-cli` using embedded HTML5, CSS3, and JavaScript served from Go `net/http`.

## Implementation Objectives
1. Create `pkg/server/ui.go` containing embedded single-page Web UI (`IndexHTML`) with dark aesthetic, clean typography, and zero external dependencies.
2. Render Instagram post cards dynamically via client-side JavaScript fetching `/api/posts` and `/api/status`.
3. Provide interactive post management capabilities:
   - Single post deletion buttons (`deleteSingle`).
   - Bulk selection checkboxes and batch delete button (`deleteSelected`).
   - Select All / Deselect All actions.
   - Dynamic toast notifications for API success and error feedback.
4. Integrate route `http.HandleFunc("/", srv.handleIndex)` into `pkg/server/server.go`.
5. Keep all Go source files strictly under 300 lines.
