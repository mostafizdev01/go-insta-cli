# Step 4: Requirements & Acceptance Criteria

## Target Scope
Implement the Post Retrieval Engine for `go-insta-cli`, enabling terminal users to fetch and display recent Instagram posts with support for flags (`--limit`, `--json`) and colorized table rendering.

---

## Requested Requirements
1. **Post Data Structure**:
   - Define `Post` struct containing `ID`, `Caption`, `Timestamp`, `LikeCount`, `CommentCount`, `MediaType`, `MediaURL`.
2. **Post Retrieval Engine (`pkg/posts/posts.go`)**:
   - Implement `FetchPosts(limit int) ([]Post, error)` with Instagram API connectivity and mock data fallback for safe offline testing.
3. **CLI Flags Support**:
   - `--limit <n>`: Restrict maximum number of posts fetched (default: 10).
   - `--json`: Output raw formatted JSON array for scripting.
4. **Terminal Output Formatting**:
   - Render structured ANSI colorized terminal list/table with post metrics.
5. **Security & Error Handling**:
   - Enforce `auth.RequireAuth(cfg)` authentication guard from Step 3.
   - Gracefully handle empty post lists, API rate limits, and network errors.

---

## Verification Criteria
- `insta posts` when logged out is blocked by `RequireAuth()` guard.
- `insta posts` when logged in outputs formatted colorized post table.
- `insta posts --limit 2` returns exactly 2 posts.
- `insta posts --json` returns valid JSON string array.
