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

	mux.HandleFunc("/api/status", srv.handleAPIStatus)
	mux.HandleFunc("/api/posts", srv.handleAPIPosts)
	mux.HandleFunc("/api/delete", srv.handleAPIDelete)

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
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"count": len(postsList),
		"data":  postsList,
	})
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
