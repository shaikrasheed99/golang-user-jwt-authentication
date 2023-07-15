package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/database"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/handlers"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/repositories"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/routes"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/services"
)

func main() {
	db := database.InitDatabase()

	ur := repositories.NewUserRepository(db)

	us := services.NewUserService(ur)

	uc := handlers.NewUserHandler(us)

	app := gin.Default()

	routes.RegisterAuthRoutes(app, uc)
	routes.RegisterUserRoutes(app, uc)

	app.Run()
}
