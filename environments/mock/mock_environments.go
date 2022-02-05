// Code generated by MockGen. DO NOT EDIT.
// Source: environments.go

// Package mock_environments is a generated GoMock package.
package mock_environments

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEnvironments is a mock of Environments interface.
type MockEnvironments struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentsMockRecorder
}

// MockEnvironmentsMockRecorder is the mock recorder for MockEnvironments.
type MockEnvironmentsMockRecorder struct {
	mock *MockEnvironments
}

// NewMockEnvironments creates a new mock instance.
func NewMockEnvironments(ctrl *gomock.Controller) *MockEnvironments {
	mock := &MockEnvironments{ctrl: ctrl}
	mock.recorder = &MockEnvironmentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvironments) EXPECT() *MockEnvironmentsMockRecorder {
	return m.recorder
}

// GetInvalidWordsSearchQueryString mocks base method.
func (m *MockEnvironments) GetInvalidWordsSearchQueryString() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvalidWordsSearchQueryString")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetInvalidWordsSearchQueryString indicates an expected call of GetInvalidWordsSearchQueryString.
func (mr *MockEnvironmentsMockRecorder) GetInvalidWordsSearchQueryString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvalidWordsSearchQueryString", reflect.TypeOf((*MockEnvironments)(nil).GetInvalidWordsSearchQueryString))
}
