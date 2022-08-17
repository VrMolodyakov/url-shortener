package service

import "github.com/VrMolodyakov/url-shortener/internal/errs"

type UrlRepository interface {
	Get(shortUrl string) (string, error)
	Save(url string) (string, error)
	IsExists(id uint64) bool
}

type urlService struct {
	repo UrlRepository
}

func NewUrlRepository(repo UrlRepository) *urlService {
	return &urlService{repo: repo}
}

func (u *urlService) FindFullUrl(shortUrl string) (string, error) {
	fullUrl, err := u.repo.Get(shortUrl)
	if err != nil {
		return "", errs.ErrUrlNotFound
	}
	return fullUrl, nil
}

func (u *urlService) CreateShortUrl(url string) (string, error) {
	shortUrl, err := u.repo.Save(url)
	if err != nil {
		return "", errs.ErrUrlNotSaved
	}
	return shortUrl, nil
}
