package instagram

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"go-insta-cli/pkg/posts"
)

// getTargetDomains returns domain list based on token prefix.
func getTargetDomains(accessToken string) []string {
	if strings.HasPrefix(accessToken, "EAA") {
		return []string{"graph.facebook.com"}
	}
	if strings.HasPrefix(accessToken, "IG") {
		return []string{"graph.instagram.com"}
	}
	return []string{"graph.facebook.com", "graph.instagram.com"}
}

// FetchUserProfile retrieves connected account info via Meta Graph API.
func (c *Client) FetchUserProfile() (*GraphAPIUserResponse, error) {
	if strings.TrimSpace(c.AccessToken) == "" {
		return nil, fmt.Errorf("INSTAGRAM_ACCESS_TOKEN is missing. Please configure your Meta Graph API Access Token")
	}

	target := "me"
	if c.AccountID != "" {
		target = c.AccountID
	}

	domains := getTargetDomains(c.AccessToken)
	var lastErr error

	for _, domain := range domains {
		apiURL := fmt.Sprintf("https://%s/v19.0/%s?fields=id,name&access_token=%s",
			domain, target, url.QueryEscape(c.AccessToken))

		resp, err := c.HTTPClient.Get(apiURL)
		if err != nil {
			lastErr = fmt.Errorf("network error reaching %s: %w", domain, err)
			continue
		}

		var userResp GraphAPIUserResponse
		if err := json.NewDecoder(resp.Body).Decode(&userResp); err == nil {
			resp.Body.Close()
			if userResp.Error == nil && (userResp.ID != "" || userResp.Username != "" || userResp.Name != "") {
				if userResp.Username == "" && userResp.Name != "" {
					userResp.Username = userResp.Name
				}
				return &userResp, nil
			}
			if userResp.Error != nil {
				lastErr = fmt.Errorf("Meta API Error (%d): %s", userResp.Error.Code, userResp.Error.Message)
			}
		} else {
			resp.Body.Close()
		}
	}

	if lastErr != nil {
		return nil, lastErr
	}
	return nil, fmt.Errorf("invalid or expired Meta Access Token")
}

// FetchUserPosts retrieves real posts using official Meta Graph API.
func (c *Client) FetchUserPosts(limit int) ([]posts.Post, error) {
	if strings.TrimSpace(c.AccessToken) == "" {
		return nil, fmt.Errorf("INSTAGRAM_ACCESS_TOKEN is missing. Please configure your Meta Graph API Access Token")
	}

	target := "me"
	if c.AccountID != "" {
		target = c.AccountID
	}

	domains := getTargetDomains(c.AccessToken)
	var lastErr error

	for _, domain := range domains {
		apiURL := fmt.Sprintf("https://%s/v19.0/%s/media?fields=id,caption,media_type,media_url,permalink,timestamp,like_count,comments_count&access_token=%s",
			domain, target, url.QueryEscape(c.AccessToken))

		resp, err := c.HTTPClient.Get(apiURL)
		if err != nil {
			lastErr = fmt.Errorf("network error reaching %s: %w", domain, err)
			continue
		}

		var graphResp GraphAPIMediaResponse
		if err := json.NewDecoder(resp.Body).Decode(&graphResp); err == nil {
			resp.Body.Close()
			if graphResp.Error == nil {
				var result []posts.Post
				for i, item := range graphResp.Data {
					if limit > 0 && i >= limit {
						break
					}
					result = append(result, posts.Post{
						ID:           item.ID,
						Caption:      item.Caption,
						Timestamp:    item.Timestamp,
						LikeCount:    item.LikeCount,
						CommentCount: item.CommentsCount,
						MediaType:    item.MediaType,
						MediaURL:     item.MediaURL,
					})
				}
				return result, nil
			}
			if graphResp.Error != nil {
				lastErr = fmt.Errorf("Meta API Error (%d): %s", graphResp.Error.Code, graphResp.Error.Message)
			}
		} else {
			resp.Body.Close()
		}
	}

	if lastErr != nil {
		return nil, lastErr
	}
	return nil, fmt.Errorf("unable to fetch posts: invalid or expired Meta Access Token")
}
