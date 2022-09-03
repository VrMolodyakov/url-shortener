package service

import (
	"errors"
	"testing"

	"github.com/VrMolodyakov/url-shortener/internal/domain/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUrl(t *testing.T) {
	cntr := gomock.NewController(t)
	repo := mocks.NewMockUrlRepository(cntr)
	shortener := mocks.NewMockShortener(cntr)
	urlService := NewUrlService(repo, shortener)
	type mockCall func()
	testCases := []struct {
		title   string
		input   string
		mock    mockCall
		want    string
		isError bool
	}{
		{
			title: "Success short url creating",
			input: "https://some-full-url",
			mock: func() {
				repo.EXPECT().IsExists(gomock.Any()).Return(false)
				repo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
				shortener.EXPECT().Encode(gomock.Any()).Return("short url")
			},
			want:    "short url",
			isError: false,
		},
		{
			title: "Cannot Save and should return error",
			input: "https://some-full-url",
			mock: func() {
				repo.EXPECT().IsExists(gomock.Any()).Return(false)
				repo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("service error"))
			},
			want:    "short url",
			isError: true,
		},
		{
			title: "Empty url and Create should return error",
			input: "",
			mock: func() {
			},
			want:    "",
			isError: true,
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.mock()
			got, err := urlService.CreateUrl(test.input)
			if !test.isError {
				assert.NoError(t, err)
				assert.Equal(t, got, test.want)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestCreateCustomUrl(t *testing.T) {
	cntr := gomock.NewController(t)
	repo := mocks.NewMockUrlRepository(cntr)
	shortener := mocks.NewMockShortener(cntr)
	urlService := NewUrlService(repo, shortener)
	type mockCall func()
	type args struct {
		customUrl string
		url       string
	}
	testCases := []struct {
		title   string
		input   args
		mock    mockCall
		want    string
		isError bool
	}{
		{
			title: "Success custom short url creating",
			input: args{customUrl: "custom url", url: "full url"},
			mock: func() {
				var randomId uint64 = 424242
				shortener.EXPECT().Decode(gomock.Any()).Return(randomId)
				repo.EXPECT().IsExists(gomock.Any()).Return(false)
				repo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
			want:    "custom url",
			isError: false,
		},
		{
			title: "url is empty and should return error",
			input: args{customUrl: "custom url", url: ""},
			mock: func() {
			},
			want:    "custom url",
			isError: true,
		},
		{
			title: "custom url is empty and should return error",
			input: args{customUrl: "", url: "full url"},
			mock: func() {
			},
			want:    "custom url",
			isError: true,
		},
		{
			title: "custom url is already exist and should return error",
			input: args{customUrl: "custom url", url: "full url"},
			mock: func() {
				var randomId uint64 = 42424242
				shortener.EXPECT().Decode(gomock.Any()).Return(randomId)
				repo.EXPECT().IsExists(gomock.Any()).Return(true)
			},
			want:    "custom url",
			isError: true,
		},
		{
			title: "cannot save url and should return error",
			input: args{customUrl: "custom url", url: "full url"},
			mock: func() {
				var randomId uint64 = 42424242
				shortener.EXPECT().Decode(gomock.Any()).Return(randomId)
				repo.EXPECT().IsExists(gomock.Any()).Return(false)
				repo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("service error"))

			},
			want:    "custom url",
			isError: true,
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.mock()
			err := urlService.CreateCustomUrl(test.input.customUrl, test.input.url)
			if !test.isError {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
func TestGetUrl(t *testing.T) {
	cntr := gomock.NewController(t)
	repo := mocks.NewMockUrlRepository(cntr)
	shortener := mocks.NewMockShortener(cntr)
	urlService := NewUrlService(repo, shortener)
	type mockCall func()
	testCases := []struct {
		title   string
		input   string
		mock    mockCall
		want    string
		isError bool
	}{
		{
			title: "successful url receiving",
			input: "https://some-short-url",
			mock: func() {
				var randomId uint64 = 42424242
				shortener.EXPECT().Decode(gomock.Any()).Return(randomId)
				repo.EXPECT().Get(gomock.Any()).Return("https://some-full-url", nil)

			},
			want:    "https://some-full-url",
			isError: false,
		},
		{
			title: "short url is empty and should return error",
			input: "",
			mock: func() {
			},
			want:    "https://some-full-url",
			isError: true,
		},
		{
			title: "couldn't get url and should return error",
			input: "https://some-short-url",
			mock: func() {
				var randomId uint64 = 42424242
				shortener.EXPECT().Decode(gomock.Any()).Return(randomId)
				repo.EXPECT().Get(gomock.Any()).Return("", errors.New("internal repo error"))
			},
			want:    "",
			isError: true,
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.mock()
			got, err := urlService.GetUrl(test.input)
			if !test.isError {
				assert.NoError(t, err)
				assert.Equal(t, got, test.want)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
