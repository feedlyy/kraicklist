// Code generated by MockGen. DO NOT EDIT.
// Source: services/record_services.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	entity "challenge.haraj.com.sa/kraicklist/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockRecordServices is a mock of RecordServices interface.
type MockRecordServices struct {
	ctrl     *gomock.Controller
	recorder *MockRecordServicesMockRecorder
}

// MockRecordServicesMockRecorder is the mock recorder for MockRecordServices.
type MockRecordServicesMockRecorder struct {
	mock *MockRecordServices
}

// NewMockRecordServices creates a new mock instance.
func NewMockRecordServices(ctrl *gomock.Controller) *MockRecordServices {
	mock := &MockRecordServices{ctrl: ctrl}
	mock.recorder = &MockRecordServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecordServices) EXPECT() *MockRecordServicesMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *MockRecordServices) Search(query string) ([]entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockRecordServicesMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRecordServices)(nil).Search), query)
}
