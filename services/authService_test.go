package services

import (
	"errors"
	"testing"
	"time"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	mocks "github.com/shaikrasheed99/golang-user-jwt-authentication/mocks/repositories"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

func TestAuthService_FindTokensByUsername(t *testing.T) {
	dbError := errors.New("db error")
	emptyTokenMock := models.Tokens{}
	username := "test_username"

	tokenMock := models.Tokens{
		Username:     username,
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
		CreatedAt:    time.Now(),
	}

	t.Run("should be able to find tokens by username", func(t *testing.T) {
		mockAuthRepo := new(mocks.AuthRepository)
		mockAuthRepo.On("FindTokensByUsername", username).Return(tokenMock, nil)
		authService := NewAuthService(mockAuthRepo)

		tokens, err := authService.FindTokensByUsername(username)

		assert.NoError(t, err)
		assert.Equal(t, tokens.Username, tokenMock.Username)
		assert.Equal(t, tokens.AccessToken, tokenMock.AccessToken)
		assert.Equal(t, tokens.RefreshToken, tokenMock.RefreshToken)
		assert.Equal(t, tokens.CreatedAt, tokenMock.CreatedAt)
		mockAuthRepo.AssertExpectations(t)
	})

	t.Run("should be able to get empty tokens when tokens are not exists", func(t *testing.T) {
		mockAuthRepo := new(mocks.AuthRepository)
		mockAuthRepo.On("FindTokensByUsername", username).Return(emptyTokenMock, gorm.ErrRecordNotFound)
		authService := NewAuthService(mockAuthRepo)

		_, err := authService.FindTokensByUsername(username)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrTokensNotFound, err.Error())
		mockAuthRepo.AssertExpectations(t)
	})

	t.Run("should be able to get empty tokens when there is error from database", func(t *testing.T) {
		mockAuthRepo := new(mocks.AuthRepository)
		mockAuthRepo.On("FindTokensByUsername", username).Return(emptyTokenMock, dbError)
		authService := NewAuthService(mockAuthRepo)

		_, err := authService.FindTokensByUsername(username)

		assert.Error(t, err)
		assert.Equal(t, dbError.Error(), err.Error())
		mockAuthRepo.AssertExpectations(t)
	})
}

func TestAuthService_DeleteTokensByUsername(t *testing.T) {
	username := "test_username"
	dbError := errors.New("db error")

	t.Run("should be able to delete tokens by username", func(t *testing.T) {
		mockAuthRepo := new(mocks.AuthRepository)
		mockAuthRepo.On("DeleteTokensByUsername", username).Return(nil)
		authService := NewAuthService(mockAuthRepo)

		err := authService.DeleteTokensByUsername(username)

		assert.NoError(t, err)
		mockAuthRepo.AssertExpectations(t)
	})

	t.Run("should not be able to delete tokens when there is error from database", func(t *testing.T) {
		mockAuthRepo := new(mocks.AuthRepository)
		mockAuthRepo.On("DeleteTokensByUsername", username).Return(dbError)
		authService := NewAuthService(mockAuthRepo)

		err := authService.DeleteTokensByUsername(username)

		assert.Error(t, err)
		assert.Equal(t, dbError.Error(), err.Error())
		mockAuthRepo.AssertExpectations(t)
	})
}
