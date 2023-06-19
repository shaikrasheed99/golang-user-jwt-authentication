package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/routes"
)

func main() {
	app := gin.Default()

	routes.RegisterUserRoutes(app)

	app.Run()
}
