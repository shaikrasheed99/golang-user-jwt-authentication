package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/handlers"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/middlewares"
)

func RegisterUserRoutes(engine *gin.Engine, uc handlers.UserHandler) {
	fmt.Println("[RegisterUserRoutes] Registering user routes of the app")

	engine.Use(middlewares.Authenticator)
	engine.GET(constants.UserByUsernameEndpoint, uc.UserByUsernameHandler)
	engine.GET(constants.GetAllUsersEndpoint, uc.GetAllUsers)
}
