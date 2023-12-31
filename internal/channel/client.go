package channel

import (
	"encoding/json"
	"fmt"
	"github.com/raspi/youtubeapi/internal/shared"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const API_URL = "https://www.googleapis.com/youtube/v3/channels"

// Client is for YouTube HTTP REST API for channels
// See https://developers.google.com/youtube/v3/docs/channels
type Client struct {
	apiKey string
	cl     *http.Client
}

func New(apikey string, client *http.Client) *Client {
	return &Client{
		apiKey: apikey,
		cl:     client,
	}
}

// getIds fetches details of given channel ID(s)
//
// The part parameter specifies a comma-separated list of one or more channel resource properties that the API response will include.
//
// If the parameter identifies a property that contains child properties, the child properties will be included in the response. For example, in a channel resource, the contentDetails property contains other properties, such as the uploads properties. As such, if you set part=contentDetails, the API response will also contain all of those nested properties.
//
// The following list contains the part names that you can include in the parameter value:
//
//	auditDetails
//	brandingSettings
//	contentDetails
//	contentOwnerDetails
//	id
//	localizations
//	snippet
//	statistics
//	status
//	topicDetails
func (s Client) getIds(ids []string, parts []string) ([]Item, error) {
	if ids == nil {
		return nil, shared.ErrEmpty
	}

	if len(ids) == 0 {
		return nil, shared.ErrEmpty
	}

	if len(ids) > 50 {
		return nil, fmt.Errorf(`over 50 ids`)
	}

	for idx, id := range ids {
		if id == `` {
			return nil, shared.NewErrEmptyIdx(uint(idx))
		}
	}

	q := url.Values{}
	q.Set(`id`, strings.Join(ids, `,`)) // Channel IDs
	q.Set(`maxResults`, `50`)
	q.Set(`part`, strings.Join(parts, `,`))

	return s.fetchURL(q)
}

// GetIds see getIds()
func (s Client) GetIds(ids []string) ([]Item, error) {
	parts := []string{`id`, `snippet`, `status`, `contentDetails`, `statistics`}
	return s.getIds(ids, parts)
}

// GetIdsStats returns statistics of given channel IDs
// Since statistics changes more than other channel details,
// it has been split here for better caching (use ETag)
func (s Client) GetIdsStats(ids []string) ([]Item, error) {
	parts := []string{`id`, `statistics`}
	return s.getIds(ids, parts)
}

// GetChannelIdFromName tries to resolve channel name such as @FooBar to channel ID
// This might fail, so try with internal/search.Search() also
func (s Client) GetChannelIdFromName(name string) (*string, error) {
	if name == `` {
		return nil, shared.ErrEmpty
	}

	parts := []string{`id`, `snippet`}

	q := url.Values{}
	q.Set(`forUsername`, name) // Channel name
	q.Set(`maxResults`, `1`)
	q.Set(`part`, strings.Join(parts, `,`))

	items, err := s.fetchURL(q)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, nil
	}

	return &items[0].ChannelID, nil
}

func (s Client) fetchURL(q url.Values) ([]Item, error) {
	q.Set(`key`, s.apiKey)

	resp, err := s.cl.Get(API_URL + `?` + q.Encode())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		var r Result
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

	// Read error JSON
	var tmperr shared.APIError
	err = json.Unmarshal(content, &tmperr)
	if err != nil {
		return nil, err
	}

	return nil, tmperr.Error
}
