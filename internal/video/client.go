package video

import (
	"encoding/json"
	"fmt"
	"github.com/raspi/youtubeapi/internal/shared"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const API_URL = "https://www.googleapis.com/youtube/v3/videos"

// Client is a YouTube REST API client for videos
// See https://developers.google.com/youtube/v3/docs/videos/list
type Client struct {
	apiKey string
	cl     *http.Client
}

// New returns a new YouTube Video client
// See Client
func New(apikey string, client *http.Client) *Client {
	return &Client{
		apiKey: apikey,
		cl:     client,
	}
}

// Parts:
// the snippet property contains the channelId, title, description, tags, and categoryId properties.
// As such, if you set part=snippet, the API response will contain all of those properties.
// The following list contains the part names that you can include in the parameter value:
//
// contentDetails (duration, definition (hd), licensed) See: ContentDetails
// fileDetails (auth?)
// id (video id) string
// liveStreamingDetails
// localizations
// player (returns iframe HTML)
// processingDetails (auth?)
// recordingDetails (empty?)
// snippet (default) See: Snippet
// statistics (views, etc) See: Statistics
// status (is embeddable) See: Status
// suggestions
// topicDetails (entertainment, etc) See: TopicDetails
func (s Client) getIds(ids []string, parts []string) ([]Item, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf(`no ids given`)
	}

	if len(ids) > 50 {
		return nil, fmt.Errorf(`over 50 ids`)
	}

	for idx, id := range ids {
		if id == `` {
			return nil, fmt.Errorf(`empty id at pos %d`, idx)
		}
	}

	if len(parts) == 0 {
		return nil, fmt.Errorf(`empty parts`)
	}

	// Query string
	q := url.Values{}
	q.Set(`id`, strings.Join(ids, `,`))
	q.Set(`key`, s.apiKey)
	q.Set(`part`, strings.Join(parts, `,`))

	// Make HTTP GET request
	resp, err := s.cl.Get(API_URL + `?` + q.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read JSON
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(`couldn't load body: %v`, err)
	}

	if resp.StatusCode == http.StatusOK {
		// Got valid result
		var r Result

		// Parse JSON
		err = json.Unmarshal(content, &r)
		if err != nil {
			return nil, err
		}

		if r.PageInfo.TotalResults == 0 || len(r.Items) == 0 {
			// No results
			return []Item{}, nil
		}

		return r.Items, nil
	}

	// Something went wrong
	var tmperr shared.APIError
	err = json.Unmarshal(content, &tmperr)
	if err != nil {
		return nil, fmt.Errorf(`JSON parsing error (error result): %v`, err)
	}

	return nil, tmperr.Error
}

// GetIds returns details for multiple videos
// Use ETag for caching
func (s Client) GetIds(ids []string) ([]Item, error) {
	parts := []string{
		`id`, `snippet`, `status`, `contentDetails`, `topicDetails`,
	}
	return s.getIds(ids, parts)
}

// GetIdsStats returns statistics of given video IDs
// Since statistics changes more than other video details,
// it has been split here for better caching (use ETag)
func (s Client) GetIdsStats(ids []string) ([]Item, error) {
	parts := []string{
		`id`, `statistics`,
	}
	return s.getIds(ids, parts)
}
