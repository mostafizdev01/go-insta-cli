# Step 5 Prompt: Post Deletion Engine & Safety Mechanism

Implement Step 5 for `go-insta-cli`:
1. Build `pkg/posts/delete.go` with interactive safety prompts (`[y/N]`) and `--force` flag.
2. Extend `pkg/instagram/client.go` with `DeletePost(mediaID string) error`.
3. Wire `handleDelete(args, cfg)` into `main.go` under `case "delete"`.
4. Ensure files stay strictly under 300 lines of clean Go code.
