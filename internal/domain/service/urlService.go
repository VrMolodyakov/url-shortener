package service

import (
	"math/rand"
	"strconv"

	"github.com/VrMolodyakov/url-shortener/internal/errs"
)

type UrlRepository interface {
	Get(shortUrl string) (string, error)
	Save(id string, url string) error
	IsExists(id uint64) bool
}

type Shortener interface {
	Encode(num uint64) string
	Decode(token string) uint64
}

type urlService struct {
	repo      UrlRepository
	shortener Shortener
}

func NewUrlService(repo UrlRepository, shortener Shortener) *urlService {
	return &urlService{repo: repo, shortener: shortener}
}

func (u *urlService) GetUrl(shortUrl string) (string, error) {
	if shortUrl == "" {
		return shortUrl, errs.ErrUrlIsEmpty
	}
	id := u.shortener.Decode(shortUrl)
	key := strconv.FormatUint(id, 10)
	fullUrl, err := u.repo.Get(key)
	if err != nil {
		return "", errs.ErrUrlNotFound
	}
	return fullUrl, nil
}

func (u *urlService) CreateUrl(url string) (string, error) {
	if url == "" {
		return url, errs.ErrUrlIsEmpty
	}
	var id uint64
	for used := true; used; used = u.repo.IsExists(id) {
		id = rand.Uint64()
	}
	err := u.repo.Save(strconv.FormatUint(id, 10), url)
	if err != nil {
		return "", errs.ErrUrlNotSaved
	}
	return u.shortener.Encode(id), nil
}

func (u *urlService) CreateCustomUrl(customUrl string, url string) error {
	if url == "" || customUrl == "" {
		return errs.ErrUrlIsEmpty
	}
	id := u.shortener.Decode(customUrl)
	if u.repo.IsExists(id) {
		return errs.ErrUrlNotSaved
	}
	err := u.repo.Save(strconv.FormatUint(id, 10), url)
	if err != nil {
		return errs.ErrUrlNotSaved
	}
	return nil

}
