package services

import (
	"errors"
	"testing"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/helpers"
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
		assert.Equal(t, constants.ErrUserAlreadyExists, err.Error())
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

func TestUserService_Login(t *testing.T) {
	emptyUserMock := models.User{}
	dbError := errors.New("db error")
	loginRequest := &requests.LoginRequest{
		Username: "test_username",
		Password: "test_password",
	}
	hashedPassword, _ := helpers.GenerateHashedPassword(loginRequest.Password)
	userMock := models.User{
		Username: loginRequest.Username,
		Password: hashedPassword,
	}

	t.Run("should be able to login with valid user details", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindUserByUsername", loginRequest.Username).Return(userMock, nil)
		userService := NewUserService(mockUserRepo)

		user, err := userService.Login(loginRequest)

		assert.NoError(t, err)
		assert.Equal(t, user.FirstName, userMock.FirstName)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("should not be able to login when the user is not exists", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindUserByUsername", loginRequest.Username).Return(emptyUserMock, gorm.ErrRecordNotFound)
		userService := NewUserService(mockUserRepo)

		_, err := userService.Login(loginRequest)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrUserNotFound, err.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("should not be able to login when the user provides incorrect password", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindUserByUsername", loginRequest.Username).Return(userMock, nil)
		userService := NewUserService(mockUserRepo)
		loginRequest.Password = "abc"

		_, err := userService.Login(loginRequest)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrWrongPassword, err.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("should not be able to login when there is a error from database", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindUserByUsername", loginRequest.Username).Return(emptyUserMock, dbError)
		userService := NewUserService(mockUserRepo)

		_, err := userService.Login(loginRequest)

		assert.Error(t, err)
		assert.Equal(t, dbError.Error(), err.Error())
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_UserByUsername(t *testing.T) {
	emptyUserMock := models.User{}
	dbError := errors.New("db error")
	userMock := models.User{
		Username: "test_username",
		Email:    "test_email",
	}

	t.Run("should be able to get user details by username", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindUserByUsername", userMock.Username).Return(userMock, nil)
		userService := NewUserService(mockUserRepo)

		user, err := userService.UserByUsername(userMock.Username)

		assert.NoError(t, err)
		assert.Equal(t, user.Username, userMock.Username)
		assert.Equal(t, user.Email, userMock.Email)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("should be able to get empty user when user is not exists", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindUserByUsername", userMock.Username).Return(emptyUserMock, gorm.ErrRecordNotFound)
		userService := NewUserService(mockUserRepo)

		_, err := userService.UserByUsername(userMock.Username)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrUserNotFound, err.Error())
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("should be able tp get empty user when there is error from database", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindUserByUsername", userMock.Username).Return(emptyUserMock, dbError)
		userService := NewUserService(mockUserRepo)

		_, err := userService.UserByUsername(userMock.Username)

		assert.Error(t, err)
		assert.Equal(t, dbError.Error(), err.Error())
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_GetAllUsers(t *testing.T) {
	emptyUsersMock := []models.User{}
	dbError := errors.New("db error")
	usersMock := []models.User{
		{
			Username: "test_username",
			Email:    "test_email",
		},
	}

	t.Run("should be able to get all users list", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindAllUsers").Return(usersMock, nil)
		userService := NewUserService(mockUserRepo)

		users, err := userService.GetAllUsers()

		assert.NoError(t, err)
		assert.Equal(t, users[0].Username, usersMock[0].Username)
		assert.Equal(t, users[0].Email, usersMock[0].Email)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("should be able to get empty users list when there is error from database", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepository)
		mockUserRepo.On("FindAllUsers").Return(emptyUsersMock, dbError)
		userService := NewUserService(mockUserRepo)

		_, err := userService.GetAllUsers()

		assert.Error(t, err)
		assert.Equal(t, dbError.Error(), err.Error())
		mockUserRepo.AssertExpectations(t)
	})
}
