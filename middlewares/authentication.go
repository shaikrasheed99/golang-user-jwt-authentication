package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticator(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {

	}
	c.Next()
}
