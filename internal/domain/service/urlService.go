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

func NewUrlRepository(repo UrlRepository, shortener Shortener) *urlService {
	return &urlService{repo: repo, shortener: shortener}
}

func (u *urlService) GetFullUrl(shortUrl string) (string, error) {
	fullUrl, err := u.repo.Get(shortUrl)
	if err != nil {
		return "", errs.ErrUrlNotFound
	}
	return fullUrl, nil
}

func (u *urlService) CreateShortUrl(url string) (string, error) {
	var id uint64
	for used := true; used; used = u.repo.IsExists(id) {
		id = rand.Uint64()
	}
	err := u.repo.Save(strconv.FormatUint(id, 10), url)
	if err != nil {
		return "", errs.ErrUrlNotSaved
	}
	return u.shortener.Encode(id), err
}
