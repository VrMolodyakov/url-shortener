package redis

import (
	"math/rand"
	"strconv"

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

func (u *urlRepository) Save(url string) error {
	u.logger.Info("try to save %v", url)
	var id string
	for used := true; used; used = u.isExists(id) {
		id = strconv.FormatUint(rand.Uint64(), 10)
	}
	return u.client.Set(id, url, 0).Err()
}

func (u *urlRepository) Get(shortUrl string) (string, error) {
	u.logger.Info("try to get %v", shortUrl)
	return u.client.Get(shortUrl).Result()
}

func (u *urlRepository) isExists(id string) bool {
	return u.client.Exists(id).Val() != 0
}
