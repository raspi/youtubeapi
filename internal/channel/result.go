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
	Etag           string         `json:"etag"` // for caching
	ChannelID      string         `json:"id"`   // Channel ID
	Snippet        Snippet        `json:"snippet"`
	ContentDetails ContentDetails `json:"contentDetails"`
	Statistics     Statistics     `json:"statistics"`
	Status         Status         `json:"status"`
}

type Snippet struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	CustomURL   string             `json:"customUrl"` // @fooBar user name
	PublishedAt time.Time          `json:"publishedAt"`
	Thumbnails  *shared.Thumbnails `json:"thumbnails,omitempty"`
	Localized   Localized          `json:"localized"`
	Country     string             `json:"country"`
}

type ContentDetails struct {
	RelatedPlaylists RelatedPlaylists `json:"relatedPlaylists"`
}

type RelatedPlaylists struct {
	Likes   *string `json:"likes,omitempty"`
	Uploads string  `json:"uploads"` // All uploads playlist
}

type Statistics struct {
	ViewCount             string `json:"viewCount"`
	SubscriberCount       string `json:"subscriberCount"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
	VideoCount            string `json:"videoCount"`
}

type Status struct {
	PrivacyStatus     string `json:"privacyStatus"` // public
	IsLinked          bool   `json:"isLinked"`
	LongUploadsStatus string `json:"longUploadsStatus"`
	MadeForKids       bool   `json:"madeForKids"`
}

type Localized struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}
