# youtubeapi

![GitHub release (latest by date)](https://img.shields.io/github/v/release/raspi/youtubeapi?style=for-the-badge)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/raspi/youtubeapi?style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspi/youtubeapi)](https://goreportcard.com/report/github.com/raspi/youtubeapi)


YouTube REST API client for Go

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

fmt.Printf(`%v %v %v`, ch.PublishedAt, ch.Title, ch.Description)

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
