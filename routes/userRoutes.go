package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/controllers"
)

const (
	HealthEndpint = "/health"
)

func RegisterUserRoutes(engine *gin.Engine) {
	engine.GET(HealthEndpint, controllers.Health)
}
