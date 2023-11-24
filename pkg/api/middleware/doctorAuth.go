package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoctorAuth(c *gin.Context) {
	tokenString, err := c.Cookie("DoctorAuth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	doctorID, err := ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Set the doctorID in the Gin context
	//fmt.Println(role)
	c.Set("doctorId", doctorID)
	c.Next()
}
