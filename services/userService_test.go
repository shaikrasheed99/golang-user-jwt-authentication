package services

import (
	"errors"
	"testing"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	mocks "github.com/shaikrasheed99/golang-user-jwt-authentication/mocks/repositories"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/requests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestUserService_Save(t *testing.T) {
	emptyUserMock := models.User{}
	dbError := errors.New("db error")
	userMock := models.User{
		ID:        1,
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Username:  "test_username",
		Password:  "test_password",
		Role:      "test_role",
		Email:     "test_email",
	}
	signupRequest := &requests.SignupRequest{
		FirstName: userMock.FirstName,
		LastName:  userMock.LastName,
		Username:  userMock.Username,
		Password:  userMock.Password,
		Email:     userMock.Email,
	}

	t.Run("should be able to save user details", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindUserByUsername", signupRequest.Username).Return(emptyUserMock, gorm.ErrRecordNotFound)
		mockUserRepo.On("Save", mock.AnythingOfType("*models.User")).Return(userMock, nil)
		userService := NewUserService(mockUserRepo)

		savedUser, err := userService.Save(signupRequest)

		assert.NoError(t, err)
		assert.Equal(t, savedUser.FirstName, userMock.FirstName)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("should not be able to save user when details already exists", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindUserByUsername", signupRequest.Username).Return(userMock, nil)
		userService := NewUserService(mockUserRepo)

		_, err := userService.Save(signupRequest)

		assert.Error(t, err)
		assert.Equal(t, constants.UserAlreadyExistsErrorMessage, err.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("should not be able to save user when there is a error from database", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindUserByUsername", signupRequest.Username).Return(emptyUserMock, gorm.ErrRecordNotFound)
		mockUserRepo.On("Save", mock.AnythingOfType("*models.User")).Return(emptyUserMock, dbError)
		userService := NewUserService(mockUserRepo)

		_, err := userService.Save(signupRequest)

		assert.Error(t, err)
		assert.Equal(t, dbError.Error(), err.Error())
		mockUserRepo.AssertExpectations(t)
	})
}
