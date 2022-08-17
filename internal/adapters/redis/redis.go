package redis

import (
	"math/rand"
	"strconv"

	"github.com/VrMolodyakov/url-shortener/internal/service/shortener"
	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/go-redis/redis"
)

type urlRepository struct {
	logger *logging.Logger
	client *redis.Client
}

func NewUrlRepository(logger *logging.Logger, client *redis.Client) *urlRepository {
	return &urlRepository{logger: logger, client: client}
}

func (u *urlRepository) Save(url string) (string, error) {
	u.logger.Infof("try to save %v", url)
	var id uint64
	for used := true; used; used = u.IsExists(id) {
		id = rand.Uint64()
	}
	err := u.client.Set(strconv.FormatUint(id, 10), url, 0).Err()
	return shortener.Encode(id), err
}

func (u *urlRepository) Get(shortUrl string) (string, error) {
	id := shortener.Decode(shortUrl)
	u.logger.Info("try to get %v", shortUrl)
	return u.client.Get(strconv.FormatUint(id, 10)).Result()
}

func (u *urlRepository) IsExists(id uint64) bool {
	return u.client.Exists(strconv.FormatUint(id, 10)).Val() != 0
}
