package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/helpers"
)

func Authentication(c *gin.Context) {
	fmt.Println("[AuthenticationMiddleware] Checking for authentication of user")

	clientToken := c.Request.Header.Get(constants.Authorization)
	tokenString := strings.Replace(clientToken, "Bearer ", "", 1)
	if tokenString == "" {
		errMessage := constants.NoAuthHeaderErrorMessage
		fmt.Println("[AuthenticationMiddleware]", errMessage)
		errRes := helpers.CreateErrorResponse(http.StatusBadRequest, errMessage)
		c.JSON(http.StatusBadRequest, errRes)
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(tokenString)
	if err != nil {
		fmt.Println("[AuthenticationMiddleware]", err.Error())
		errRes := helpers.CreateErrorResponse(http.StatusUnauthorized, err.Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}

	fmt.Println("[AuthenticationMiddleware] User is authenticated")
	c.Set("username", claims.Username)
	c.Set("role", claims.Role)
	c.Next()
}
