package playlist

import (
	"github.com/raspi/youtubeapi/internal/shared"
	"time"
)

type Result struct {
	//Kind          string          `json:"kind"`
	Etag          string          `json:"etag"`                    // for caching
	NextPageToken *string         `json:"nextPageToken,omitempty"` // for `pageToken` parameter
	PageInfo      shared.PageInfo `json:"pageInfo"`                // How many results etc
	Items         []Item          `json:"items"`                   // Playlist
}

type Item struct {
	//Kind           string           `json:"kind"`
	Etag           string         `json:"etag"` // for caching
	ID             string         `json:"id"`   // Playlist ID
	Snippet        Snippet        `json:"snippet"`
	Status         Status         `json:"status"`
	ContentDetails ContentDetails `json:"contentDetails"` // How many videos etc
}

type ContentDetails struct {
	ItemCount uint64 `json:"itemCount"` // How many videos in playlist
}

type Snippet struct {
	ChannelID    string             `json:"channelId"`    // Channel id
	ChannelTitle string             `json:"channelTitle"` // Channel name
	PublishedAt  time.Time          `json:"publishedAt"`
	Title        string             `json:"title"`                 // Playlist name
	Description  *string            `json:"description,omitempty"` // Playlist description
	Thumbnails   *shared.Thumbnails `json:"thumbnails"`
	Localized    *Localized         `json:"localized"` // might be different when using `hl` parameter in query?
}

type Localized struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}

type Status struct {
	PrivacyStatus string `json:"privacyStatus"` // public
}
