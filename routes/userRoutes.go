package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/controllers"
)

const (
	SignupUserEndpoint = "/signup"
	LoginUserEndpoint  = "/login"
	HealthEndpint      = "/health"
)

func RegisterUserRoutes(engine *gin.Engine, uc controllers.UserController) {
	engine.POST(SignupUserEndpoint, uc.SignupHandler)
	engine.POST(LoginUserEndpoint, uc.LoginHandler)
	engine.GET(HealthEndpint, uc.Health)
}
