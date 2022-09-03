package handler

import (
	"bytes"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/VrMolodyakov/url-shortener/internal/errs"
	"github.com/VrMolodyakov/url-shortener/internal/handler/mocks"
	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestEncodeUrl(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	urlService := mocks.NewMockUrlService(ctrl)
	handler := NewUrlHandler(logging.GetLogger("debug"), urlService)
	handler.InitRoutes(router)
	type mockCall func()
	testCases := []struct {
		title          string
		inputRequest   string
		want           string
		mock           mockCall
		expectedStatus int
	}{
		{
			title:        "create short url and 201 response",
			inputRequest: `{"url":"random.test.url"}`,
			mock: func() {
				urlService.EXPECT().CreateUrl(gomock.Any()).Return("random.short.url", nil)
			},
			want:           `{"shortUrl": "random.short.url","fulltUrl": "random.test.url"}`,
			expectedStatus: 201,
		},
		{
			title:        "url notfound and 400 response",
			inputRequest: `{"url":"random.test.url"}`,
			mock: func() {
				urlService.EXPECT().CreateUrl(gomock.Any()).Return("", errs.ErrUrlNotFound)
			},
			want:           `Short url not found`,
			expectedStatus: 500,
		},
		{
			title:        "url is empty and 400 response",
			inputRequest: `{"url":""}`,
			mock: func() {
				urlService.EXPECT().CreateUrl(gomock.Any()).Return("", errs.ErrUrlIsEmpty)
			},
			want:           `url is empty`,
			expectedStatus: 400,
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.mock()
			req := httptest.NewRequest(
				"POST",
				"/encode",
				bytes.NewBufferString(test.inputRequest),
			)
			req.Header.Set("Content-type", "application/json")
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)
			assert.Equal(t, clearResponse(recorder.Body.String()), test.want)
			assert.Equal(t, test.expectedStatus, recorder.Code)
		})
	}
}

func TestEncodeCustomUrl(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	urlService := mocks.NewMockUrlService(ctrl)
	handler := NewUrlHandler(logging.GetLogger("debug"), urlService)
	handler.InitRoutes(router)
	type mockCall func()
	testCases := []struct {
		title          string
		inputRequest   string
		want           string
		mock           mockCall
		expectedStatus int
	}{
		{
			title:        "create custom short url and 201 response",
			inputRequest: `{"url":"random.test.url","custom":"custom.url"}`,
			mock: func() {
				urlService.EXPECT().CreateCustomUrl(gomock.Any(), gomock.Any()).Return(nil)
			},
			want:           `{"shortUrl": "custom.url","fulltUrl": "random.test.url"}`,
			expectedStatus: 201,
		},
		{
			title:        "url notfound and 400 response",
			inputRequest: `{"url":"random.test.url","custom":"custom.url"}`,
			mock: func() {
				urlService.EXPECT().CreateCustomUrl(gomock.Any(), gomock.Any()).Return(errs.ErrUrlNotFound)
			},
			want:           `Short url not found`,
			expectedStatus: 500,
		},
		{
			title:        "url is empty and 400 response",
			inputRequest: `{"url":""}`,
			mock: func() {
				urlService.EXPECT().CreateCustomUrl(gomock.Any(), gomock.Any()).Return(errs.ErrUrlIsEmpty)
			},
			want:           `url is empty`,
			expectedStatus: 400,
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.mock()
			req := httptest.NewRequest(
				"POST",
				"/encode/custom",
				bytes.NewBufferString(test.inputRequest),
			)
			req.Header.Set("Content-type", "application/json")
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)
			assert.Equal(t, clearResponse(recorder.Body.String()), test.want)
			assert.Equal(t, test.expectedStatus, recorder.Code)
		})
	}
}

func TestDecodeUrl(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	urlService := mocks.NewMockUrlService(ctrl)
	handler := NewUrlHandler(logging.GetLogger("debug"), urlService)
	handler.InitRoutes(router)
	type mockCall func()
	testCases := []struct {
		title          string
		want           string
		mock           mockCall
		expectedStatus int
		isError        bool
	}{
		{
			title: "get full url and 200 response",
			mock: func() {
				urlService.EXPECT().GetUrl(gomock.Any()).Return("full.url", nil)
			},
			want:           `/full.url`,
			expectedStatus: 303,
			isError:        false,
		},
		{
			title: "url notfound and 400 response",
			mock: func() {
				urlService.EXPECT().GetUrl(gomock.Any()).Return("", errs.ErrUrlNotFound)
			},
			want:           `Short url not found`,
			expectedStatus: 500,
			isError:        true,
		},
		{
			title: "url is empty and 400 response",
			mock: func() {
				urlService.EXPECT().GetUrl(gomock.Any()).Return("", errs.ErrUrlIsEmpty)
			},
			want:           `url is empty`,
			expectedStatus: 400,
			isError:        true,
		},
	}
	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.mock()
			req := httptest.NewRequest(
				"GET",
				"/shortUrl",
				nil)
			req.Header.Set("Content-type", "application/json")
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)
			if !test.isError {
				url, _ := recorder.Result().Location()
				assert.Equal(t, test.want, url.String())
				assert.Equal(t, test.expectedStatus, recorder.Code)
			} else {
				assert.Equal(t, clearResponse(recorder.Body.String()), test.want)
				assert.Equal(t, test.expectedStatus, recorder.Code)
			}

		})
	}
}

func clearResponse(s string) string {
	temp := strings.ReplaceAll(s, "   ", "")
	return strings.ReplaceAll(temp, "\n", "")
}
