package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/handlers"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/utils"
)

func RegisterUserRoutes(engine *gin.Engine, uc handlers.UserHandler) {
	fmt.Println("[RegisterUserRoutes] Registering routes of the app")

	engine.POST(utils.SignupUserEndpoint, uc.SignupHandler)
	engine.POST(utils.LoginUserEndpoint, uc.LoginHandler)
	engine.GET(utils.UserByUsernameEndpoint, uc.UserByUsernameHandler)
	engine.GET(utils.GetAllUsersEndpoint, uc.GetAllUsers)
	engine.GET(utils.HealthEndpint, uc.Health)
}
