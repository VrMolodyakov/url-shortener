package handler

type UrlRequest struct {
	FullUrl string `json:"url"`
}

type UrlResponse struct {
	ShortUrl string `json:"shortUrl"`
	FullUrl  string `json:"fulltUrl"`
}
