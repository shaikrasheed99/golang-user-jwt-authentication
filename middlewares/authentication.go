package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/helpers"
)

func Authentication(c *gin.Context) {
	clientToken := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(clientToken, "Bearer ", "", 1)
	if tokenString == "" {
		errMessage := "No Authorization Header Provided"
		fmt.Println("[Authentication]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, errMessage)
		c.JSON(http.StatusForbidden, errRes)
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(tokenString)
	if err != nil {
		fmt.Println("[Authentication]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusForbidden, errRes)
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
	c.Set("role", claims.Role)
	c.Next()
}
