package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/helpers"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/requests"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/services"
)

type AuthHandler interface {
	SignupHandler(*gin.Context)
	LoginHandler(*gin.Context)
	Health(*gin.Context)
}

type authHandler struct {
	us services.UserService
}

func NewAuthHandler(us services.UserService) AuthHandler {
	fmt.Println("[NewAuthHandler] Initiating New Auth Handler")
	return &authHandler{
		us: us,
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

	savedUserRes := helpers.CreateAuthenticationResponse(savedUser, accessToken, refreshToken)
	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully saved user details", savedUserRes)

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

	userRes := helpers.CreateAuthenticationResponse(user, accessToken, refreshToken)
	res := helpers.CreateSuccessResponse(http.StatusOK, "successfully logged in", userRes)

	fmt.Println("[LoginHandler] Finished execution of login handler")
	c.JSON(http.StatusOK, res)
}

func (ah *authHandler) Health(c *gin.Context) {
	fmt.Println("[HealthHandler] Hitting health function in auth handler")

	c.JSON(http.StatusOK, gin.H{
		"message": "UP!",
	})
}
