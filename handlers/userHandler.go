package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/requests"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/responses"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/services"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/utils"
)

type UserHandler interface {
	SignupHandler(*gin.Context)
	LoginHandler(*gin.Context)
	UserByUsernameHandler(*gin.Context)
	Health(*gin.Context)
}

type userHandler struct {
	us services.UserService
}

func NewUserHandler(us services.UserService) UserHandler {
	fmt.Println("[NewUserHandler] Initiating New User Handler")
	return &userHandler{
		us: us,
	}
}

func (uc *userHandler) SignupHandler(c *gin.Context) {
	fmt.Println("[SignupHandler] Hitting signup handler function in user handler")

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

func (uc *userHandler) LoginHandler(c *gin.Context) {
	fmt.Println("[LoginHandler] Hitting login handler function in user handler")

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

func (uc *userHandler) UserByUsernameHandler(c *gin.Context) {
	fmt.Println("[UserByUsernameHandler] Hitting user by username handler function in user handler")

	username := c.Param("username")
	_, err := strconv.Atoi(username)
	if err == nil || username == "" {
		fmt.Println("[UserByUsernameHandler]", err.Error())
		errRes := createErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user, err := uc.us.UserByUsername(username)
	if err != nil {
		fmt.Println("[UserByUsernameHandler]", err.Error())
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

	res := createSuccessResponse(http.StatusOK, "successfully got user details", userRes)

	c.JSON(http.StatusCreated, res)
}

func createSuccessResponse(code int, message string, data interface{}) responses.SuccessResponse {
	fmt.Println("[createSuccessResponse] Creating success response")

	res := responses.SuccessResponse{
		Status:  utils.Success,
		Code:    http.StatusText(code),
		Message: message,
		Data:    data,
	}

	return res
}

func createErrorResponse(code int, message string) responses.ErrorResponse {
	fmt.Println("[createErrorResponse] Creating error response")

	res := responses.ErrorResponse{
		Status:  utils.Error,
		Code:    http.StatusText(code),
		Message: message,
	}

	return res
}

func (uc *userHandler) Health(c *gin.Context) {
	fmt.Println("[Health] Hitting health function in user handler")

	c.JSON(http.StatusOK, gin.H{
		"message": "UP!",
	})
}
