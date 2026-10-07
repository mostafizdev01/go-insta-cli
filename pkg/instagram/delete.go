package instagram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// DeleteMedia removes an Instagram media post by Post ID using Meta Graph API.
func (c *Client) DeleteMedia(mediaID string) error {
	mediaID = strings.TrimSpace(mediaID)
	if mediaID == "" {
		return fmt.Errorf("media ID cannot be empty")
	}

	if strings.TrimSpace(c.AccessToken) == "" {
		return fmt.Errorf("INSTAGRAM_ACCESS_TOKEN is missing. Please configure your Meta Graph API Access Token")
	}

	domains := getTargetDomains(c.AccessToken)
	var lastErr error

	for _, domain := range domains {
		apiURL := fmt.Sprintf("https://%s/v19.0/%s?access_token=%s",
			domain, url.QueryEscape(mediaID), url.QueryEscape(c.AccessToken))

		req, err := http.NewRequest("DELETE", apiURL, nil)
		if err != nil {
			lastErr = fmt.Errorf("failed to construct DELETE request: %w", err)
			continue
		}

		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("network error reaching %s: %w", domain, err)
			continue
		}

		var delResp struct {
			Success bool `json:"success"`
			Error   *struct {
				Message string `json:"message"`
				Type    string `json:"type"`
				Code    int    `json:"code"`
			} `json:"error"`
		}

		err = json.NewDecoder(resp.Body).Decode(&delResp)
		resp.Body.Close()

		if err == nil {
			if delResp.Success {
				return nil
			}
			if delResp.Error != nil {
				lastErr = fmt.Errorf("Meta API Deletion Error (%d): %s", delResp.Error.Code, delResp.Error.Message)
			} else {
				lastErr = fmt.Errorf("deletion request failed for media ID '%s' (HTTP %d)", mediaID, resp.StatusCode)
			}
		} else {
			lastErr = fmt.Errorf("failed to parse deletion response (HTTP %d): %w", resp.StatusCode, err)
		}
	}

	if lastErr != nil {
		return lastErr
	}
	return fmt.Errorf("unable to delete post '%s': invalid Post ID or missing permissions", mediaID)
}
