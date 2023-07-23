package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/handlers"
)

func RegisterAuthRoutes(engine *gin.Engine, ah handlers.AuthHandler) {
	fmt.Println("[RegisterAuthRoutes] Registering auth routes of the app")

	engine.POST(constants.SignupUserEndpoint, ah.SignupHandler)
	engine.POST(constants.LoginUserEndpoint, ah.LoginHandler)
	engine.POST(constants.LogoutUserEndpoint, ah.LogoutHandler)
	engine.POST(constants.RefreshAccessTokenEndpoint, ah.RefreshTokenHandler)
	engine.GET(constants.HealthEndpint, ah.Health)
}
