// Code generated by MockGen. DO NOT EDIT.
// Source: app/usecase/user (interfaces: IRepositoryUser)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "app/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepositoryUser is a mock of IRepositoryUser interface.
type MockIRepositoryUser struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryUserMockRecorder
}

// MockIRepositoryUserMockRecorder is the mock recorder for MockIRepositoryUser.
type MockIRepositoryUserMockRecorder struct {
	mock *MockIRepositoryUser
}

// NewMockIRepositoryUser creates a new mock instance.
func NewMockIRepositoryUser(ctrl *gomock.Controller) *MockIRepositoryUser {
	mock := &MockIRepositoryUser{ctrl: ctrl}
	mock.recorder = &MockIRepositoryUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepositoryUser) EXPECT() *MockIRepositoryUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIRepositoryUser) CreateUser(arg0 *entity.EntityUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIRepositoryUserMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIRepositoryUser)(nil).CreateUser), arg0)
}

// DeleteUser mocks base method.
func (m *MockIRepositoryUser) DeleteUser(arg0 *entity.EntityUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIRepositoryUserMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIRepositoryUser)(nil).DeleteUser), arg0)
}

// GetByID mocks base method.
func (m *MockIRepositoryUser) GetByID(arg0 int) (*entity.EntityUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*entity.EntityUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIRepositoryUserMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIRepositoryUser)(nil).GetByID), arg0)
}

// GetByMail mocks base method.
func (m *MockIRepositoryUser) GetByMail(arg0 string) (*entity.EntityUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMail", arg0)
	ret0, _ := ret[0].(*entity.EntityUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMail indicates an expected call of GetByMail.
func (mr *MockIRepositoryUserMockRecorder) GetByMail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMail", reflect.TypeOf((*MockIRepositoryUser)(nil).GetByMail), arg0)
}

// GetUser mocks base method.
func (m *MockIRepositoryUser) GetUser(arg0 int) (*entity.EntityUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*entity.EntityUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIRepositoryUserMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIRepositoryUser)(nil).GetUser), arg0)
}

// GetUsers mocks base method.
func (m *MockIRepositoryUser) GetUsers() ([]entity.EntityUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]entity.EntityUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockIRepositoryUserMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockIRepositoryUser)(nil).GetUsers))
}

// UpdateUser mocks base method.
func (m *MockIRepositoryUser) UpdateUser(arg0 *entity.EntityUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIRepositoryUserMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIRepositoryUser)(nil).UpdateUser), arg0)
}
