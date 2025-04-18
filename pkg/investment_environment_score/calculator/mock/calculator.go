// Code generated by MockGen. DO NOT EDIT.
// Source: calculator.go
//
// Generated by this command:
//
//	mockgen -source=calculator.go -package=mock -destination=./mock/calculator.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	response "github.com/supermarine1377/yebis/pkg/fred/series/response"
	gomock "go.uber.org/mock/gomock"
)

// MockSeriesFetcher is a mock of SeriesFetcher interface.
type MockSeriesFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockSeriesFetcherMockRecorder
}

// MockSeriesFetcherMockRecorder is the mock recorder for MockSeriesFetcher.
type MockSeriesFetcherMockRecorder struct {
	mock *MockSeriesFetcher
}

// NewMockSeriesFetcher creates a new mock instance.
func NewMockSeriesFetcher(ctrl *gomock.Controller) *MockSeriesFetcher {
	mock := &MockSeriesFetcher{ctrl: ctrl}
	mock.recorder = &MockSeriesFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeriesFetcher) EXPECT() *MockSeriesFetcherMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockSeriesFetcher) Fetch(ctx context.Context, seriesID string, obeservationEnd time.Time) (*response.Res, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, seriesID, obeservationEnd)
	ret0, _ := ret[0].(*response.Res)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockSeriesFetcherMockRecorder) Fetch(ctx, seriesID, obeservationEnd any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockSeriesFetcher)(nil).Fetch), ctx, seriesID, obeservationEnd)
}
