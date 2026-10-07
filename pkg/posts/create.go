package posts

import (
	"fmt"
	"strings"

	"go-insta-cli/pkg/cli"
)

// CreatePostParams defines the parameters for creating a new Instagram post.
type CreatePostParams struct {
	ImageURL string
	Caption  string
}

// Validate checks if the required post creation parameters are present.
func (p CreatePostParams) Validate() error {
	if strings.TrimSpace(p.ImageURL) == "" {
		return fmt.Errorf("image URL cannot be empty. Usage: insta create <image_url> [caption...]")
	}
	return nil
}

// RenderCreateSuccess displays a colorized terminal notification for successful post creation.
func RenderCreateSuccess(postID string) {
	fmt.Printf("\n%sSuccessfully published post to Instagram!%s\n", cli.ColorGreen, cli.ColorReset)
	if postID != "" {
		fmt.Printf("  Post ID: %s%s%s\n", cli.ColorYellow, postID, cli.ColorReset)
	}
	fmt.Printf("  Status:  %sLIVE on Instagram Profile%s\n\n", cli.ColorCyan, cli.ColorReset)
}
