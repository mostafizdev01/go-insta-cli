# Step 6: Requirements & Acceptance Criteria

## Target Scope
Implement the Embedded Local Web Server Core for `go-insta-cli`, enabling authenticated users to launch a lightweight HTTP REST API server (`http://localhost:8080`) via `insta show ui`.

---

## Detailed Requirements

1. **Embedded Web Server (`pkg/server/server.go`)**:
   - Use Go standard library `net/http` exclusively.
   - Implement HTTP server routing with `http.NewServeMux()`.
   - Provide graceful shutdown with `context.WithTimeout` handling `SIGINT`/`SIGTERM`.

2. **REST API Endpoints**:
   - `GET /api/status`: Return authentication & session status JSON (`status`, `user_id`, `username`, `last_login`).
   - `GET /api/posts`: Return recent Instagram posts list as JSON (reusing Step 4 post retrieval).
   - `POST /api/delete`: Delete specified post ID passed in JSON body or query param (reusing Step 5 post deletion).

3. **Security & Authentication Guard**:
   - Protect `insta show ui` subcommand and REST API endpoints using `auth.RequireAuth(cfg)`.
   - Return `HTTP 401 Unauthorized` for unauthenticated requests.

4. **Terminal Feedback**:
   - Display Green message: `✓ Local Web Server running at http://localhost:8080 (Press Ctrl+C to stop)`.
   - Display Yellow shutdown notification upon `Ctrl+C`.

---

## Verification Criteria
- Unauthenticated invocation `insta show ui` is blocked by `RequireAuth()`.
- Executing `insta show ui` launches HTTP server on `http://localhost:8080`.
- `curl http://localhost:8080/api/status` returns `HTTP 200 OK` JSON with user profile data.
- `curl http://localhost:8080/api/posts` returns `HTTP 200 OK` JSON post list or Meta API status.
- `curl -X POST http://localhost:8080/api/delete?id=12345` invokes Meta Graph API deletion logic.
- Pressing `Ctrl+C` triggers graceful shutdown without resource leaks.
