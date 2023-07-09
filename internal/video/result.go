package video

import (
	"github.com/raspi/youtubeapi/internal/shared"
	"time"
)

// TODO: move to shared?

type Result struct {
	//Kind     string          `json:"kind"` // youtube#videoListResponse
	Etag     string          `json:"etag"` // for caching
	Items    []Item          `json:"items"`
	PageInfo shared.PageInfo `json:"pageInfo"` // n Results etc
}

type Item struct {
	//Kind           string         `json:"kind"` // youtube#video
	Etag           string         `json:"etag"` // for caching
	ID             string         `json:"id"`   // Video ID
	Snippet        Snippet        `json:"snippet"`
	ContentDetails ContentDetails `json:"contentDetails"`
	Status         Status         `json:"status"`
	Statistics     Statistics     `json:"statistics"`
	TopicDetails   TopicDetails   `json:"topicDetails"`
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

type ContentDetails struct {
	Duration        string `json:"duration"`   // ISO 8601
	Dimension       string `json:"dimension"`  // 2d
	Definition      string `json:"definition"` // hd
	Caption         string `json:"caption"`
	LicensedContent bool   `json:"licensedContent"`
	//ContentRating   ContentRating `json:"contentRating"`
	Projection string `json:"projection"`
}

type Status struct {
	UploadStatus        string `json:"uploadStatus"`  // processed
	PrivacyStatus       string `json:"privacyStatus"` // public
	License             string `json:"license"`       // youtube
	Embeddable          bool   `json:"embeddable"`
	PublicStatsViewable bool   `json:"publicStatsViewable"`
	MadeForKids         bool   `json:"madeForKids"`
}

type Statistics struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}

type TopicDetails struct {
	TopicCategories []string `json:"topicCategories"` // URLs?
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
