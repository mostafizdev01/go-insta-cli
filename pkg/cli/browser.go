package cli

import (
	"fmt"
	"os/exec"
	"runtime"
)

// OpenBrowser launches the OS default web browser to the specified URL.
// Supports Windows, macOS, and Linux platforms.
func OpenBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux", "freebsd":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return cmd.Start()
}
