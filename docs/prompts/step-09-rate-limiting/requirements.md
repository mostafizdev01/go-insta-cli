# Step 9: Requirements & Acceptance Criteria

## Objective
Implement API Rate Limiting, Request Throttling, Exponential Backoff, and Circuit Breaker protection controls in `go-insta-cli` to keep Instagram accounts safe from Meta anti-bot WAF blocks during API calls and bulk operations.

---

## Detailed Requirements

1. **Request Throttling & Rate Limiter Engine (`pkg/instagram/ratelimit.go`)**:
   - Central rate limiter middleware intercepting all Meta Graph API requests (Post Retrieval, Post Deletion, Authentication).
   - Mandatory safety delay (1-2 seconds) between API requests.

2. **Bulk Action Throttling**:
   - Execute bulk post deletions (via CLI or Web UI) sequentially with a forced safety interval (3 seconds per deletion) instead of parallel spamming.

3. **Exponential Backoff & Circuit Breaker**:
   - Automatic exponential backoff retries when HTTP 429 (Too Many Requests) or rate limit error occurs.
   - Circuit Breaker: Halt execution immediately if rate limit errors persist, display a red/yellow terminal warning (`⚠ Rate limit detected. Pausing requests for account safety.`), and return JSON error feedback to Web UI.

4. **Rate Limit Status Exposure**:
   - Include rate limit status (quota/throttle status) in `insta status` CLI output and `GET /api/status` REST API response.

---

## Acceptance Criteria
- Bulk deletions execute sequentially with 3-second safety delays.
- Rate limiting middleware intercepts Meta API HTTP requests without breaking existing authentication or retrieval handlers.
- Circuit breaker halts requests upon repeated HTTP 429 responses and displays warnings in terminal and Web UI.
- `insta status` and `GET /api/status` include active rate limit status information.
- All Go files remain strictly under 300 lines of code.
