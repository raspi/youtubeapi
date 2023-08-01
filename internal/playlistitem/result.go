package playlistitem

import (
	"github.com/raspi/youtubeapi/internal/shared"
	"time"
)

type Result struct {
	//Kind          string   `json:"kind"`
	Etag          string          `json:"etag"` // for caching
	NextPageToken *string         `json:"nextPageToken,omitempty"`
	Items         []Item          `json:"items"`
	PageInfo      shared.PageInfo `json:"pageInfo"` // next page etc
}

type Item struct {
	//Kind           string       `json:"kind"`
	Etag           string         `json:"etag"` // for caching
	ID             string         `json:"id"`
	Snippet        Snippet        `json:"snippet"`
	ContentDetails ContentDetails `json:"contentDetails"`
	Status         Status         `json:"status"`
}

type ContentDetails struct {
	VideoID          string    `json:"videoId"` // Video's id
	VideoPublishedAt time.Time `json:"videoPublishedAt"`
}

type Snippet struct {
	PlaylistID             string             `json:"playlistId"`
	Position               uint64             `json:"position"`            // position in playlist 0-N
	ChannelID              string             `json:"channelId"`           // YouTube channel ID
	VideoOwnerChannelID    string             `json:"videoOwnerChannelId"` // it might be other user's playlist
	PublishedAt            time.Time          `json:"publishedAt"`
	ResourceID             ResourceID         `json:"resourceId"`
	ChannelTitle           string             `json:"channelTitle"`           // YouTube channel name
	VideoOwnerChannelTitle string             `json:"videoOwnerChannelTitle"` // it might be other user's playlist
	Title                  string             `json:"title"`                  // Video title
	Description            string             `json:"description"`            // Video description
	Thumbnails             *shared.Thumbnails `json:"thumbnails,omitempty"`   // Video thumbnails
}

type ResourceID struct {
	//Kind    string `json:"kind"`
	VideoID string `json:"videoId"` // YouTube video ID
}

type Status struct {
	PrivacyStatus string `json:"privacyStatus"` // public
}
