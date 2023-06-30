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
	Maxres   *ThumbImage `json:"maxres,omitempty"`
	Standard *ThumbImage `json:"standard,omitempty"`
	High     *ThumbImage `json:"high,omitempty"`
	Medium   *ThumbImage `json:"medium,omitempty"`
	Default  *ThumbImage `json:"default,omitempty"`
}
