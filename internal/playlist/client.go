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

const API_URL = "https://www.googleapis.com/youtube/v3/playlistItems"

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

func (s Client) GetId(id string, custom map[string]string) ([]Item, error) {
	if id == `` {
		return nil, fmt.Errorf(`no id`)
	}

	// The following list contains the part names that you can include in the parameter value:
	//    contentDetails, id, snippet, status
	parts := []string{`id`, `snippet`, `status`, `contentDetails`}

	q := url.Values{}
	q.Set(`playlistId`, id) // Playlist ID
	//q.Set(`videoId`, strings.Join(ids, `,`)) // Return playlist(s) which contains this video id
	q.Set(`maxResults`, `50`)
	q.Set(`part`, strings.Join(parts, `,`))
	q.Set(`key`, s.apiKey)

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

// GetItemIds fetches information about video item INSIDE some playlist
// You can get the playlist item ids with GetId()
func (s Client) GetItemIds(ids []string, custom map[string]string) ([]Item, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf(`no ids given`)
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
	q.Set(`key`, s.apiKey)

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
