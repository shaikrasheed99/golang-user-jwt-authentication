package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/helpers"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/responses"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/services"
)

type UserHandler interface {
	UserByUsernameHandler(*gin.Context)
	GetAllUsers(*gin.Context)
}

type userHandler struct {
	us services.UserService
	as services.AuthService
}

func NewUserHandler(us services.UserService, as services.AuthService) UserHandler {
	fmt.Println("[NewUserHandler] Initiating New User Handler")
	return &userHandler{
		us: us,
		as: as,
	}
}

func (uh *userHandler) UserByUsernameHandler(c *gin.Context) {
	fmt.Println("[UserByUsernameHandler] Hitting user by username handler function in user handler")

	username := c.Param(constants.Username)
	_, err := strconv.Atoi(username)
	if err == nil || username == "" {
		errMessage := constants.ErrInvalidUsername
		fmt.Println("[UserByUsernameHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusBadRequest, errMessage)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if !helpers.IsUserMatchesWith(c, username) {
		errMessage := constants.ErrUserIsNotAuthorised
		fmt.Println("[UserByUsernameHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, errMessage)
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	if !uh.isUserProvidesValidAccessToken(c) {
		errMessage := constants.ErrMaliciousToken
		fmt.Println("[UserByUsernameHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, errMessage)
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	user, err := uh.us.UserByUsername(username)
	if err != nil {
		fmt.Println("[UserByUsernameHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	userRes := helpers.CreateUserResponse(user)
	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully got user details", userRes)

	fmt.Println("[UserByUsernameHandler] Finished execution of user by username handler")
	c.JSON(http.StatusOK, res)
}

func (uh *userHandler) GetAllUsers(c *gin.Context) {
	fmt.Println("[GetAllUsersHandler] Hitting get all users handler function in user handler")

	if !helpers.IsAdmin(c) {
		errMessage := constants.ErrUserIsNotAuthorised
		fmt.Println("[GetAllUsersHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, errMessage)
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	if !uh.isUserProvidesValidAccessToken(c) {
		errMessage := constants.ErrMaliciousToken
		fmt.Println("[GetAllUsersHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, errMessage)
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	userList, err := uh.us.GetAllUsers()
	if err != nil {
		fmt.Println("[GetAllUsersHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	var users []responses.UserResponse
	for _, user := range userList {
		users = append(users, helpers.CreateUserResponse(&user))
	}

	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully got list of users", users)

	fmt.Println("[GetAllUsersHandler] Finished execution of get all users handler")
	c.JSON(http.StatusOK, res)
}

func (uh *userHandler) isUserProvidesValidAccessToken(c *gin.Context) bool {
	clientToken := c.Request.Header.Get(constants.Authorization)
	tokenString := strings.Replace(clientToken, "Bearer ", "", 1)
	username := c.GetString(constants.Username)

	tokens, err := uh.as.FindTokensByUsername(username)
	if err != nil {
		fmt.Println("[isUserProvidesValidAccessToken]", err.Error())
		return false
	}

	if !helpers.AreTokensEqual(tokenString, tokens.AccessToken) {
		errMessage := constants.ErrMaliciousToken
		fmt.Println("[isUserProvidesValidAccessToken]", errMessage)
		return false
	}

	return true
}
