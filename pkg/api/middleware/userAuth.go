package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	tokenString, err := c.Cookie("UserAuth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	// Assuming this is the correct way to access your configuration

	userID, err := ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized%v", "err": err.Error()})
		c.Abort()
		return
	}
	//c.Set("role", role)
	c.Set("userId", userID)
	c.Next()
}
