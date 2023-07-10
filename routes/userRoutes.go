package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/controllers"
)

const (
	SignupUserEndpoint = "/signup"
	HealthEndpint      = "/health"
)

func RegisterUserRoutes(engine *gin.Engine, uc controllers.UserController) {
	engine.POST(SignupUserEndpoint, uc.Signup)
	engine.GET(HealthEndpint, uc.Health)
}
