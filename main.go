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
	ar := repositories.NewAuthRepository(db)

	us := services.NewUserService(ur)
	as := services.NewAuthService(ar)

	ah := handlers.NewAuthHandler(us, as)
	uh := handlers.NewUserHandler(us)

	app := gin.Default()

	routes.RegisterAuthRoutes(app, ah)
	routes.RegisterUserRoutes(app, uh)

	app.Run()
}
