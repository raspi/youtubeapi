# youtubeapi

![GitHub release (latest by date)](https://img.shields.io/github/v/release/raspi/youtubeapi?style=for-the-badge)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/raspi/youtubeapi?style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspi/youtubeapi)](https://goreportcard.com/report/github.com/raspi/youtubeapi)


YouTube REST API v3 client for Go. 
For searching and getting details of YouTube videos and channels.

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

* https://developers.google.com/youtube/v3/docs/videos/list
* https://developers.google.com/youtube/v3/docs/channels
* https://developers.google.com/youtube/v3/docs/search/list