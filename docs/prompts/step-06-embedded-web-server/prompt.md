# Step 6 Prompt: Embedded Local Web Server Core

## System Role & Objective
Act as a Principal Software Architect and Lead Developer. Implement and verify Step 6 (**Embedded Local Web Server Core**) for `go-insta-cli` using Go's standard library `net/http`.

## Implementation Objectives
1. Create `pkg/server/server.go` implementing embedded HTTP web server and REST API endpoints (`GET /api/status`, `GET /api/posts`, `POST /api/delete`).
2. Update `main.go` to handle `show ui` subcommand and launch server on `http://localhost:8080`.
3. Protect web server and REST API endpoints using `auth.RequireAuth(cfg)`.
4. Provide ANSI colorized terminal startup message: `✓ Local Web Server running at http://localhost:8080 (Press Ctrl+C to stop)`.
5. Implement graceful HTTP server shutdown using `context.WithTimeout` upon `Ctrl+C` (SIGINT/SIGTERM) signals.
6. Strictly avoid heavy third-party web frameworks (no Gin, Echo, or Fiber). Keep every file under 300 lines of clean Go code.
