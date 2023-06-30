package youtubeapi

import (
	"github.com/raspi/youtubeapi/internal/channel"
	"github.com/raspi/youtubeapi/internal/search"
	"github.com/raspi/youtubeapi/internal/video"
	"net/http"
)

// YoutubeAPI is REST client for YouTube REST API
// See [internal](internal) directory for implementation details
// See https://console.cloud.google.com/apis/api/youtube.googleapis.com/quotas
type YoutubeAPI struct {
	videoClient   *video.Client
	channelClient *channel.Client
	searchClient  *search.Client
}

func New(cl *http.Client, apikey string) *YoutubeAPI {
	if cl == nil {
		panic(`nil http client`)
	}

	if apikey == `` {
		panic(`empty api key`)
	}

	return &YoutubeAPI{
		videoClient:   video.New(apikey, cl),
		channelClient: channel.New(apikey, cl),
		searchClient:  search.New(apikey, cl),
	}
}

// GetVideo fetches video information from YouTube REST API, see internal/video for more details
func (api YoutubeAPI) GetVideo(id string) (*video.Item, error) {
	return api.videoClient.Get(id)
}

// GetChannel fetches channel information from YouTube REST API, see internal/channel for more details
func (api YoutubeAPI) GetChannel(id string) (*channel.Snippet, error) {
	return api.channelClient.Get(id)
}

// Search searches YouTube REST API with given query, see internal/search for more details
// Note: the search consumes more quota
// See https://developers.google.com/youtube/v3/docs/search/list
func (api YoutubeAPI) Search(query string, customParameters map[string]string) (*search.Result, error) {
	return api.searchClient.Search(query, customParameters)
}
