package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/requests"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/responses"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/services"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/utils"
)

type UserController interface {
	SignupHandler(*gin.Context)
	LoginHandler(*gin.Context)
	Health(*gin.Context)
}

type userController struct {
	us services.UserService
}

func NewUserController(us services.UserService) UserController {
	return &userController{
		us: us,
	}
}

func (uc *userController) SignupHandler(c *gin.Context) {
	var req *requests.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("[SignupHandler]", err.Error())
		errRes := createErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	savedUser, err := uc.us.Save(req)
	if err != nil {
		fmt.Println("[SignupHandler]", err.Error())
		errRes := createErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	savedUserRes := responses.SavedUserResponse{
		ID:        savedUser.ID,
		FirstName: savedUser.FirstName,
		LastName:  savedUser.LastName,
		Username:  savedUser.Username,
		Email:     savedUser.Email,
	}

	res := createSuccessResponse(http.StatusOK, "successfully saved user details", savedUserRes)

	c.JSON(http.StatusCreated, res)
}

func (uc *userController) LoginHandler(c *gin.Context) {
	var req *requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("[LoginHandler]", err.Error())
		errRes := createErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user, err := uc.us.Login(req)
	if err != nil {
		fmt.Println("[LoginHandler]", err.Error())
		errRes := createErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	userRes := responses.SavedUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}

	res := createSuccessResponse(http.StatusOK, "successfully logged in", userRes)

	c.JSON(http.StatusCreated, res)

}

func createSuccessResponse(code int, message string, data interface{}) responses.SuccessResponse {
	res := responses.SuccessResponse{
		Status:  utils.Success,
		Code:    http.StatusText(code),
		Message: message,
		Data:    data,
	}

	return res
}

func createErrorResponse(code int, message string) responses.ErrorResponse {
	res := responses.ErrorResponse{
		Status:  utils.Error,
		Code:    http.StatusText(code),
		Message: message,
	}

	return res
}

func (uc *userController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UP!",
	})
}
