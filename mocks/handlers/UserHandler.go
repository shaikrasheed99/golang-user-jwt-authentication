// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// UserHandler is an autogenerated mock type for the UserHandler type
type UserHandler struct {
	mock.Mock
}

// GetAllUsers provides a mock function with given fields: _a0
func (_m *UserHandler) GetAllUsers(_a0 *gin.Context) {
	_m.Called(_a0)
}

// UserByUsernameHandler provides a mock function with given fields: _a0
func (_m *UserHandler) UserByUsernameHandler(_a0 *gin.Context) {
	_m.Called(_a0)
}

// NewUserHandler creates a new instance of UserHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserHandler {
	mock := &UserHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
