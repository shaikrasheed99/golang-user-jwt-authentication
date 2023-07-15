package helpers

import (
	"fmt"
	"net/http"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/responses"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/utils"
)

func CreateUserResponse(user *models.User) responses.UserResponse {
	fmt.Println("[createUserResponse] Creating user response")

	return responses.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}
}

func CreateSuccessResponse(code int, message string, data interface{}) responses.SuccessResponse {
	fmt.Println("[createSuccessResponse] Creating success response")

	return responses.SuccessResponse{
		Status:  utils.Success,
		Code:    http.StatusText(code),
		Message: message,
		Data:    data,
	}
}

func CreateErrorResponse(code int, message string) responses.ErrorResponse {
	fmt.Println("[createErrorResponse] Creating error response")

	return responses.ErrorResponse{
		Status:  utils.Error,
		Code:    http.StatusText(code),
		Message: message,
	}
}
