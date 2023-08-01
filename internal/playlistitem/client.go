package playlistitem

import (
	"encoding/json"
	"fmt"
	"github.com/raspi/youtubeapi/internal/shared"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const API_URL = "https://www.googleapis.com/youtube/v3/playlistItems"

// Client is HTTP client for YouTube playlist item REST API v3
// See https://developers.google.com/youtube/v3/docs/playlistItems/list
type Client struct {
	apiKey string
	cl     *http.Client
}

// New uses YouTube playlist items REST API
func New(apikey string, client *http.Client) *Client {
	return &Client{
		apiKey: apikey,
		cl:     client,
	}
}

// GetPlaylistIdItems fetches a playlist with given ID
// Use custom[`pageToken`] = meta.NextPageToken for iterating all pages
func (s Client) GetPlaylistIdItems(id string, custom map[string]string) ([]Item, *shared.Meta, error) {
	if id == `` {
		return nil, nil, shared.ErrEmpty
	}

	// The following list contains the part names that you can include in the parameter value:
	//    contentDetails, id, snippet, status
	parts := []string{`id`, `snippet`, `status`, `contentDetails`}

	q := url.Values{}
	q.Set(`playlistId`, id) // Playlist ID
	//q.Set(`videoId`, strings.Join(ids, `,`)) // Return playlist(s) which contains this video id
	q.Set(`maxResults`, `50`)
	q.Set(`part`, strings.Join(parts, `,`))

	if custom != nil {
		for k, v := range custom {
			switch k {
			case `key`, `part`, `playlistId`, `id`:
				// Do not allow to change these key(s)
				continue
			}

			q.Set(k, v)
		}
	}

	return s.fetchUrl(q)
}

// GetItemIds fetches information about video item INSIDE some playlist
// You can get the playlist item ids with GetPlaylistIdItems()
// Use custom[`pageToken`] = meta.NextPageToken for iterating all pages
func (s Client) GetItemIds(ids []string, custom map[string]string) ([]Item, error) {
	if len(ids) == 0 {
		return nil, shared.ErrEmpty
	}

	for idx, id := range ids {
		if id == `` {
			return nil, fmt.Errorf(`empty id at pos %d`, idx)
		}
	}
	// The following list contains the part names that you can include in the parameter value:
	//    contentDetails, id, snippet, status
	parts := []string{`id`, `snippet`, `status`, `contentDetails`}

	q := url.Values{}
	q.Set(`id`, strings.Join(ids, `,`)) // Playlist item IDs
	q.Set(`maxResults`, `50`)
	q.Set(`part`, strings.Join(parts, `,`))

	if custom != nil {
		for k, v := range custom {
			switch k {
			case `key`, `part`, `playlistId`, `id`:
				// Do not allow to change these key(s)
				continue
			}

			q.Set(k, v)
		}
	}

	items, _, err := s.fetchUrl(q)
	if err != nil {
		return nil, err
	}
	return items, nil

}

func (s Client) fetchUrl(q url.Values) ([]Item, *shared.Meta, error) {
	q.Set(`key`, s.apiKey)

	resp, err := s.cl.Get(API_URL + `?` + q.Encode())
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode == http.StatusOK {
		var r Result
		err = json.Unmarshal(content, &r)
		if err != nil {
			return nil, nil, err
		}

		if r.PageInfo.TotalResults == 0 || len(r.Items) == 0 {
			// No results
			return []Item{}, nil, nil
		}

		return r.Items,
			&shared.Meta{
				ETag:          r.Etag,
				NextPageToken: r.NextPageToken,
			}, nil
	}

	// Read error JSON
	var tmperr shared.APIError
	err = json.Unmarshal(content, &tmperr)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, tmperr.Error

}
