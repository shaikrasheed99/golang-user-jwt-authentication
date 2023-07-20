package helpers

import (
	"fmt"
	"net/http"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/responses"
)

func CreateSuccessResponse(code int, message string, data interface{}) responses.SuccessResponse {
	fmt.Println("[CreateSuccessResponseHelper] Creating success response")

	return responses.SuccessResponse{
		Status:  constants.Success,
		Code:    http.StatusText(code),
		Message: message,
		Data:    data,
	}
}

func CreateErrorResponse(code int, message string) responses.ErrorResponse {
	fmt.Println("[CreateErrorResponseHelper] Creating error response")

	return responses.ErrorResponse{
		Status:  constants.Error,
		Code:    http.StatusText(code),
		Message: message,
	}
}

func CreateUserResponse(user *models.User) responses.UserResponse {
	fmt.Println("[CreateUserResponseHelper] Creating user response")

	return responses.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
	}
}
