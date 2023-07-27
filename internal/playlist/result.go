package playlist

import (
	"github.com/raspi/youtubeapi/internal/shared"
	"time"
)

type Result struct {
	//Kind          string          `json:"kind"`
	Etag          string          `json:"etag"` // for caching
	NextPageToken *string         `json:"nextPageToken,omitempty"`
	PageInfo      shared.PageInfo `json:"pageInfo"`
	Items         []Item          `json:"items"`
}

type Item struct {
	//Kind           string           `json:"kind"`
	Etag           string         `json:"etag"` // for caching
	ID             string         `json:"id"`
	Snippet        Snippet        `json:"snippet"`
	Status         Status         `json:"status"`
	ContentDetails ContentDetails `json:"contentDetails"`
}

type ContentDetails struct {
	ItemCount int64 `json:"itemCount"` // How many videos in playlist
}

type Snippet struct {
	ChannelID    string             `json:"channelId"`    // Channel id
	ChannelTitle string             `json:"channelTitle"` // Channel name
	PublishedAt  time.Time          `json:"publishedAt"`
	Title        string             `json:"title"` // Playlist name
	Description  string             `json:"description"`
	Thumbnails   *shared.Thumbnails `json:"thumbnails"`
	Localized    *Localized         `json:"localized"` // might be different when using `hl` parameter in query?
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Status struct {
	PrivacyStatus string `json:"privacyStatus"` // public
}
