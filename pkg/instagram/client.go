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

type rateLimitingTransport struct {
	base http.RoundTripper
}

func (t *rateLimitingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := GlobalLimiter.Throttle(); err != nil {
		return nil, err
	}
	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}
	resp, err := base.RoundTrip(req)
	GlobalLimiter.HandleResponse(resp, err)
	return resp, err
}

// NewClient creates a new Meta Graph API client with automatic rate limiting.
func NewClient(accessToken, accountID string) *Client {
	return &Client{
		AccessToken: accessToken,
		AccountID:   accountID,
		HTTPClient: &http.Client{
			Timeout:   20 * time.Second,
			Transport: &rateLimitingTransport{base: http.DefaultTransport},
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
