package server

import "net/http"

// handleIndex serves the embedded single-page Web Dashboard UI.
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/index.html" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(indexHTMLHeader + indexJSFooter))
}

const indexHTMLHeader = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Instagram Management Dashboard</title>
    <style>
        :root {
            --bg-color: #0f172a;
            --card-bg: #1e293b;
            --card-border: #334155;
            --text-primary: #f8fafc;
            --text-secondary: #94a3b8;
            --accent-blue: #3897f0;
            --accent-red: #ef4444;
            --accent-green: #10b981;
            --hover-bg: #334155;
        }

        * { box-sizing: border-box; margin: 0; padding: 0; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
            background-color: var(--bg-color);
            color: var(--text-primary);
            min-height: 100vh;
            display: flex;
            flex-direction: column;
        }

        .navbar {
            background-color: var(--card-bg);
            border-bottom: 1px solid var(--card-border);
            padding: 1rem 2rem;
            display: flex;
            justify-content: space-between;
            align-items: center;
            position: sticky;
            top: 0;
            z-index: 100;
        }
        .logo { font-size: 1.25rem; font-weight: 700; color: var(--text-primary); display: flex; align-items: center; gap: 0.5rem; }
        .logo span { color: var(--accent-blue); }
        .status-badge {
            background: rgba(16, 185, 129, 0.15);
            color: var(--accent-green);
            padding: 0.35rem 0.75rem;
            border-radius: 9999px;
            font-size: 0.85rem;
            font-weight: 600;
            display: flex;
            align-items: center;
            gap: 0.4rem;
        }

        .container { max-width: 1200px; margin: 0 auto; width: 100%; padding: 2rem 1.5rem; flex: 1; }

        .toolbar {
            background-color: var(--card-bg);
            border: 1px solid var(--card-border);
            border-radius: 12px;
            padding: 1rem 1.5rem;
            margin-bottom: 2rem;
            display: flex;
            flex-wrap: wrap;
            justify-content: space-between;
            align-items: center;
            gap: 1rem;
        }
        .stats { display: flex; gap: 1.5rem; color: var(--text-secondary); font-size: 0.95rem; }
        .stats strong { color: var(--text-primary); }
        .actions { display: flex; gap: 0.75rem; align-items: center; }

        button {
            background-color: var(--card-border);
            color: var(--text-primary);
            border: none;
            padding: 0.5rem 1rem;
            border-radius: 8px;
            font-size: 0.9rem;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.2s ease;
        }
        button:hover { background-color: var(--hover-bg); opacity: 0.9; }
        button.btn-danger { background-color: var(--accent-red); color: white; }
        button.btn-danger:disabled { opacity: 0.5; cursor: not-allowed; }
        button.btn-primary { background-color: var(--accent-blue); color: white; }

        .post-grid {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
            gap: 1.5rem;
        }

        .card {
            background-color: var(--card-bg);
            border: 1px solid var(--card-border);
            border-radius: 12px;
            overflow: hidden;
            display: flex;
            flex-direction: column;
            position: relative;
            transition: transform 0.2s ease, border-color 0.2s ease;
        }
        .card.selected { border-color: var(--accent-blue); transform: translateY(-2px); }
        .card-header {
            padding: 0.75rem 1rem;
            display: flex;
            justify-content: space-between;
            align-items: center;
            border-bottom: 1px solid var(--card-border);
        }
        .checkbox-wrapper { display: flex; align-items: center; gap: 0.5rem; cursor: pointer; font-size: 0.85rem; color: var(--text-secondary); }
        .checkbox-wrapper input { width: 18px; height: 18px; cursor: pointer; accent-color: var(--accent-blue); }

        .image-container {
            width: 100%;
            height: 240px;
            background-color: #090d16;
            display: flex;
            align-items: center;
            justify-content: center;
            overflow: hidden;
            position: relative;
        }
        .image-container img { width: 100%; height: 100%; object-fit: cover; }
        .media-badge {
            position: absolute;
            top: 10px;
            right: 10px;
            background: rgba(0, 0, 0, 0.6);
            color: white;
            font-size: 0.75rem;
            padding: 2px 8px;
            border-radius: 4px;
        }

        .card-body { padding: 1rem; flex: 1; display: flex; flex-direction: column; gap: 0.75rem; }
        .metrics { display: flex; gap: 1rem; font-size: 0.9rem; font-weight: 600; color: var(--text-primary); }
        .metric-item { display: flex; align-items: center; gap: 0.3rem; }
        .caption { font-size: 0.9rem; color: var(--text-secondary); line-height: 1.4; flex: 1; word-break: break-word; }
        .timestamp { font-size: 0.75rem; color: #64748b; margin-top: auto; }

        .card-footer {
            padding: 0.75rem 1rem;
            border-top: 1px solid var(--card-border);
            display: flex;
            justify-content: flex-end;
        }

        .toast-container { position: fixed; bottom: 20px; right: 20px; z-index: 1000; display: flex; flex-direction: column; gap: 10px; }
        .toast { background-color: var(--card-bg); border: 1px solid var(--card-border); color: white; padding: 12px 20px; border-radius: 8px; font-size: 0.9rem; box-shadow: 0 10px 15px -3px rgba(0,0,0,0.5); }
        .toast.success { border-color: var(--accent-green); }
        .toast.error { border-color: var(--accent-red); }

        .empty-state { text-align: center; padding: 4rem 1rem; color: var(--text-secondary); grid-column: 1 / -1; }
        .empty-state h3 { font-size: 1.5rem; color: var(--text-primary); margin-bottom: 0.5rem; }
    </style>
</head>
<body>

    <nav class="navbar">
        <div class="logo">
            <span>insta</span>-cli Dashboard
        </div>
        <div class="status-badge" id="statusBadge">
            ● Connected (<span id="userAccount">Stavo Nelson</span>)
        </div>
    </nav>

    <div class="container">
        <div class="toolbar">
            <div class="stats">
                <div>Total Posts: <strong id="totalPostsCount">0</strong></div>
                <div>Selected: <strong id="selectedCount">0</strong></div>
            </div>
            <div class="actions">
                <button onclick="selectAll(true)">Select All</button>
                <button onclick="selectAll(false)">Deselect</button>
                <button class="btn-primary" onclick="loadPosts(true)">Refresh</button>
                <button class="btn-danger" id="deleteSelectedBtn" onclick="deleteSelected()" disabled>Delete Selected</button>
            </div>
        </div>

        <div class="post-grid" id="postGrid">
        </div>
    </div>

    <div class="toast-container" id="toastContainer"></div>
`
