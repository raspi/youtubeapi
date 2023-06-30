package channel

import (
	"github.com/raspi/youtubeapi/internal/shared"
	"time"
)

// TODO: move to shared?

type Result struct {
	Kind     string          `json:"kind"`
	Etag     string          `json:"etag"`
	PageInfo shared.PageInfo `json:"pageInfo"`
	Items    []Item          `json:"items"`
}

type Item struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	ID      string  `json:"id"`
	Snippet Snippet `json:"snippet"`
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
