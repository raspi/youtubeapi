package channel

import (
	"github.com/raspi/youtubeapi/internal/shared"
	"time"
)

// TODO: move to shared?

type Result struct {
	//Kind     string          `json:"kind"`
	Etag     string          `json:"etag"` // for caching
	PageInfo shared.PageInfo `json:"pageInfo"`
	Items    []Item          `json:"items"`
}

type Item struct {
	//Kind    string  `json:"kind"`
	Etag      string  `json:"etag"` // for caching
	ChannelID string  `json:"id"`   // Channel ID
	Snippet   Snippet `json:"snippet"`
}

type Snippet struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	CustomURL   string             `json:"customUrl"`
	PublishedAt time.Time          `json:"publishedAt"`
	Thumbnails  *shared.Thumbnails `json:"thumbnails,omitempty"`
	Localized   Localized          `json:"localized"`
	Country     string             `json:"country"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
