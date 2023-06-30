package main

import (
	"fmt"
	"github.com/raspi/youtubeapi"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv(`YT_APIKEY`)

	httpc := http.DefaultClient

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
}
