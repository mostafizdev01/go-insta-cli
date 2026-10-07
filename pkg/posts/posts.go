package posts

import (
	"encoding/json"
	"fmt"

	"go-insta-cli/pkg/cli"
)

// Post represents an Instagram post data model.
type Post struct {
	ID           string `json:"id"`
	Caption      string `json:"caption"`
	Timestamp    string `json:"timestamp"`
	LikeCount    int    `json:"like_count"`
	CommentCount int    `json:"comment_count"`
	MediaType    string `json:"media_type"`
	MediaURL     string `json:"media_url"`
}

// RenderPostTable displays formatted ANSI colorized list of posts in the terminal.
func RenderPostTable(posts []Post) {
	if len(posts) == 0 {
		fmt.Printf("%sNo Instagram posts found.%s\n", cli.ColorYellow, cli.ColorReset)
		return
	}

	fmt.Printf("%s--- Instagram Posts (%d) ---%s\n\n", cli.ColorCyan, len(posts), cli.ColorReset)
	for i, post := range posts {
		fmt.Printf("%s[%d] Post ID: %s%s (%s)\n", cli.ColorGreen, i+1, post.ID, cli.ColorReset, post.Timestamp)
		fmt.Printf("    Caption:  %s\n", post.Caption)
		fmt.Printf("    Type:     %s\n", post.MediaType)
		fmt.Printf("    Metrics:  %s%d Likes%s | %s%d Comments%s\n", cli.ColorYellow, post.LikeCount, cli.ColorReset, cli.ColorCyan, post.CommentCount, cli.ColorReset)
		fmt.Println("    --------------------------------------------------")
	}
}

// RenderPostJSON outputs raw JSON formatted string for scripting.
func RenderPostJSON(posts []Post) (string, error) {
	data, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
