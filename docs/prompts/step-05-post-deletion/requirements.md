# Step 5: Requirements & Acceptance Criteria

## Target Scope
Implement the Post Deletion Engine & Safety Mechanism for `go-insta-cli`, allowing authenticated users to delete a post by Post ID with interactive safety confirmation prompts and `--force` flag support.

---

## Requested Requirements
1. **Post Deletion Engine (`pkg/posts/delete.go`)**:
   - Create post deletion handler with validation for `post_id`.
   - Implement interactive CLI prompt: `Are you sure you want to delete post <id>? [y/N]: `.
2. **Safety & Override Flags**:
   - Support `--force` (or `-f`) flag to skip interactive confirmation prompt.
3. **Instagram Network Integration (`pkg/instagram/client.go`)**:
   - Add `DeletePost(mediaID string) error` API method.
   - Include mock fallback handling for offline/test environments.
4. **Security & Authentication Guard**:
   - Enforce `auth.RequireAuth(cfg)` prior to deletion attempt.
5. **Terminal Feedback**:
   - Display colorized terminal notifications for confirmation, cancellation, and deletion results.

---

## Verification Criteria
- `insta delete` without authentication is blocked by `RequireAuth()`.
- `insta delete` without post ID displays error and usage guide.
- `insta delete 32849102839` prompts for `[y/N]`; entering `n` cancels deletion.
- `insta delete 32849102839 --force` skips prompt and completes deletion.
