package instagram

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// ValidateGraphToken checks if the Meta Graph API access token is valid.
func (c *Client) ValidateGraphToken() bool {
	if strings.TrimSpace(c.AccessToken) == "" {
		return false
	}
	user, err := c.FetchUserProfile()
	return err == nil && user != nil && user.ID != ""
}

// PublishMedia creates and publishes a post using Meta Content Publishing API.
func (c *Client) PublishMedia(imageURL, caption string) (string, error) {
	if strings.TrimSpace(c.AccessToken) == "" {
		return "", fmt.Errorf("INSTAGRAM_ACCESS_TOKEN is missing. Please configure your Meta Graph API Access Token")
	}

	if strings.TrimSpace(imageURL) == "" {
		return "", fmt.Errorf("image URL cannot be empty")
	}

	target := "me"
	if c.AccountID != "" {
		target = c.AccountID
	}

	domains := getTargetDomains(c.AccessToken)
	domain := domains[0]

	// 1. Create Media Container
	createURL := fmt.Sprintf("https://%s/v19.0/%s/media?image_url=%s&caption=%s&access_token=%s",
		domain, target, url.QueryEscape(imageURL), url.QueryEscape(caption), url.QueryEscape(c.AccessToken))

	resp, err := c.HTTPClient.Post(createURL, "application/json", nil)
	if err != nil {
		return "", fmt.Errorf("network error creating media container: %w", err)
	}
	defer resp.Body.Close()

	var containerResp struct {
		ID    string `json:"id"`
		Error *struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		} `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&containerResp); err != nil {
		return "", fmt.Errorf("failed to parse media container response: %w", err)
	}

	if containerResp.Error != nil {
		return "", fmt.Errorf("Meta API Creation Error (%d): %s", containerResp.Error.Code, containerResp.Error.Message)
	}

	// 2. Publish Media Container
	publishURL := fmt.Sprintf("https://%s/v19.0/%s/media_publish?creation_id=%s&access_token=%s",
		domain, target, containerResp.ID, url.QueryEscape(c.AccessToken))

	pubResp, err := c.HTTPClient.Post(publishURL, "application/json", nil)
	if err != nil {
		return "", fmt.Errorf("network error publishing media container: %w", err)
	}
	defer pubResp.Body.Close()

	var pubData struct {
		ID    string `json:"id"`
		Error *struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		} `json:"error"`
	}

	if err := json.NewDecoder(pubResp.Body).Decode(&pubData); err != nil {
		return "", fmt.Errorf("failed to parse publish response: %w", err)
	}

	if pubData.Error != nil {
		return "", fmt.Errorf("Meta API Publish Error (%d): %s", pubData.Error.Code, pubData.Error.Message)
	}

	return pubData.ID, nil
}

// UpdateMediaCaption sends POST request to Meta Graph API to update media caption.
func (c *Client) UpdateMediaCaption(mediaID, caption string) error {
	mediaID = strings.TrimSpace(mediaID)
	if mediaID == "" {
		return fmt.Errorf("media ID cannot be empty")
	}
	if strings.TrimSpace(c.AccessToken) == "" {
		return fmt.Errorf("INSTAGRAM_ACCESS_TOKEN is missing. Please configure your Meta Graph API Access Token")
	}

	domains := getTargetDomains(c.AccessToken)
	domain := domains[0]

	apiURL := fmt.Sprintf("https://%s/v19.0/%s?caption=%s&access_token=%s",
		domain, url.QueryEscape(mediaID), url.QueryEscape(caption), url.QueryEscape(c.AccessToken))

	resp, err := c.HTTPClient.Post(apiURL, "application/json", nil)
	if err != nil {
		return fmt.Errorf("network error updating caption: %w", err)
	}
	defer resp.Body.Close()

	var resData struct {
		Success bool `json:"success"`
		Error   *struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		} `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&resData); err != nil {
		return fmt.Errorf("failed to parse update caption response: %w", err)
	}

	if resData.Error != nil {
		return fmt.Errorf("Meta API Error (%d): %s", resData.Error.Code, resData.Error.Message)
	}

	if !resData.Success {
		return fmt.Errorf("Meta API did not accept caption update for post ID '%s'", mediaID)
	}

	return nil
}
