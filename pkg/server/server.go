package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-insta-cli/pkg/auth"
	"go-insta-cli/pkg/cli"
	"go-insta-cli/pkg/config"
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
	mux.HandleFunc("/api/create", srv.handleAPICreate)

	srv.HTTPServer = &http.Server{
		Addr:    ":" + srv.Port,
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	targetURL := fmt.Sprintf("http://localhost:%s/", srv.Port)
	go func() {
		fmt.Printf("\n%s✓ Local Web Server running at %s (Press Ctrl+C to stop)%s\n",
			cli.ColorGreen, targetURL, cli.ColorReset)

		go func() {
			time.Sleep(500 * time.Millisecond)
			if err := cli.OpenBrowser(targetURL); err != nil {
				fmt.Printf("%s🔗 Open Dashboard: %s%s\n\n", cli.ColorCyan, targetURL, cli.ColorReset)
			}
		}()

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
