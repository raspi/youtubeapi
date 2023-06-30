package search

import (
	"github.com/raspi/youtubeapi/internal/shared"
	"time"
)

// TODO: move to shared?

type Result struct {
	Kind          string          `json:"kind"`
	Etag          string          `json:"etag"`
	NextPageToken string          `json:"nextPageToken"`
	RegionCode    string          `json:"regionCode"`
	PageInfo      shared.PageInfo `json:"pageInfo"`
	Items         []Item          `json:"items"`
}

type Item struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	ID      ID      `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type ID struct {
	Kind    string `json:"kind"`
	VideoID string `json:"videoId"`
}

type Snippet struct {
	PublishedAt          time.Time         `json:"publishedAt"`
	ChannelID            string            `json:"channelId"`
	Title                string            `json:"title"`
	Description          string            `json:"description"`
	Thumbnails           shared.Thumbnails `json:"thumbnails"`
	ChannelTitle         string            `json:"channelTitle"`
	LiveBroadcastContent string            `json:"liveBroadcastContent"`
	PublishTime          string            `json:"publishTime"`
}
