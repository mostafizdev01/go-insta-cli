# Step 7: Requirements & Acceptance Criteria

## Target Scope
Implement the Interactive Web Dashboard UI for `go-insta-cli`, enabling authenticated users accessing `http://localhost:8080` via browser to view Instagram posts as cards and perform single or bulk post deletion operations.

---

## Detailed Requirements

1. **Web Dashboard UI (`pkg/server/ui.go`)**:
   - Serve embedded single-page Web UI at `/` and `/index.html`.
   - Implement responsive CSS grid layout (`repeat(auto-fill, minmax(280px, 1fr))`).
   - Use clean, minimal dark theme styling (`#0f172a`, `#1e293b`, `#3897f0`, `#ef4444`).

2. **Card Visualization**:
   - Render post cards displaying ID, media badge (IMAGE/VIDEO), image preview, like count, comment count, caption, timestamp, and action buttons.

3. **Interactive Post Management**:
   - Checkboxes for selecting individual post cards.
   - Bulk toolbar with "Select All", "Deselect All", and "Delete Selected" buttons.
   - Single post "Delete" button per card.
   - Toast notifications for real-time operation feedback.

---

## Verification Criteria
- Navigating to `http://localhost:8080` loads the web dashboard HTML interface.
- Client-side JS fetches `/api/status` and displays user profile info (`Stavo Nelson`).
- Client-side JS fetches `/api/posts` and populates the responsive card grid.
- Selecting post cards activates the bulk deletion toolbar.
- Deleting posts triggers calls to `/api/delete` and updates the UI live with toast alerts.
