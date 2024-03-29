// Code generated by MockGen. DO NOT EDIT.
// Source: investment_score.go
//
// Generated by this command:
//
//	mockgen -source=investment_score.go -package=investment_score -destination=./mock/investment_score/investment_score.go
//

// Package investment_score is a generated GoMock package.
package investment_score

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockpartCalculator is a mock of partCalculator interface.
type MockpartCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockpartCalculatorMockRecorder
}

// MockpartCalculatorMockRecorder is the mock recorder for MockpartCalculator.
type MockpartCalculatorMockRecorder struct {
	mock *MockpartCalculator
}

// NewMockpartCalculator creates a new mock instance.
func NewMockpartCalculator(ctrl *gomock.Controller) *MockpartCalculator {
	mock := &MockpartCalculator{ctrl: ctrl}
	mock.recorder = &MockpartCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpartCalculator) EXPECT() *MockpartCalculatorMockRecorder {
	return m.recorder
}

// BAA10Y mocks base method.
func (m *MockpartCalculator) BAA10Y(ctx context.Context) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BAA10Y", ctx)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BAA10Y indicates an expected call of BAA10Y.
func (mr *MockpartCalculatorMockRecorder) BAA10Y(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BAA10Y", reflect.TypeOf((*MockpartCalculator)(nil).BAA10Y), ctx)
}

// FEDFUNDS mocks base method.
func (m *MockpartCalculator) FEDFUNDS(ctx context.Context) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FEDFUNDS", ctx)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FEDFUNDS indicates an expected call of FEDFUNDS.
func (mr *MockpartCalculatorMockRecorder) FEDFUNDS(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FEDFUNDS", reflect.TypeOf((*MockpartCalculator)(nil).FEDFUNDS), ctx)
}

// T10YFF mocks base method.
func (m *MockpartCalculator) T10YFF(ctx context.Context) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "T10YFF", ctx)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// T10YFF indicates an expected call of T10YFF.
func (mr *MockpartCalculatorMockRecorder) T10YFF(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "T10YFF", reflect.TypeOf((*MockpartCalculator)(nil).T10YFF), ctx)
}

// US10Y mocks base method.
func (m *MockpartCalculator) US10Y(ctx context.Context) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "US10Y", ctx)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// US10Y indicates an expected call of US10Y.
func (mr *MockpartCalculatorMockRecorder) US10Y(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "US10Y", reflect.TypeOf((*MockpartCalculator)(nil).US10Y), ctx)
}

// USDINDEX mocks base method.
func (m *MockpartCalculator) USDINDEX(ctx context.Context) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "USDINDEX", ctx)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// USDINDEX indicates an expected call of USDINDEX.
func (mr *MockpartCalculatorMockRecorder) USDINDEX(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "USDINDEX", reflect.TypeOf((*MockpartCalculator)(nil).USDINDEX), ctx)
}
