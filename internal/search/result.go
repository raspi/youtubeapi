package search

import (
	"github.com/raspi/youtubeapi/internal/shared"
	"time"
)

// TODO: move to shared?

type Result struct {
	//Kind          string          `json:"kind"`
	Etag          string          `json:"etag"` // for cache
	NextPageToken *string         `json:"nextPageToken,omitempty"`
	RegionCode    string          `json:"regionCode"`
	PageInfo      shared.PageInfo `json:"pageInfo"`
	Items         []Item          `json:"items"`
}

type Item struct {
	//Kind    string  `json:"kind"`
	Etag    string  `json:"etag"` // for cache
	ID      ID      `json:"id,omitempty"`
	Snippet Snippet `json:"snippet"`
}

type ID struct {
	Kind       string  `json:"kind"` // channel, playlist, video
	VideoID    *string `json:"videoId,omitempty"`
	ChannelID  *string `json:"channelId,omitempty"`
	PlaylistID *string `json:"playlistId,omitempty"`
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
