package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/helpers"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/middlewares"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/requests"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/services"
)

type IAuthHandler interface {
	SignupHandler(*gin.Context)
	LoginHandler(*gin.Context)
	LogoutHandler(*gin.Context)
	RefreshTokenHandler(*gin.Context)
	Health(*gin.Context)
}

type authHandler struct {
	us services.IUserService
	as services.IAuthService
}

func NewAuthHandler(us services.IUserService, as services.IAuthService) IAuthHandler {
	fmt.Println("[NewAuthHandler] Initiating New Auth Handler")
	return &authHandler{
		us: us,
		as: as,
	}
}

func (ah *authHandler) SignupHandler(c *gin.Context) {
	fmt.Println("[SignupHandler] Hitting signup handler function in auth handler")

	var req *requests.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("[SignupHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	savedUser, err := ah.us.Save(req)
	if err != nil {
		fmt.Println("[SignupHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	accessToken, refreshToken, err := helpers.GenerateToken(savedUser.Username, savedUser.Role)
	if err != nil {
		fmt.Println("[SignupHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	err = ah.as.SaveTokensByUsername(savedUser.Username, accessToken, refreshToken)
	if err != nil {
		fmt.Println("[SignupHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully saved user details", nil)

	helpers.SetAccessAndRefreshTokenCookies(c, accessToken, refreshToken)

	fmt.Println("[SignupHandler] Finished execution of signup handler")
	c.JSON(http.StatusCreated, res)
}

func (ah *authHandler) LoginHandler(c *gin.Context) {
	fmt.Println("[LoginHandler] Hitting login handler function in auth handler")

	var req *requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("[LoginHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user, err := ah.us.Login(req)
	if err != nil {
		fmt.Println("[LoginHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	accessToken, refreshToken, err := helpers.GenerateToken(user.Username, user.Role)
	if err != nil {
		fmt.Println("[LoginHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	err = ah.as.SaveTokensByUsername(user.Username, accessToken, refreshToken)
	if err != nil {
		fmt.Println("[LoginHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully logged in", nil)

	helpers.SetAccessAndRefreshTokenCookies(c, accessToken, refreshToken)

	fmt.Println("[LoginHandler] Finished execution of login handler")
	c.JSON(http.StatusOK, res)
}

func (ah *authHandler) LogoutHandler(c *gin.Context) {
	fmt.Println("[LogoutHandler] Hitting logout handler function in auth handler")

	middlewares.Authentication(c)
	if c.IsAborted() {
		return
	}

	var req *requests.LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("[LogoutHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if !helpers.IsUserMatchesWith(c, req.Username) {
		errMessage := constants.ErrUserIsNotAuthorised
		fmt.Println("[LogoutHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, errMessage)
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	if !ah.isUserProvidesValidAccessToken(c) {
		errMessage := constants.ErrMaliciousToken
		fmt.Println("[LogoutHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, errMessage)
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	err := ah.as.DeleteTokensByUsername(req.Username)
	if err != nil {
		fmt.Println("[LogoutHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully logged out", nil)

	helpers.SetAccessAndRefreshTokenCookies(c, "", "")

	fmt.Println("[LogoutHandler] Finished execution of login handler")
	c.JSON(http.StatusOK, res)
}

func (ah *authHandler) RefreshTokenHandler(c *gin.Context) {
	fmt.Println("[RefreshTokenHandler] Hitting refresh token handler function in auth handler")

	middlewares.Authentication(c)
	if c.IsAborted() {
		return
	}

	var req *requests.LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("[RefreshTokenHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if !helpers.IsUserMatchesWith(c, req.Username) {
		errMessage := constants.ErrUserIsNotAuthorised
		fmt.Println("[RefreshTokenHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, errMessage)
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	if !ah.isUserProvidesValidRefreshToken(c) {
		errMessage := constants.ErrMaliciousToken
		fmt.Println("[RefreshTokenHandler]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, errMessage)
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	role := c.GetString(constants.Role)
	accessToken, refreshToken, err := helpers.GenerateToken(req.Username, role)
	if err != nil {
		fmt.Println("[RefreshTokenHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	err = ah.as.SaveTokensByUsername(req.Username, accessToken, refreshToken)
	if err != nil {
		fmt.Println("[RefreshTokenHandler]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully received access token", nil)

	helpers.SetAccessAndRefreshTokenCookies(c, accessToken, refreshToken)

	fmt.Println("[RefreshTokenHandler] Finished execution of refresh token handler")
	c.JSON(http.StatusOK, res)
}

func (ah *authHandler) Health(c *gin.Context) {
	fmt.Println("[HealthHandler] Hitting health function in auth handler")

	c.JSON(http.StatusOK, gin.H{
		"message": "UP!",
	})
}

func (ah *authHandler) isUserProvidesValidAccessToken(c *gin.Context) bool {
	clientToken := c.Request.Header.Get(constants.Authorization)
	tokenString := strings.Replace(clientToken, "Bearer ", "", 1)
	username := c.GetString(constants.Username)

	tokens, err := ah.as.FindTokensByUsername(username)
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

func (ah *authHandler) isUserProvidesValidRefreshToken(c *gin.Context) bool {
	clientToken := c.Request.Header.Get(constants.Authorization)
	tokenString := strings.Replace(clientToken, "Bearer ", "", 1)
	username := c.GetString(constants.Username)

	tokens, err := ah.as.FindTokensByUsername(username)
	if err != nil {
		fmt.Println("[isUserProvidesValidRefreshToken]", err.Error())
		return false
	}

	if !helpers.AreTokensEqual(tokenString, tokens.RefreshToken) {
		errMessage := constants.ErrMaliciousToken
		fmt.Println("[isUserProvidesValidRefreshToken]", errMessage)
		return false
	}

	return true
}
