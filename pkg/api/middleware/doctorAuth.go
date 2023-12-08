package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DoctorAuth(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

	doctorID, err := ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Set the doctorID in the Gin context	//fmt.Println(role)
	c.Set("doctorId", doctorID)
	c.Next()
}
