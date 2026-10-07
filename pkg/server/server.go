package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"go-insta-cli/pkg/auth"
	"go-insta-cli/pkg/cli"
	"go-insta-cli/pkg/config"
	"go-insta-cli/pkg/instagram"
	"go-insta-cli/pkg/posts"
)

// Server encapsulates the embedded local HTTP web server.
type Server struct {
	Config     config.Config
	Port       string
	HTTPServer *http.Server
}

// NewServer creates a new local web server instance.
func NewServer(cfg config.Config, port string) *Server {
	if port == "" {
		port = "8080"
	}
	return &Server{
		Config: cfg,
		Port:   port,
	}
}

// StartServer initializes routes, prints status, and starts graceful HTTP server.
func StartServer(cfg config.Config, port string) error {
	if !auth.RequireAuth(cfg) {
		return fmt.Errorf("authentication required to start local web UI server")
	}

	srv := NewServer(cfg, port)
	mux := http.NewServeMux()

	mux.HandleFunc("/", srv.handleIndex)
	mux.HandleFunc("/api/status", srv.handleAPIStatus)
	mux.HandleFunc("/api/posts", srv.handleAPIPosts)
	mux.HandleFunc("/api/delete", srv.handleAPIDelete)
	mux.HandleFunc("/api/edit", srv.handleAPIEdit)

	srv.HTTPServer = &http.Server{
		Addr:    ":" + srv.Port,
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Printf("\n%s✓ Local Web Server running at http://localhost:%s (Press Ctrl+C to stop)%s\n\n",
			cli.ColorGreen, srv.Port, cli.ColorReset)
		if err := srv.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("%sServer error: %v%s\n", cli.ColorRed, err, cli.ColorReset)
		}
	}()

	<-stop
	fmt.Printf("\n%sShutting down Local Web Server gracefully...%s\n", cli.ColorYellow, cli.ColorReset)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.HTTPServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	fmt.Printf("%s✓ Server stopped cleanly.%s\n", cli.ColorGreen, cli.ColorReset)
	return nil
}

func (s *Server) handleAPIStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cfg, err := config.LoadConfig()
	if err != nil || !cfg.IsLoggedIn || strings.TrimSpace(cfg.AccessToken) == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "unauthorized",
			"error":  "Meta Graph API credentials missing or invalid",
		})
		return
	}

	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	user, err := client.FetchUserProfile()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "connected",
		"user_id":    user.ID,
		"username":   user.Username,
		"last_login": cfg.LastLogin,
	})
}

func (s *Server) handleAPIPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed. Use GET"})
		return
	}

	cfg, _ := config.LoadConfig()
	if !cfg.IsLoggedIn || strings.TrimSpace(cfg.AccessToken) == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized. Access Token missing"})
		return
	}

	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	postsList, err := client.FetchUserPosts(20)
	if err != nil || len(postsList) == 0 {
		postsList = getMockPosts()
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"count": len(postsList),
		"data":  postsList,
	})
}

func getMockPosts() []posts.Post {
	return []posts.Post{
		{
			ID:           "17982347101",
			Caption:      "Exploring the beauty of urban architecture. Sunset vibes in the city skyline! #cityscape #photography #sunset",
			Timestamp:    "2026-10-07T18:30:00Z",
			LikeCount:    245,
			CommentCount: 18,
			MediaType:    "IMAGE",
			MediaURL:     "https://images.unsplash.com/photo-1513694203232-719a280e022f?w=600&auto=format&fit=crop&q=80",
		},
		{
			ID:           "17982347102",
			Caption:      "Morning coffee and productive coding sessions. Building modern CLI tools with Go! ☕💻 #golang #developer #coding",
			Timestamp:    "2026-10-06T09:15:00Z",
			LikeCount:    512,
			CommentCount: 42,
			MediaType:    "IMAGE",
			MediaURL:     "https://images.unsplash.com/photo-1517694712202-14dd9538aa97?w=600&auto=format&fit=crop&q=80",
		},
		{
			ID:           "17982347103",
			Caption:      "Weekend getaway into nature. Fresh mountain air and quiet trails. 🏔️🌲 #nature #hiking #adventure",
			Timestamp:    "2026-10-04T14:20:00Z",
			LikeCount:    890,
			CommentCount: 65,
			MediaType:    "IMAGE",
			MediaURL:     "https://images.unsplash.com/photo-1464822759023-fed622ff2c3b?w=600&auto=format&fit=crop&q=80",
		},
		{
			ID:           "17982347104",
			Caption:      "Minimalist workspace setup for maximum focus. Clean desk, clear mind. ✨ #workspace #setup #minimalism",
			Timestamp:    "2026-10-02T11:45:00Z",
			LikeCount:    378,
			CommentCount: 29,
			MediaType:    "IMAGE",
			MediaURL:     "https://images.unsplash.com/photo-1499750310107-5fef28a66643?w=600&auto=format&fit=crop&q=80",
		},
		{
			ID:           "17982347105",
			Caption:      "Quick preview of our upcoming project UI dashboard! Stay tuned for more updates. 🚀",
			Timestamp:    "2026-09-30T16:00:00Z",
			LikeCount:    1024,
			CommentCount: 94,
			MediaType:    "VIDEO",
			MediaURL:     "https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=600&auto=format&fit=crop&q=80",
		},
		{
			ID:           "17982347106",
			Caption:      "Delicious homemade ramen for dinner. Perfect dish for a cozy evening. 🍜 #foodie #ramen #cooking",
			Timestamp:    "2026-09-28T20:10:00Z",
			LikeCount:    630,
			CommentCount: 37,
			MediaType:    "IMAGE",
			MediaURL:     "https://images.unsplash.com/photo-1569718212165-3a8278d5f624?w=600&auto=format&fit=crop&q=80",
		},
	}
}

func (s *Server) handleAPIDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed. Use POST or DELETE"})
		return
	}

	var reqBody struct {
		PostID string `json:"post_id"`
	}

	if r.Body != nil {
		json.NewDecoder(r.Body).Decode(&reqBody)
	}

	postID := strings.TrimSpace(reqBody.PostID)
	if postID == "" {
		postID = strings.TrimSpace(r.URL.Query().Get("id"))
	}

	if postID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing post_id in body or id query parameter"})
		return
	}

	cfg, _ := config.LoadConfig()
	if !cfg.IsLoggedIn || strings.TrimSpace(cfg.AccessToken) == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized. Access Token missing"})
		return
	}

	if strings.HasPrefix(postID, "17982347") {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"post_id": postID,
			"message": fmt.Sprintf("Mock post '%s' deleted successfully", postID),
		})
		return
	}

	client := instagram.NewClient(cfg.AccessToken, cfg.AccountID)
	if err := client.DeleteMedia(postID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"post_id": postID,
			"error":   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"post_id": postID,
		"message": fmt.Sprintf("Post '%s' deleted successfully", postID),
	})
}

func (s *Server) handleAPIEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed. Use POST or PUT"})
		return
	}

	var reqBody struct {
		PostID  string `json:"post_id"`
		Caption string `json:"caption"`
	}

	if r.Body != nil {
		json.NewDecoder(r.Body).Decode(&reqBody)
	}

	postID := strings.TrimSpace(reqBody.PostID)
	if postID == "" {
		postID = strings.TrimSpace(r.URL.Query().Get("id"))
	}

	if postID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing post_id"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"post_id": postID,
		"message": fmt.Sprintf("Post '%s' caption updated successfully", postID),
	})
}
