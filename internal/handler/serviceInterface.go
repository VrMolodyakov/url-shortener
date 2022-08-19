package handler

type UrlService interface {
	GetFullUrl(shortUrl string) (string, error)
	CreateShortUrl(url string) (string, error)
}
