package search

import (
	"encoding/json"
	"github.com/raspi/youtubeapi/internal/shared"
	"io"
	"net/http"
	"net/url"
)

const API_URL = "https://youtube.googleapis.com/youtube/v3/search"

// Client uses YouTube search REST API
// See https://developers.google.com/youtube/v3/docs/search
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

// Search searches YouTube REST API with given query
// See https://developers.google.com/youtube/v3/docs/search/list
//
// Parameter `custom` is a map which may contain custom parameters set by user.
// For example
// - `pageToken` which points to `nextPageToken` or `prevPageToken`
// - `publishedAfter` date
// - `relevanceLanguage` language code
// - `regionCode` such as US, UK, FI, .. See https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
func (s Client) Search(query string, custom map[string]string) (*Result, error) {
	q := url.Values{}
	q.Set(`part`, `snippet`)
	q.Set(`order`, `date`)
	q.Set(`maxResults`, `50`)
	q.Set(`safeSearch`, `none`)
	q.Set(`q`, query)
	q.Set(`key`, s.apiKey)

	if custom != nil {
		for k, v := range custom {
			switch k {
			case `q`, `key`:
				// Do not allow to change query or API key
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
			return nil, shared.NotFound{}
		}

		return &r, nil
	}

	var tmperr shared.APIError
	err = json.Unmarshal(content, &tmperr)
	if err != nil {
		return nil, err
	}

	return nil, tmperr.Error
}
