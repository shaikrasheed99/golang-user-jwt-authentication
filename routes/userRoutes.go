package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/controllers"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/utils"
)

func RegisterUserRoutes(engine *gin.Engine, uc controllers.UserController) {
	engine.POST(utils.SignupUserEndpoint, uc.SignupHandler)
	engine.POST(utils.LoginUserEndpoint, uc.LoginHandler)
	engine.GET(utils.HealthEndpint, uc.Health)
}
