package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/controllers"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/utils"
)

func RegisterUserRoutes(engine *gin.Engine, uc controllers.UserController) {
	fmt.Println("[RegisterUserRoutes] Registering routes of the app")

	engine.POST(utils.SignupUserEndpoint, uc.SignupHandler)
	engine.POST(utils.LoginUserEndpoint, uc.LoginHandler)
	engine.GET(utils.UserByUsernameEndpoint, uc.UserByUsernameHandler)
	engine.GET(utils.HealthEndpint, uc.Health)
}
