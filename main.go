package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/configs"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/database"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/handlers"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/repositories"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/routes"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/services"
)

func main() {
	configs.LoadConfigs()

	db := database.InitDatabase()

	ur := repositories.NewUserRepository(db)

	us := services.NewUserService(ur)

	ah := handlers.NewAuthHandler(us)
	uc := handlers.NewUserHandler(us)

	app := gin.Default()

	routes.RegisterAuthRoutes(app, ah)
	routes.RegisterUserRoutes(app, uc)

	app.Run()
}
