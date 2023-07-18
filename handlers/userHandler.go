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

	username := c.Param("username")
	_, err := strconv.Atoi(username)
	if err == nil || username == "" {
		errMessage := constants.InvalidUsernameErrorMessage
		fmt.Println("[UserByUsernameHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusBadRequest, errMessage)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if !isUserMatchesWith(c, username) {
		errMessage := constants.UserIsNotAuthorisedErrorMessage
		fmt.Println("[UserByUsernameHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, errMessage)
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	if !uh.isUserProvidesValidToken(c) {
		errMessage := constants.MaliciousTokenErrorMessage
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

	if !isAdmin(c) {
		errMessage := constants.UserIsNotAuthorisedErrorMessage
		fmt.Println("[GetAllUsersHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, errMessage)
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	if !uh.isUserProvidesValidToken(c) {
		errMessage := constants.MaliciousTokenErrorMessage
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

func (uh *userHandler) isUserProvidesValidToken(c *gin.Context) bool {
	clientToken := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(clientToken, "Bearer ", "", 1)
	username := c.GetString("username")

	tokens, err := uh.as.FindTokensByUsername(username)
	if err != nil {
		fmt.Println("[UserByUsernameHandler]", err.Error())
		return false
	}

	if !areTokensEqual(tokenString, tokens.AccessToken) {
		errMessage := constants.MaliciousTokenErrorMessage
		fmt.Println("[UserByUsernameHandler]", errMessage)
		return false
	}

	return true
}

func areTokensEqual(tokenOne string, tokenTwo string) bool {
	return tokenOne == tokenTwo
}

func isUserMatchesWith(c *gin.Context, inputUsername string) bool {
	if isAdmin(c) {
		return true
	}

	if !isUser(c) {
		return false
	}

	username := c.GetString("username")
	if isEmpty(username) || !isEqual(username, inputUsername) {
		return false
	}

	return true
}

func isAdmin(c *gin.Context) bool {
	role := c.GetString("role")
	if isEmpty(role) || !isEqual(role, "admin") {
		return false
	}

	return true
}

func isUser(c *gin.Context) bool {
	role := c.GetString("role")
	if isEmpty(role) || !isEqual(role, "user") {
		return false
	}

	return true
}

func isEmpty(value string) bool {
	return value == ""
}

func isEqual(valueOne string, valueTwo string) bool {
	valueOne = strings.ToLower(valueOne)
	valueTwo = strings.ToLower(valueTwo)
	return valueOne == valueTwo
}
