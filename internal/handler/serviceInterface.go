package handler

type UrlService interface {
	GetUrl(shortUrl string) (string, error)
	CreateUrl(url string) (string, error)
	CreateCustomUrl(customUrl string, url string) (string, error)
}
