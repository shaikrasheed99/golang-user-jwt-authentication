package helpers

import (
	"fmt"
	"net/http"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/responses"
)

func CreateSuccessResponse(code int, message string, data interface{}) responses.SuccessResponse {
	fmt.Println("[CreateSuccessResponse] Creating success response")

	return responses.SuccessResponse{
		Status:  constants.Success,
		Code:    http.StatusText(code),
		Message: message,
		Data:    data,
	}
}

func CreateErrorResponse(code int, message string) responses.ErrorResponse {
	fmt.Println("[CreateErrorResponse] Creating error response")

	return responses.ErrorResponse{
		Status:  constants.Error,
		Code:    http.StatusText(code),
		Message: message,
	}
}

func CreateUserResponse(user *models.User) responses.UserResponse {
	fmt.Println("[CreateUserResponse] Creating user response")

	return responses.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
	}
}

func CreateAuthenticationResponse(user *models.User, accessToken, refreshToken string) responses.AuthenticationResponse {
	fmt.Println("[CreateAuthenticationResponse] Creating authentication response")

	return responses.AuthenticationResponse{
		Username:     user.Username,
		Token:        accessToken,
		RefreshToken: refreshToken,
	}
}
