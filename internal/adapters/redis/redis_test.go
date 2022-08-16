package redis

import (
	"testing"

	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var (
	redisServer *miniredis.Miniredis
	redisClient *redis.Client
)

func TestUrlSave(t *testing.T) {
	setUp()
	defer teardown()
	logger := logging.GetLogger("debug")
	repo := NewUrlRepository(logger, redisClient)
	type mockCall func()
	testCases := []struct {
		title   string
		input   string
		isError bool
		mock    mockCall
	}{
		{
			title:   "should save successfully",
			input:   "url to save",
			isError: false,
			mock:    func() {},
		},
		{
			title:   "should doesn't save and return error ",
			input:   "url to save",
			isError: true,
			mock: func() {
				redisServer.SetError("interanl redis error")
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.mock()
			short, err := repo.Save(test.input)
			t.Log(short)
			if test.isError == true {
				assert.Error(t, err)
			} else {
				assert.Equal(t, err, nil)
			}

		})
	}
}

func TestGetUrl(t *testing.T) {
	setUp()
	defer teardown()
	logger := logging.GetLogger("debug")
	repo := NewUrlRepository(logger, redisClient)
	type mockCall func(string) (string, error)
	testCases := []struct {
		title   string
		input   string
		isError bool
		mock    mockCall
	}{
		{
			title:   "should save successfully",
			input:   "try to save it",
			isError: false,
			mock: func(input string) (string, error) {
				return repo.Save(input)
			},
		},
		{
			title:   "Get doens't find key and should return error",
			input:   "try to save it",
			isError: true,
			mock: func(input string) (string, error) {
				return "wrong key", nil
			},
		},
		{
			title:   "reddis internal error and Get return error ",
			input:   "try to save it",
			isError: true,
			mock: func(input string) (string, error) {
				redisServer.SetError("interanl redis error")
				return "good key", nil
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			key, _ := test.mock(test.input)
			url, err := repo.Get(key)
			if test.isError == false {
				assert.Equal(t, err, nil)
				assert.Equal(t, test.input, url)
			} else {
				assert.Error(t, err)
			}

		})
	}

}

func setUp() {
	redisServer = mockRedis()
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})
}

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}

func teardown() {
	redisServer.Close()
}
