package instagram

import (
	"net/http"
	"time"
)

// Client handles Meta Instagram Graph API requests.
type Client struct {
	AccessToken string
	AccountID   string
	HTTPClient  *http.Client
}

// NewClient creates a new Meta Graph API client.
func NewClient(accessToken, accountID string) *Client {
	return &Client{
		AccessToken: accessToken,
		AccountID:   accountID,
		HTTPClient: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

// GraphAPIUserResponse represents the Meta Graph API /me response.
type GraphAPIUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Error    *struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FBTraceID string `json:"fbtrace_id"`
	} `json:"error"`
}

// GraphAPIMediaData represents media items returned by /me/media.
type GraphAPIMediaData struct {
	ID            string `json:"id"`
	Caption       string `json:"caption"`
	MediaType     string `json:"media_type"`
	MediaURL      string `json:"media_url"`
	Permalink     string `json:"permalink"`
	Timestamp     string `json:"timestamp"`
	LikeCount     int    `json:"like_count"`
	CommentsCount int    `json:"comments_count"`
}

// GraphAPIMediaResponse represents the media query list response.
type GraphAPIMediaResponse struct {
	Data  []GraphAPIMediaData `json:"data"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    int    `json:"code"`
	} `json:"error"`
}
