package playlist

import (
	"encoding/json"
	"fmt"
	"github.com/raspi/youtubeapi/internal/shared"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const API_URL = "https://www.googleapis.com/youtube/v3/playlists"

// Client ...
// See https://developers.google.com/youtube/v3/docs/playlists/list
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

// GetChannelPlaylists fetches a list of playlist(s) created by some channel
// Use custom[`pageToken`] = meta.NextPageToken for iterating all pages
func (s Client) GetChannelPlaylists(channelId string, custom map[string]string) ([]Item, shared.Meta, error) {
	meta := shared.Meta{}
	if channelId == `` {
		return nil, meta, fmt.Errorf(`no id`)
	}

	// The following list contains the part names that you can include in the parameter value:
	//    contentDetails, id, localizations, player, snippet, status
	parts := []string{`id`, `snippet`, `status`, `contentDetails`}

	q := url.Values{}
	q.Set(`channelId`, channelId) // Channel ID
	q.Set(`maxResults`, `50`)
	q.Set(`part`, strings.Join(parts, `,`))
	q.Set(`key`, s.apiKey)

	if custom != nil {
		for k, v := range custom {
			switch k {
			case `key`, `part`, `channelId`, `id`:
				// Do not allow to change these key(s)
				continue
			}

			q.Set(k, v)
		}
	}

	resp, err := s.cl.Get(API_URL + `?` + q.Encode())
	if err != nil {
		return nil, meta, err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, meta, err
	}

	if resp.StatusCode == http.StatusOK {
		var r Result
		err = json.Unmarshal(content, &r)
		if err != nil {
			return nil, meta, err
		}

		if r.PageInfo.TotalResults == 0 || len(r.Items) == 0 {
			// No results
			return []Item{}, meta, nil
		}

		// For custom parameter so that next page can be queried
		meta.NextPageToken = r.NextPageToken
		meta.ETag = r.Etag

		return r.Items, meta, nil
	}

	var tmperr shared.APIError
	err = json.Unmarshal(content, &tmperr)
	if err != nil {
		return nil, meta, err
	}

	return nil, meta, tmperr.Error
}

func (s Client) GetPlaylists(ids []string) ([]Item, error) {
	if ids == nil {
		return nil, fmt.Errorf(`nil ids`)
	}

	if len(ids) == 0 {
		return nil, fmt.Errorf(`no ids`)
	}
	if len(ids) > 50 {
		return nil, fmt.Errorf(`over 50 ids`)
	}

	for idx, id := range ids {
		if id == `` {
			return nil, fmt.Errorf(`empty id at pos %d`, idx)
		}
	}

	// The following list contains the part names that you can include in the parameter value:
	//    contentDetails, id, localizations, player, snippet, status
	parts := []string{`id`, `snippet`, `status`, `contentDetails`}

	q := url.Values{}
	q.Set(`id`, strings.Join(ids, ``)) // Playlist IDs
	q.Set(`maxResults`, `50`)
	q.Set(`part`, strings.Join(parts, `,`))
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

	var tmperr shared.APIError
	err = json.Unmarshal(content, &tmperr)
	if err != nil {
		return nil, err
	}

	return nil, tmperr.Error
}
