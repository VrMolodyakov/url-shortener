// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/handler/serviceInterface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUrlService is a mock of UrlService interface.
type MockUrlService struct {
	ctrl     *gomock.Controller
	recorder *MockUrlServiceMockRecorder
}

// MockUrlServiceMockRecorder is the mock recorder for MockUrlService.
type MockUrlServiceMockRecorder struct {
	mock *MockUrlService
}

// NewMockUrlService creates a new mock instance.
func NewMockUrlService(ctrl *gomock.Controller) *MockUrlService {
	mock := &MockUrlService{ctrl: ctrl}
	mock.recorder = &MockUrlServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUrlService) EXPECT() *MockUrlServiceMockRecorder {
	return m.recorder
}

// CreateCustomUrl mocks base method.
func (m *MockUrlService) CreateCustomUrl(customUrl, url string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomUrl", customUrl, url)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCustomUrl indicates an expected call of CreateCustomUrl.
func (mr *MockUrlServiceMockRecorder) CreateCustomUrl(customUrl, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomUrl", reflect.TypeOf((*MockUrlService)(nil).CreateCustomUrl), customUrl, url)
}

// CreateUrl mocks base method.
func (m *MockUrlService) CreateUrl(url string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUrl", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUrl indicates an expected call of CreateUrl.
func (mr *MockUrlServiceMockRecorder) CreateUrl(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUrl", reflect.TypeOf((*MockUrlService)(nil).CreateUrl), url)
}

// GetUrl mocks base method.
func (m *MockUrlService) GetUrl(shortUrl string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl", shortUrl)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUrl indicates an expected call of GetUrl.
func (mr *MockUrlServiceMockRecorder) GetUrl(shortUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockUrlService)(nil).GetUrl), shortUrl)
}
