package youtubeapi

import (
	"github.com/raspi/youtubeapi/internal/channel"
	"github.com/raspi/youtubeapi/internal/playlist"
	"github.com/raspi/youtubeapi/internal/playlistitem"
	"github.com/raspi/youtubeapi/internal/search"
	"github.com/raspi/youtubeapi/internal/shared"
	"github.com/raspi/youtubeapi/internal/video"
	"net/http"
)

// YoutubeAPI is REST client for YouTube REST API v3
// See [internal](internal) directory for implementation details
// See https://console.cloud.google.com/apis/api/youtube.googleapis.com/quotas
type YoutubeAPI struct {
	videoClient        *video.Client
	channelClient      *channel.Client
	searchClient       *search.Client
	playlistItemClient *playlistitem.Client // More details, video IDs, etc
	playlistClient     *playlist.Client     // Only video count, etc
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
		videoClient:        video.New(apikey, cl),
		channelClient:      channel.New(apikey, cl),
		searchClient:       search.New(apikey, cl),
		playlistItemClient: playlistitem.New(apikey, cl),
		playlistClient:     playlist.New(apikey, cl),
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

// GetChannelStats fetches statistics of given channel ID from YouTube REST API, see internal/channel for more details
func (api YoutubeAPI) GetChannelStats(channelId string) (*channel.Item, error) {
	res, err := api.GetChannelStatsMany([]string{channelId})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return &res[0], nil
}

// GetChannelStatsMany fetches statistics of given channel ID(s) from YouTube REST API, see internal/channel for more details
func (api YoutubeAPI) GetChannelStatsMany(channelIds []string) ([]channel.Item, error) {
	return api.channelClient.GetIdsStats(channelIds)
}

// Search searches YouTube REST API with given query, see internal/search for more details
// Note: the search consumes more API quota
// See https://developers.google.com/youtube/v3/docs/search/list
func (api YoutubeAPI) Search(query string, customParameters map[string]string) (*search.Result, error) {
	return api.searchClient.Search(query, customParameters)
}

// GetPlaylistVideos fetches given playlists' videos
func (api YoutubeAPI) GetPlaylistVideos(playlistId string, customParameters map[string]string) ([]playlistitem.Item, shared.Meta, error) {
	return api.playlistItemClient.GetPlaylistIdItems(playlistId, customParameters)
}

// GetPlaylistVideoItems fetches items (videos) which are *in* some playlistClient
func (api YoutubeAPI) GetPlaylistVideoItems(plItemIds []string, customParameters map[string]string) ([]playlistitem.Item, error) {
	return api.playlistItemClient.GetItemIds(plItemIds, customParameters)
}

// GetChannelPlaylists fetches playlists meta information (no videos) from given channel ID
func (api YoutubeAPI) GetChannelPlaylists(channelId string, customParameters map[string]string) ([]playlist.Item, shared.Meta, error) {
	return api.playlistClient.GetChannelPlaylists(channelId, customParameters)
}

// GetPlaylist fetches given playlists' meta information (no videos)
func (api YoutubeAPI) GetPlaylist(playlistId string) (*playlist.Item, error) {
	res, err := api.GetPlaylists([]string{playlistId})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return &res[0], nil
}

// GetPlaylists fetches given list of playlists' meta information (no videos)
func (api YoutubeAPI) GetPlaylists(playlistIds []string) ([]playlist.Item, error) {
	return api.playlistClient.GetPlaylists(playlistIds)
}
