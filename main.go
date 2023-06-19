package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/database"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/routes"
)

func main() {
	db := database.InitDatabase()

	fmt.Printf("db: %v\n", db)

	app := gin.Default()

	routes.RegisterUserRoutes(app)

	app.Run()
}
