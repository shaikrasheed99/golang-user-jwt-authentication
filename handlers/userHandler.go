package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/requests"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/responses"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/services"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/utils"
)

type UserHandler interface {
	SignupHandler(*gin.Context)
	LoginHandler(*gin.Context)
	UserByUsernameHandler(*gin.Context)
	GetAllUsers(*gin.Context)
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

func (uh *userHandler) SignupHandler(c *gin.Context) {
	fmt.Println("[SignupHandler] Hitting signup handler function in user handler")

	var req *requests.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("[SignupHandler]", err.Error())
		errRes := createErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	savedUser, err := uh.us.Save(req)
	if err != nil {
		fmt.Println("[SignupHandler]", err.Error())
		errRes := createErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	savedUserRes := createUserResponse(savedUser)
	res := createSuccessResponse(http.StatusOK, "successfully saved user details", savedUserRes)

	c.JSON(http.StatusCreated, res)
}

func (uh *userHandler) LoginHandler(c *gin.Context) {
	fmt.Println("[LoginHandler] Hitting login handler function in user handler")

	var req *requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("[LoginHandler]", err.Error())
		errRes := createErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user, err := uh.us.Login(req)
	if err != nil {
		fmt.Println("[LoginHandler]", err.Error())
		errRes := createErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	userRes := createUserResponse(user)
	res := createSuccessResponse(http.StatusOK, "successfully logged in", userRes)

	c.JSON(http.StatusOK, res)
}

func (uh *userHandler) UserByUsernameHandler(c *gin.Context) {
	fmt.Println("[UserByUsernameHandler] Hitting user by username handler function in user handler")

	username := c.Param("username")
	_, err := strconv.Atoi(username)
	if err == nil || username == "" {
		fmt.Println("[UserByUsernameHandler]", err.Error())
		errRes := createErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user, err := uh.us.UserByUsername(username)
	if err != nil {
		fmt.Println("[UserByUsernameHandler]", err.Error())
		errRes := createErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	userRes := createUserResponse(user)
	res := createSuccessResponse(http.StatusOK, "successfully got user details", userRes)

	c.JSON(http.StatusOK, res)
}
func (uh *userHandler) GetAllUsers(c *gin.Context) {
	fmt.Println("[GetAllUsersHandler] Hitting get all users handler function in user handler")

	userList, err := uh.us.GetAllUsers()
	if err != nil {
		fmt.Println("[GetAllUsersHandler]", err.Error())
		errRes := createErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	var users []responses.UserResponse
	for _, user := range userList {
		users = append(users, createUserResponse(&user))
	}

	res := createSuccessResponse(http.StatusOK, "successfully got list of users", users)

	c.JSON(http.StatusOK, res)
}

func createUserResponse(user *models.User) responses.UserResponse {
	fmt.Println("[createUserResponse] Creating user response")

	return responses.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}
}

func createSuccessResponse(code int, message string, data interface{}) responses.SuccessResponse {
	fmt.Println("[createSuccessResponse] Creating success response")

	return responses.SuccessResponse{
		Status:  utils.Success,
		Code:    http.StatusText(code),
		Message: message,
		Data:    data,
	}
}

func createErrorResponse(code int, message string) responses.ErrorResponse {
	fmt.Println("[createErrorResponse] Creating error response")

	return responses.ErrorResponse{
		Status:  utils.Error,
		Code:    http.StatusText(code),
		Message: message,
	}
}

func (uh *userHandler) Health(c *gin.Context) {
	fmt.Println("[Health] Hitting health function in user handler")

	c.JSON(http.StatusOK, gin.H{
		"message": "UP!",
	})
}
