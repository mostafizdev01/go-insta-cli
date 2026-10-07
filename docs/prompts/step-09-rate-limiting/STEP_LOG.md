# Step 9: 8-Phase Architectural Plan & Audit Log

- **Target Step**: Step 9 - API Rate Limiting & Account Protection Controls
- **Status**: PLAN_APPROVED_READY_FOR_IMPLEMENTATION
- **Timestamp**: 2026-10-07

---

## Phase 1 — Understand
- Request throttling middleware with 1-2s safety delays.
- Sequential bulk post deletion with 3s interval between requests.
- Exponential backoff & Circuit Breaker on HTTP 429 errors.
- Rate limit metrics in `insta status` and `GET /api/status`.

---

## Phase 2 — Audit of Existing Work
- `pkg/instagram/client.go`: HTTP client wrapper for Meta Graph API calls.
- `pkg/instagram/fetch.go`, `delete.go`, `publish.go`: API query methods.
- `pkg/server/handlers.go`: Web REST API endpoints (`/api/status`, `/api/posts`, `/api/delete`).
- `cmd_handlers.go`: CLI command handlers.

---

## Phase 3 — Gap Analysis
- [ ] Task 1: `pkg/instagram/ratelimit.go` Rate Limiter Engine & Circuit Breaker.
- [ ] Task 2: Rate limiter integration into `Client.HTTPClient` round tripper.
- [ ] Task 3: Sequential bulk deletion throttling in CLI & Web UI.
- [ ] Task 4: Expose rate limit status in `insta status` & `/api/status`.

---

## Phase 4 — Task Breakdown
1. **Task 1: Rate Limiter Engine & Circuit Breaker**: Create `pkg/instagram/ratelimit.go` with request throttling, exponential backoff, and state tracking.
2. **Task 2: HTTP Client Interception**: Wrap Meta Graph API HTTP requests in `client.go` with rate limiting transport.
3. **Task 3: Bulk Action Throttling**: Update CLI bulk delete handler in `cmd_handlers.go` and JS handler in `ui_script.go` for sequential execution with forced 3s delay.
4. **Task 4: Status Exposure**: Expose rate limit metrics in `handleStatus` and `handleAPIStatus`.

---

## Phase 5 — Requirement Mapping
- Req 1 & 3 -> Task 1 & Task 2
- Req 2 -> Task 3
- Req 4 -> Task 4

---

## Phase 6 — Reuse & Duplication Check
- 100% reuse of existing HTTP client methods, server handlers, and CLI commands.

---

## Phase 7 — Validation Plan
- Test rate limiter safety delay during API calls.
- Test sequential bulk deletion throttling.
- Test rate limit status reporting in CLI and REST API.
