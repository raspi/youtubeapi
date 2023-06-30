package video

import (
	"encoding/json"
	"github.com/raspi/youtubeapi/internal/shared"
	"io"
	"net/http"
	"net/url"
)

const API_URL = "https://www.googleapis.com/youtube/v3/videos"

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

func (s Client) Get(id string) (*Item, error) {
	q := url.Values{}
	q.Set(`id`, id)
	q.Set(`key`, s.apiKey)
	q.Set(`part`, `snippet`)

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

		return &r.Items[0], nil
	}

	var tmperr shared.APIError
	err = json.Unmarshal(content, &tmperr)
	if err != nil {
		return nil, err
	}

	return nil, tmperr.Error
}
