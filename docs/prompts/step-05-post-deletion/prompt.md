# Step 5 Prompt: Post Deletion Engine & Safety Mechanism

## System Role & Objective
Act as a Principal Software Architect and QA Engineer. Implement and verify Step 5 (**Post Deletion Engine & Safety Mechanism**) for `go-insta-cli` on top of Meta Graph API v19.0.

## Implementation Objectives
1. Create `pkg/posts/delete.go` containing deletion parameter structs, validation, interactive confirmation prompts (`[y/N]`), and colorized terminal feedback.
2. Create `pkg/instagram/delete.go` implementing official Meta Graph API deletion (`DELETE https://graph.facebook.com/v19.0/{post_id}` or `graph.instagram.com/v19.0/{post_id}`).
3. Wire `handleDelete(args, cfg)` in `main.go` under subcommand `case "delete"`.
4. Support `--force` / `-f` CLI flags to bypass interactive confirmation for automation and CI/CD pipelines.
5. Protect post deletion using `auth.RequireAuth(cfg)` security guard.
6. Enforce zero mock fallbacks, zero browser cookie dependencies, and strict file length constraints (<300 lines per file).
