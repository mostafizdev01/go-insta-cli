# Step 5: Requirements & Acceptance Criteria

## Target Scope
Implement the Post Deletion Engine & Safety Mechanism for `go-insta-cli`, allowing authenticated users to delete an Instagram post by Post ID via Meta Graph API, with interactive confirmation prompts (`[y/N]`) and `--force` flag support.

---

## Detailed Requirements

1. **Post Deletion Engine (`pkg/posts/delete.go`)**:
   - Create post deletion parameter validation for `<post_id>`.
   - Display interactive terminal prompt: `Are you sure you want to delete post <post_id>? [y/N]: `.
   - Cancel deletion if user enters anything other than `y` or `Y`.

2. **Safety Override Flags**:
   - Support `--force` and `-f` flags to skip the interactive confirmation prompt for automated scripting.

3. **Meta Graph API Integration (`pkg/instagram/delete.go`)**:
   - Add `DeleteMedia(mediaID string) error` API method making HTTP DELETE calls to Meta Graph API (`v19.0`).
   - Eliminate all mock fallbacks; return real API response or explicit error on failure.

4. **Security & Authentication Guard**:
   - Protect subcommand execution using `auth.RequireAuth(cfg)` guard.
   - Prohibit request execution if `INSTAGRAM_ACCESS_TOKEN` is missing or invalid.

5. **Terminal Feedback (ANSI Colors)**:
   - Green: `✓ Post <post_id> deleted successfully.`
   - Yellow: `⚠ Deletion cancelled by user.`
   - Red: `✗ Error: Post <post_id> not found or deletion failed.`

---

## Verification Criteria
- `insta delete` without authentication is blocked by `RequireAuth()`.
- `insta delete` without post ID displays error and usage guide.
- `insta delete <post_id>` prompts for `[y/N]`; entering `n` cancels deletion with yellow warning.
- `insta delete <post_id> --force` skips prompt and invokes API deletion.
- Invalid post ID or API rate limit errors output clear red error messages.
