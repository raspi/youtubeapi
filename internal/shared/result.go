package shared

// Common shared API JSON structs

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}

type ThumbImage struct {
	URL    string `json:"url"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}

type Thumbnails struct {
	Maxres   *ThumbImage `json:"maxres,omitempty"`   // 1280x720 https://i.ytimg.com/vi/<video id>/maxresdefault.jpg
	Standard *ThumbImage `json:"standard,omitempty"` // 640x480 https://i.ytimg.com/vi/<video id>/sddefault.jpg
	High     *ThumbImage `json:"high,omitempty"`     // 480x360 https://i.ytimg.com/vi/<video id>/hqdefault.jpg
	Medium   *ThumbImage `json:"medium,omitempty"`   // 320x180 https://i.ytimg.com/vi/<video id>/mqdefault.jpg
	Default  *ThumbImage `json:"default,omitempty"`  // 120x90 https://i.ytimg.com/vi/<video id>/default.jpg
}

// Meta is for custom parameter(s) so that next page(s) can be queried
type Meta struct {
	ETag          string  `json:"etag"`                      // For caching
	NextPageToken *string `json:"next_page_token,omitempty"` // Next page, nil if there's no next page, for custom parameters (`pageToken`)
}
