package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuth(c *gin.Context) {
	TokenString, err := c.Cookie("AdminAuth")
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	adminID, err := ValidateJWT(TokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("adminId", adminID)

	c.Next()
}
