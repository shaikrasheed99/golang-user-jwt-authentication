package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/helpers"
	mocks "github.com/shaikrasheed99/golang-user-jwt-authentication/mocks/services"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/requests"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/responses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthHandler_SignupHandler(t *testing.T) {
	userServiceError := errors.New("error from user service")
	authServiceError := errors.New("error from auth service")
	userMock := &models.User{
		ID:        1,
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Username:  "test_username",
		Password:  "test_password",
		Role:      "test_role",
		Email:     "test_email@gmail.com",
	}
	signupReq := &requests.SignupRequest{
		FirstName: userMock.FirstName,
		LastName:  userMock.LastName,
		Username:  userMock.Username,
		Password:  userMock.Password,
		Email:     userMock.Email,
	}

	t.Run("should be able to sign up a user", func(t *testing.T) {
		mockUserService := new(mocks.UserService)
		mockAuthService := new(mocks.AuthService)
		mockUserService.On("Save", signupReq).Return(userMock, nil)
		mockAuthService.On("SaveTokensByUsername", userMock.Username, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
		at, rt, _ := helpers.GenerateToken(userMock.Username, userMock.Role)

		ah := NewAuthHandler(mockUserService, mockAuthService)
		router := gin.Default()
		router.POST(constants.SignupUserEndpoint, ah.SignupHandler)

		res := httptest.NewRecorder()
		body, _ := json.Marshal(signupReq)
		req, _ := http.NewRequest("POST", constants.SignupUserEndpoint, bytes.NewBuffer(body))
		router.ServeHTTP(res, req)

		var resBody responses.SuccessResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Success, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusOK), resBody.Code)
		assert.Equal(t, "successfully saved user details", resBody.Message)
		assert.Equal(t, nil, resBody.Data)
		assert.Equal(t, constants.AccessTokenCookie, res.Result().Cookies()[0].Name)
		assert.Equal(t, at, res.Result().Cookies()[0].Value)
		assert.Equal(t, constants.RefreshTokenCookie, res.Result().Cookies()[1].Name)
		assert.Equal(t, rt, res.Result().Cookies()[1].Value)
		mockAuthService.AssertExpectations(t)
		mockUserService.AssertExpectations(t)
	})

	t.Run("should not be able to sign up user when first name is not provided", func(t *testing.T) {
		invalidSignupReq := &requests.SignupRequest{
			LastName: userMock.LastName,
			Username: userMock.Username,
			Password: userMock.Password,
			Email:    userMock.Email,
		}
		mockUserService := new(mocks.UserService)
		mockAuthService := new(mocks.AuthService)

		ah := NewAuthHandler(mockUserService, mockAuthService)
		router := gin.Default()
		router.POST(constants.SignupUserEndpoint, ah.SignupHandler)

		res := httptest.NewRecorder()
		body, _ := json.Marshal(invalidSignupReq)
		req, _ := http.NewRequest("POST", constants.SignupUserEndpoint, bytes.NewBuffer(body))
		router.ServeHTTP(res, req)

		var resBody responses.SuccessResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Error, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusBadRequest), resBody.Code)
		assert.Equal(t, nil, resBody.Data)
		mockAuthService.AssertExpectations(t)
		mockUserService.AssertExpectations(t)
	})

	t.Run("should not be able to sign up user when there is error from user service", func(t *testing.T) {
		mockUserService := new(mocks.UserService)
		mockAuthService := new(mocks.AuthService)
		mockUserService.On("Save", signupReq).Return(nil, userServiceError)

		ah := NewAuthHandler(mockUserService, mockAuthService)
		router := gin.Default()
		router.POST(constants.SignupUserEndpoint, ah.SignupHandler)

		res := httptest.NewRecorder()
		body, _ := json.Marshal(signupReq)
		req, _ := http.NewRequest("POST", constants.SignupUserEndpoint, bytes.NewBuffer(body))
		router.ServeHTTP(res, req)

		var resBody responses.SuccessResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Error, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusInternalServerError), resBody.Code)
		assert.Equal(t, userServiceError.Error(), resBody.Message)
		assert.Equal(t, nil, resBody.Data)
		mockAuthService.AssertExpectations(t)
		mockUserService.AssertExpectations(t)
	})

	t.Run("should not be able to sign up user when there is error from auth service", func(t *testing.T) {
		mockUserService := new(mocks.UserService)
		mockAuthService := new(mocks.AuthService)
		mockUserService.On("Save", signupReq).Return(userMock, nil)
		mockAuthService.On("SaveTokensByUsername", userMock.Username, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(authServiceError)

		ah := NewAuthHandler(mockUserService, mockAuthService)
		router := gin.Default()
		router.POST(constants.SignupUserEndpoint, ah.SignupHandler)

		res := httptest.NewRecorder()
		body, _ := json.Marshal(signupReq)
		req, _ := http.NewRequest("POST", constants.SignupUserEndpoint, bytes.NewBuffer(body))
		router.ServeHTTP(res, req)

		var resBody responses.SuccessResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Error, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusInternalServerError), resBody.Code)
		assert.Equal(t, authServiceError.Error(), resBody.Message)
		assert.Equal(t, nil, resBody.Data)
		mockAuthService.AssertExpectations(t)
		mockUserService.AssertExpectations(t)
	})
}
