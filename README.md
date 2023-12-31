# youtubeapi

![GitHub release (latest by date)](https://img.shields.io/github/v/release/raspi/youtubeapi?style=for-the-badge)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/raspi/youtubeapi?style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspi/youtubeapi)](https://goreportcard.com/report/github.com/raspi/youtubeapi)


YouTube REST API v3 client for Go. 
For searching and getting details of YouTube videos, channels and playlists.

## Currently implemented

* Video details
* Channel details
* Search
* Playlists

## Example

```go
yt := youtubeapi.New(httpc, apiKey)
item, err := yt.GetVideo(`dQw4w9WgXcQ`)
if err != nil {
	panic(err)
}

fmt.Printf(`%v %v %v`, item.Snippet.ChannelID, item.Snippet.PublishedAt, item.Snippet.Title)

ch, err := yt.GetChannel(`UCuAXFkgsw1L7xaCfnd5JJOw`)
if err != nil {
	panic(err)
}

fmt.Printf(`%v %v %v`, ch.Snippet.PublishedAt, ch.Snippet.Title, ch.Snippet.Description)

custom := map[string]string{
	`regionCode`:        `US`,
	`relevanceLanguage`: `en`,
}

res, err := yt.Search(`never gonna give you up`, custom)
if err != nil {
	panic(err)
}

fmt.Printf(`%v`, res)
```

## References

* https://developers.google.com/youtube/v3
* [internal/video](internal/video) - https://developers.google.com/youtube/v3/docs/videos/list
* [internal/channel](internal/channel) - https://developers.google.com/youtube/v3/docs/channels
* [internal/search](internal/search) - https://developers.google.com/youtube/v3/docs/search/list
* [internal/playlist](internal/playlist) - https://developers.google.com/youtube/v3/docs/playlists/list
* [internal/playlistitem](internal/playlistitem) - https://developers.google.com/youtube/v3/docs/playlistItems/list
* https://console.cloud.google.com/apis/api/youtube.googleapis.com
