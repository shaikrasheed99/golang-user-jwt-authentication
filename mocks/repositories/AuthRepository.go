// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	models "github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	mock "github.com/stretchr/testify/mock"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// DeleteTokensByUsername provides a mock function with given fields: _a0
func (_m *AuthRepository) DeleteTokensByUsername(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindTokensByUsername provides a mock function with given fields: _a0
func (_m *AuthRepository) FindTokensByUsername(_a0 string) (models.Tokens, error) {
	ret := _m.Called(_a0)

	var r0 models.Tokens
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Tokens, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) models.Tokens); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Tokens)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveTokens provides a mock function with given fields: _a0, _a1, _a2
func (_m *AuthRepository) SaveTokens(_a0 string, _a1 string, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAuthRepository creates a new instance of AuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthRepository {
	mock := &AuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
