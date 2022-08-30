package handler

type UrlRequest struct {
	Url string `json:"url"`
}

type UrlResponse struct {
	ShortUrl string `json:"shortUrl"`
	FullUrl  string `json:"fulltUrl"`
}

type CustomUrlRequest struct {
	Url    string `json:"url"`
	Custom string `json:"custom"`
}
