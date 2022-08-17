package handler

type UrlRequest struct {
	Url string `json:"url"`
}

type UrlResponse struct {
	ShortUrl string `json:"shortUrl"`
}
