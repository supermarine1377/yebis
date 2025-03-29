// Code generated by MockGen. DO NOT EDIT.
// Source: common.go
//
// Generated by this command:
//
//	mockgen -source=common.go -package=mock -destination=./mock/common.go
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockConfig is a mock of Config interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig.
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// FEDAPIKEY mocks base method.
func (m *MockConfig) FEDAPIKEY() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FEDAPIKEY")
	ret0, _ := ret[0].(string)
	return ret0
}

// FEDAPIKEY indicates an expected call of FEDAPIKEY.
func (mr *MockConfigMockRecorder) FEDAPIKEY() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FEDAPIKEY", reflect.TypeOf((*MockConfig)(nil).FEDAPIKEY))
}
