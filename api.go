package youtubeapi

import (
	"github.com/raspi/youtubeapi/internal/channel"
	"github.com/raspi/youtubeapi/internal/search"
	"github.com/raspi/youtubeapi/internal/video"
	"net/http"
)

// YoutubeAPI is REST client for YouTube REST API v3
// See [internal](internal) directory for implementation details
// See https://console.cloud.google.com/apis/api/youtube.googleapis.com/quotas
type YoutubeAPI struct {
	videoClient   *video.Client
	channelClient *channel.Client
	searchClient  *search.Client
}

// New creates a new YouTube HTTP REST API v3 client
//
// You can use caching in http.Client .Transport http.RoundTripper:
//
//	   httpc := http.DefaultClient
//		  httpc.Transport = myCachingRoundTripper
//		  ytc := New(httpc, `...`)
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
func (api YoutubeAPI) GetVideo(videoId string) (*video.Item, error) {
	res, err := api.GetVideos([]string{videoId})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return &res[0], err
}

// GetVideos fetches multiple video information from YouTube REST API, see internal/video for more details
func (api YoutubeAPI) GetVideos(videoIds []string) ([]video.Item, error) {
	return api.videoClient.GetIds(videoIds)
}

// GetVideoStats fetches video statistics from YouTube REST API, see internal/video for more details
func (api YoutubeAPI) GetVideoStats(videoId string) (*video.Item, error) {
	res, err := api.GetVideosStats([]string{videoId})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return &res[0], nil
}

// GetVideosStats fetches multiple video information from YouTube REST API, see internal/video for more details
func (api YoutubeAPI) GetVideosStats(videoIds []string) ([]video.Item, error) {
	return api.videoClient.GetIdsStats(videoIds)
}

// GetChannel fetches channel information from YouTube REST API, see internal/channel for more details
func (api YoutubeAPI) GetChannel(channelId string) (*channel.Item, error) {
	res, err := api.GetChannels([]string{channelId})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return &res[0], nil
}

// GetChannels fetches multiple channel information from YouTube REST API, see internal/channel for more details
func (api YoutubeAPI) GetChannels(channelIds []string) ([]channel.Item, error) {
	return api.channelClient.GetIds(channelIds)
}

// Search searches YouTube REST API with given query, see internal/search for more details
// Note: the search consumes more API quota
// See https://developers.google.com/youtube/v3/docs/search/list
func (api YoutubeAPI) Search(query string, customParameters map[string]string) (*search.Result, error) {
	return api.searchClient.Search(query, customParameters)
}
