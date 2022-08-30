package urlDb

import (
	"strconv"

	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/go-redis/redis"
)

const (
	short string = "short"
	full  string = "long"
)

type urlRepository struct {
	logger *logging.Logger
	client *redis.Client
}

func NewUrlRepository(logger *logging.Logger, client *redis.Client) *urlRepository {
	return &urlRepository{logger: logger, client: client}
}

func (u *urlRepository) Save(id string, url string) error {
	err := u.client.Set(id, url, 0).Err()
	if err != nil {
		u.logger.Errorf("cannot save id = %v wih url = %v due to ", err)
		return err
	}
	return nil
}

func (u *urlRepository) Get(key string) (string, error) {
	u.logger.Info("try to get %v", key)
	url, err := u.client.Get(key).Result()
	if err != nil {
		u.logger.Errorf("cannot get full url for short url = %v due to ", err)
		return "", err
	}
	return url, err
}

func (u *urlRepository) IsExists(id uint64) bool {
	return u.client.Exists(strconv.FormatUint(id, 10)).Val() != 0
}
