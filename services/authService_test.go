package services

import (
	"errors"
	"testing"

	mocks "github.com/shaikrasheed99/golang-user-jwt-authentication/mocks/repositories"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_SaveTokensByUsername(t *testing.T) {
	username := "test_username"
	at := "test_access_token"
	rt := "test_refresh_token"
	dbError := errors.New("db error")

	t.Run("should be able to save tokens by username", func(t *testing.T) {
		mockAuthRepo := new(mocks.AuthRepository)
		mockAuthRepo.On("SaveTokens", username, at, rt).Return(nil)
		authService := NewAuthService(mockAuthRepo)

		err := authService.SaveTokensByUsername(username, at, rt)

		assert.NoError(t, err)
		mockAuthRepo.AssertExpectations(t)
	})

	t.Run("should not be able to save tokens when there is error from database", func(t *testing.T) {
		mockAuthRepo := new(mocks.AuthRepository)
		mockAuthRepo.On("SaveTokens", username, at, rt).Return(dbError)
		authService := NewAuthService(mockAuthRepo)

		err := authService.SaveTokensByUsername(username, at, rt)

		assert.Error(t, err)
		assert.Equal(t, dbError.Error(), err.Error())
		mockAuthRepo.AssertExpectations(t)
	})
}
