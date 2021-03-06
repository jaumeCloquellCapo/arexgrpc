// Code generated by MockGen. DO NOT EDIT.
// Source: AuthRepository.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/jaumeCloquellCapo/authGrpc/app/model"
	reflect "reflect"
)

/**
CreateToken(user model.User) (td model.TokenDetails, err error)
CreateAuth(user model.User, td model.TokenDetails) error
GetAuth(AccessUUID string) (int64, error)
DeleteAuth(AccessUUID string) error
*/

// MockAuthRepository is a mock of UserRedisRepository interface
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

func (m *MockAuthRepository) CreateAuth(user model.User, td model.TokenDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuth", user, td)
	ret0, _ := ret[0].(error)
	return ret0
}
func (mr *MockAuthRepositoryMockRecorder) CreateAuth(user interface{}, td interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuth", reflect.TypeOf((*MockAuthRepository)(nil).CreateToken), user, td)
}

// GetByIDCtx mocks base method
func (m *MockAuthRepository) CreateToken(user model.User) (td model.TokenDetails, err error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", user)
	ret0, _ := ret[0].(model.TokenDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDCtx indicates an expected call of GetByIDCtx
func (mr *MockAuthRepositoryMockRecorder) CreateToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockAuthRepository)(nil).CreateToken), user)
}

// SetUserCtx mocks base method
func (m *MockAuthRepository) GetAuth(AccessUUID string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuth", AccessUUID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUserCtx indicates an expected call of SetUserCtx
func (mr *MockAuthRepositoryMockRecorder) GetAuth(AccessUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuth", reflect.TypeOf((*MockAuthRepository)(nil).GetAuth), AccessUUID)
}

// DeleteUserCtx mocks base method
func (m *MockAuthRepository) DeleteAuth(AccessUUID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuth", AccessUUID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserCtx indicates an expected call of DeleteUserCtx
func (mr *MockAuthRepositoryMockRecorder) DeleteAuth(AccessUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuth", reflect.TypeOf((*MockAuthRepository)(nil).DeleteAuth), AccessUUID)
}
