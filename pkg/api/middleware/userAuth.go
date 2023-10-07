package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	tokenString, err := c.Cookie("userAuth")
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	userID, err := ValidateJWT(tokenString)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	c.Set("userId", userID)
	c.Next()
}
