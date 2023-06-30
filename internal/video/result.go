package video

import (
	"github.com/raspi/youtubeapi/internal/shared"
	"time"
)

// TODO: move to shared?

type Result struct {
	Kind     string          `json:"kind"`
	Etag     string          `json:"etag"`
	Items    []Item          `json:"items"`
	PageInfo shared.PageInfo `json:"pageInfo"`
}

type Item struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	ID      string  `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	PublishedAt          time.Time          `json:"publishedAt"`
	ChannelID            string             `json:"channelId"`
	Title                string             `json:"title"`
	Description          string             `json:"description"`
	Thumbnails           *shared.Thumbnails `json:"thumbnails,omitempty"`
	ChannelTitle         string             `json:"channelTitle"`
	CategoryID           string             `json:"categoryId"`
	LiveBroadcastContent string             `json:"liveBroadcastContent"`
	DefaultLanguage      string             `json:"defaultLanguage"`
	Localized            Localized          `json:"localized"`
	DefaultAudioLanguage string             `json:"defaultAudioLanguage"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
