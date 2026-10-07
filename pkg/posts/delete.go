package posts

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-insta-cli/pkg/cli"
)

// DeletePostParams holds parameters for deleting an Instagram post.
type DeletePostParams struct {
	PostID string
	Force  bool
}

// Validate checks if the required post ID argument is present.
func (p DeletePostParams) Validate() error {
	if strings.TrimSpace(p.PostID) == "" {
		return fmt.Errorf("missing required Post ID argument. Usage: insta delete <post_id> [--force|-f]")
	}
	return nil
}

// ConfirmDeletion displays an interactive confirmation prompt asking [y/N].
func ConfirmDeletion(postID string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%sAre you sure you want to delete post '%s'? [y/N]: %s", cli.ColorYellow, postID, cli.ColorReset)

	input, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	answer := strings.ToLower(strings.TrimSpace(input))
	return answer == "y" || answer == "yes"
}

// RenderDeleteSuccess displays colorized terminal notification for successful post deletion.
func RenderDeleteSuccess(postID string) {
	fmt.Printf("%s✓ Post '%s' deleted successfully.%s\n", cli.ColorGreen, postID, cli.ColorReset)
}

// RenderDeleteCancelled displays colorized terminal notification when user cancels deletion.
func RenderDeleteCancelled(postID string) {
	fmt.Printf("%s⚠ Deletion cancelled by user.%s\n", cli.ColorYellow, cli.ColorReset)
}

// RenderDeleteError displays colorized terminal notification when deletion fails.
func RenderDeleteError(postID string, err error) {
	fmt.Printf("%s✗ Error: Post '%s' not found or deletion failed: %v%s\n", cli.ColorRed, postID, err, cli.ColorReset)
}
