package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/handlers"
)

func RegisterAuthRoutes(engine *gin.Engine, uc handlers.UserHandler) {
	fmt.Println("[RegisterAuthRoutes] Registering auth routes of the app")

	engine.POST(constants.SignupUserEndpoint, uc.SignupHandler)
	engine.POST(constants.LoginUserEndpoint, uc.LoginHandler)
	engine.GET(constants.HealthEndpint, uc.Health)
}
