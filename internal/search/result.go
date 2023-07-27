package search

import (
	"github.com/raspi/youtubeapi/internal/shared"
	"time"
)

// TODO: move to shared?

type Result struct {
	//Kind          string          `json:"kind"`
	Etag          string          `json:"etag"`
	NextPageToken string          `json:"nextPageToken"`
	RegionCode    string          `json:"regionCode"`
	PageInfo      shared.PageInfo `json:"pageInfo"`
	Items         []Item          `json:"items"`
}

type Item struct {
	//Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	ID      ID      `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type ID struct {
	//Kind    string `json:"kind"`
	VideoID string `json:"videoId"`
}

type Snippet struct {
	ChannelID            string            `json:"channelId"`
	ChannelTitle         string            `json:"channelTitle"`
	PublishedAt          time.Time         `json:"publishedAt"`
	PublishTime          time.Time         `json:"publishTime"`
	Title                string            `json:"title"`
	Description          string            `json:"description"`
	Thumbnails           shared.Thumbnails `json:"thumbnails"`
	LiveBroadcastContent string            `json:"liveBroadcastContent"` // "none"
}
