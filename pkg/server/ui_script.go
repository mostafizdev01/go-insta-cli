package server

const indexJSFooter = `    <script>
        var allPosts = [];
        var selectedIDs = new Set();

        document.addEventListener("DOMContentLoaded", function() {
            loadStatus();
            loadPosts(true);
            setInterval(function() {
                loadPosts(false);
            }, 10000);
        });

        function loadStatus() {
            fetch('/api/status')
                .then(function(res) { return res.json(); })
                .then(function(data) {
                    if (data.status === "connected") {
                        document.getElementById('userAccount').innerText = data.username || data.user_id;
                    }
                })
                .catch(function(e) { console.error("Status fetch failed", e); });
        }

        function loadPosts(showSpinner) {
            var grid = document.getElementById('postGrid');
            if (showSpinner !== false) {
                grid.innerHTML = '<div class="empty-state">Loading posts from Meta Graph API...</div>';
                selectedIDs.clear();
                updateSelectionUI();
            }

            fetch('/api/posts')
                .then(function(res) { return res.json(); })
                .then(function(data) {
                    if (data.error) {
                        if (showSpinner !== false) {
                            grid.innerHTML = '<div class="empty-state"><h3>Meta Graph API Status</h3><p>' + escapeHtml(data.error) + '</p></div>';
                            document.getElementById('totalPostsCount').innerText = "0";
                        }
                        return;
                    }

                    allPosts = data.data || [];
                    document.getElementById('totalPostsCount').innerText = allPosts.length;

                    if (allPosts.length === 0) {
                        grid.innerHTML = '<div class="empty-state"><h3>No Posts Found</h3><p>Your Instagram profile currently has 0 posts.</p></div>';
                        return;
                    }

                    renderGrid();
                })
                .catch(function(e) {
                    if (showSpinner !== false) {
                        grid.innerHTML = '<div class="empty-state"><h3>Connection Error</h3><p>Failed to load posts from API.</p></div>';
                    }
                });
        }

        function renderGrid() {
            var grid = document.getElementById('postGrid');
            grid.innerHTML = '';

            allPosts.forEach(function(post) {
                var isSelected = selectedIDs.has(post.id);
                var card = document.createElement('div');
                card.className = 'card' + (isSelected ? ' selected' : '');
                card.id = 'card-' + post.id;

                var imgURL = post.media_url || 'https://via.placeholder.com/300x300/1e293b/94a3b8?text=Instagram+Media';
                var mediaType = post.media_type || 'IMAGE';
                var captionText = post.caption || 'No caption provided.';
                var likes = post.like_count || 0;
                var comments = post.comment_count || 0;
                var timeStr = post.timestamp || '';

                var html = '<div class="card-header">';
                html += '<label class="checkbox-wrapper">';
                html += '<input type="checkbox" ' + (isSelected ? 'checked' : '') + ' onchange="toggleSelect(\'' + post.id + '\', this.checked)">';
                html += 'ID: ' + post.id.substring(0, 10) + '...';
                html += '</label></div>';
                html += '<div class="image-container">';
                html += '<img src="' + escapeHtml(imgURL) + '" alt="Post image" onerror="this.src=\'https://via.placeholder.com/300x300/1e293b/94a3b8?text=Media+Preview\'">';
                html += '<div class="media-badge">' + mediaType + '</div></div>';
                html += '<div class="card-body">';
                html += '<div class="metrics">';
                html += '<div class="metric-item">♥ ' + likes + '</div>';
                html += '<div class="metric-item">💬 ' + comments + '</div></div>';
                html += '<div class="caption">' + escapeHtml(captionText) + '</div>';
                html += '<div class="timestamp">' + timeStr + '</div></div>';
                html += '<div class="card-footer" style="display:flex; gap:0.5rem; justify-content:flex-end;">';
                html += '<button class="btn-primary" onclick="editSingle(\'' + post.id + '\')">Edit</button>';
                html += '<button class="btn-danger" onclick="deleteSingle(\'' + post.id + '\')">Delete</button></div>';

                card.innerHTML = html;
                grid.appendChild(card);
            });
        }

        function editSingle(id) {
            var post = allPosts.find(function(p) { return p.id === id; });
            var currentCaption = post ? post.caption : '';
            var newCaption = prompt('Edit caption for post ID ' + id + ':', currentCaption);
            if (newCaption === null || newCaption.trim() === '') return;

            fetch('/api/edit?id=' + encodeURIComponent(id), {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ post_id: id, caption: newCaption })
            })
            .then(function(res) { return res.json(); })
            .then(function(data) {
                if (data.success) {
                    showToast('Post caption updated successfully', 'success');
                    if (post) post.caption = newCaption;
                    renderGrid();
                } else {
                    showToast('Failed to edit post: ' + (data.error || 'API Error'), 'error');
                }
            })
            .catch(function(e) {
                showToast('Network error editing post ' + id, 'error');
            });
        }

        function toggleSelect(id, checked) {
            if (checked) {
                selectedIDs.add(id);
            } else {
                selectedIDs.delete(id);
            }
            var card = document.getElementById('card-' + id);
            if (card) {
                card.classList.toggle('selected', checked);
            }
            updateSelectionUI();
        }

        function selectAll(check) {
            selectedIDs.clear();
            if (check) {
                allPosts.forEach(function(p) { selectedIDs.add(p.id); });
            }
            updateSelectionUI();
            renderGrid();
        }

        function updateSelectionUI() {
            document.getElementById('selectedCount').innerText = selectedIDs.size;
            document.getElementById('deleteSelectedBtn').disabled = selectedIDs.size === 0;
        }

        function deleteSingle(id) {
            if (!confirm('Are you sure you want to delete post ID ' + id + '?')) return;
            executeDelete([id]);
        }

        function deleteSelected() {
            var count = selectedIDs.size;
            if (count === 0) return;
            if (!confirm('Are you sure you want to delete ' + count + ' selected post(s)?')) return;
            executeDelete(Array.from(selectedIDs));
        }

        function executeDelete(ids) {
            var index = 0;
            function deleteNext() {
                if (index >= ids.length) {
                    loadPosts(true);
                    return;
                }
                var id = ids[index];
                showToast('Deleting post (' + (index + 1) + '/' + ids.length + ') - 3s safety pause...', 'info');
                fetch('/api/delete?id=' + encodeURIComponent(id), { method: 'POST' })
                    .then(function(res) { return res.json(); })
                    .then(function(data) {
                        if (data.success) {
                            showToast('Post ' + id + ' deleted successfully', 'success');
                        } else {
                            showToast('Failed to delete post ' + id + ': ' + (data.error || 'API Error'), 'error');
                        }
                    })
                    .catch(function(e) {
                        showToast('Network error deleting post ' + id, 'error');
                    })
                    .finally(function() {
                        index++;
                        if (index < ids.length) {
                            setTimeout(deleteNext, 3000);
                        } else {
                            setTimeout(function() { loadPosts(true); }, 1000);
                        }
                    });
            }
            deleteNext();
        }

        function showToast(msg, type) {
            var container = document.getElementById('toastContainer');
            var toast = document.createElement('div');
            toast.className = 'toast ' + type;
            toast.innerText = msg;
            container.appendChild(toast);
            setTimeout(function() { toast.remove(); }, 4000);
        }

        function escapeHtml(str) {
            return String(str).replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');
        }
    </script>
</body>
</html>`
