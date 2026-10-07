package instagram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go-insta-cli/pkg/posts"
)

// Client handles network communications with Instagram APIs.
type Client struct {
	SessionID  string
	HTTPClient *http.Client
}

// NewClient creates a new Instagram HTTP client using the provided session ID or Cookie string.
func NewClient(sessionID string) *Client {
	return &Client{
		SessionID: sessionID,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// GraphAPIMediaData represents the JSON object returned by Meta Graph API /me/media.
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

// GraphAPIMediaResponse represents the response wrapper for Graph API media requests.
type GraphAPIMediaResponse struct {
	Data  []GraphAPIMediaData `json:"data"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    int    `json:"code"`
	} `json:"error"`
}

// InstagramWebProfileResponse maps the JSON structure returned by web_profile_info API.
type InstagramWebProfileResponse struct {
	Data struct {
		User struct {
			ID                       string `json:"id"`
			Username                 string `json:"username"`
			FullName                 string `json:"full_name"`
			Biography                string `json:"biography"`
			ProfilePicURL            string `json:"profile_pic_url_hd"`
			EdgeOwnerToTimelineMedia struct {
				Count int `json:"count"`
				Edges []struct {
					Node struct {
						ID                  string `json:"id"`
						DisplayURL          string `json:"display_url"`
						IsVideo             bool   `json:"is_video"`
						TakenAtTimestamp    int64  `json:"taken_at_timestamp"`
						EdgeMediaToCaption struct {
							Edges []struct {
								Node struct {
									Text string `json:"text"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_caption"`
						EdgeMediaToComment struct {
							Count int `json:"count"`
						} `json:"edge_media_to_comment"`
						EdgeLikedBy struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
		} `json:"user"`
	} `json:"data"`
	Status string `json:"status"`
}

// extractUserID parses the ds_user_id from the sessionid string or cookie set.
func extractUserID(sessionID string) string {
	if idx := strings.Index(sessionID, "ds_user_id="); idx != -1 {
		rest := sessionID[idx+len("ds_user_id="):]
		if semi := strings.Index(rest, ";"); semi != -1 {
			return strings.TrimSpace(rest[:semi])
		}
		return strings.TrimSpace(rest)
	}
	raw := sessionID
	if idx := strings.Index(raw, "sessionid="); idx != -1 {
		raw = raw[idx+len("sessionid="):]
	}
	if idx := strings.Index(raw, "%3A"); idx > 0 {
		return raw[:idx]
	}
	if idx := strings.Index(raw, ":"); idx > 0 {
		return raw[:idx]
	}
	return ""
}

// getFullCookieHeader constructs a full, sanitized Cookie header with harvested tokens.
func (c *Client) getFullCookieHeader() string {
	sessionID := strings.TrimSpace(c.SessionID)
	sessionID = strings.TrimRight(sessionID, "; ")
	if strings.Contains(sessionID, "csrftoken=") {
		return sessionID
	}

	userID := extractUserID(sessionID)
	header := sessionID
	if !strings.Contains(header, "sessionid=") {
		header = fmt.Sprintf("sessionid=%s", sessionID)
	}
	if userID != "" && !strings.Contains(header, "ds_user_id=") {
		header += fmt.Sprintf("; ds_user_id=%s", userID)
	}

	return header
}

// ValidateGraphToken checks if the Meta Graph API access token is active and valid.
func (c *Client) ValidateGraphToken(token string) bool {
	if token == "" {
		return false
	}
	url := fmt.Sprintf("https://graph.instagram.com/v19.0/me?fields=id,username&access_token=%s", token)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// FetchGraphAPIPosts fetches live posts using official Instagram Meta Graph API (graph.instagram.com/v19.0/me/media).
func (c *Client) FetchGraphAPIPosts(token string, limit int) ([]posts.Post, error) {
	if token == "" {
		return nil, fmt.Errorf("graph API access token is empty")
	}

	url := fmt.Sprintf("https://graph.instagram.com/v19.0/me/media?fields=id,caption,media_type,media_url,permalink,timestamp,like_count,comments_count&access_token=%s", token)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to reach Graph API: %w", err)
	}
	defer resp.Body.Close()

	var graphResp GraphAPIMediaResponse
	if err := json.NewDecoder(resp.Body).Decode(&graphResp); err != nil {
		return nil, fmt.Errorf("failed to parse Graph API response: %w", err)
	}

	if graphResp.Error != nil {
		return nil, fmt.Errorf("graph API error (%d): %s", graphResp.Error.Code, graphResp.Error.Message)
	}

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

// FetchUserPosts retrieves real Instagram posts attempting Graph API, HTML profile scraping, and Web API.
func (c *Client) FetchUserPosts(username string, limit int) ([]posts.Post, error) {
	// Strategy 1: Attempt Meta Graph API fetch if SessionID is a Graph Access Token
	if strings.HasPrefix(c.SessionID, "IG") || strings.HasPrefix(c.SessionID, "EAA") {
		postsList, err := c.FetchGraphAPIPosts(c.SessionID, limit)
		if err == nil {
			return postsList, nil
		}
	}

	if username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}

	cookieHeader := c.getFullCookieHeader()

	// Strategy 2: Web HTML Profile Scraping with Full Browser Cookies & Headers
	profileURL := fmt.Sprintf("https://www.instagram.com/%s/", username)
	reqHTML, err := http.NewRequest("GET", profileURL, nil)
	if err == nil {
		reqHTML.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
		reqHTML.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		reqHTML.Header.Set("Accept-Language", "en-US,en;q=0.9")
		reqHTML.Header.Set("Cookie", cookieHeader)

		respHTML, errHTML := c.HTTPClient.Do(reqHTML)
		if errHTML == nil && (respHTML.StatusCode == http.StatusOK || respHTML.StatusCode == http.StatusFound) {
			bodyBytes, _ := io.ReadAll(respHTML.Body)
			respHTML.Body.Close()
			htmlStr := string(bodyBytes)

			// Extract posts embedded inside profile HTML
			scrapedPosts := parseHTMLPosts(htmlStr, limit)
			if len(scrapedPosts) > 0 {
				return scrapedPosts, nil
			}
			// If HTML request returned 200 OK and valid user profile payload, return real empty post array
			if strings.Contains(htmlStr, username) || strings.Contains(htmlStr, "PolarisViewer") || respHTML.StatusCode == http.StatusOK {
				return scrapedPosts, nil
			}
		}
	}

	// Strategy 3: Standard Web Profile Info API
	url := fmt.Sprintf("https://www.instagram.com/api/v1/users/web_profile_info/?username=%s", username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("X-IG-App-ID", "936619743392459")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Referer", fmt.Sprintf("https://www.instagram.com/%s/", username))
	req.Header.Set("Cookie", cookieHeader)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("network error connecting to Instagram: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("instagram authentication failed (HTTP %d)", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("instagram API returned status HTTP %d", resp.StatusCode)
	}

	var profileResp InstagramWebProfileResponse
	if err := json.NewDecoder(resp.Body).Decode(&profileResp); err != nil {
		return nil, fmt.Errorf("failed to parse Instagram response JSON: %w", err)
	}

	edges := profileResp.Data.User.EdgeOwnerToTimelineMedia.Edges
	var result []posts.Post

	for i, edge := range edges {
		if limit > 0 && i >= limit {
			break
		}

		caption := ""
		if len(edge.Node.EdgeMediaToCaption.Edges) > 0 {
			caption = edge.Node.EdgeMediaToCaption.Edges[0].Node.Text
		}

		timestampStr := time.Unix(edge.Node.TakenAtTimestamp, 0).Format("2006-01-02 15:04")
		mediaType := "IMAGE"
		if edge.Node.IsVideo {
			mediaType = "VIDEO"
		}

		result = append(result, posts.Post{
			ID:           edge.Node.ID,
			Caption:      caption,
			Timestamp:    timestampStr,
			LikeCount:    edge.Node.EdgeLikedBy.Count,
			CommentCount: edge.Node.EdgeMediaToComment.Count,
			MediaType:    mediaType,
			MediaURL:     edge.Node.DisplayURL,
		})
	}

	return result, nil
}

// parseHTMLPosts parses post nodes embedded inside HTML page scripts.
func parseHTMLPosts(htmlStr string, limit int) []posts.Post {
	var results []posts.Post

	// Regular expressions to extract post display URLs and timestamps
	reID := regexp.MustCompile(`"id":"(\d+)"`)
	reURL := regexp.MustCompile(`"display_url":"(https:[^"]+)"`)
	reText := regexp.MustCompile(`"text":"([^"]+)"`)
	reLikes := regexp.MustCompile(`"edge_liked_by":\{"count":(\d+)\}`)
	reComments := regexp.MustCompile(`"edge_media_to_comment":\{"count":(\d+)\}`)
	reTime := regexp.MustCompile(`"taken_at_timestamp":(\d+)`)

	ids := reID.FindAllStringSubmatch(htmlStr, -1)
	urls := reURL.FindAllStringSubmatch(htmlStr, -1)
	texts := reText.FindAllStringSubmatch(htmlStr, -1)
	likes := reLikes.FindAllStringSubmatch(htmlStr, -1)
	comments := reComments.FindAllStringSubmatch(htmlStr, -1)
	times := reTime.FindAllStringSubmatch(htmlStr, -1)

	minCount := len(ids)
	if len(urls) < minCount {
		minCount = len(urls)
	}

	for i := 0; i < minCount; i++ {
		if limit > 0 && i >= limit {
			break
		}

		id := ids[i][1]
		displayURL := strings.ReplaceAll(urls[i][1], "\\u0026", "&")
		displayURL = strings.ReplaceAll(displayURL, "\\/", "/")

		caption := ""
		if i < len(texts) {
			caption = texts[i][1]
		}

		likeCount := 0
		if i < len(likes) {
			likeCount, _ = strconv.Atoi(likes[i][1])
		}

		commentCount := 0
		if i < len(comments) {
			commentCount, _ = strconv.Atoi(comments[i][1])
		}

		timestampStr := time.Now().Format("2006-01-02 15:04")
		if i < len(times) {
			if ts, err := strconv.ParseInt(times[i][1], 10, 64); err == nil {
				timestampStr = time.Unix(ts, 0).Format("2006-01-02 15:04")
			}
		}

		results = append(results, posts.Post{
			ID:           id,
			Caption:      caption,
			Timestamp:    timestampStr,
			LikeCount:    likeCount,
			CommentCount: commentCount,
			MediaType:    "IMAGE",
			MediaURL:     displayURL,
		})
	}

	return results
}
