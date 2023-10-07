package handlerurtl

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdminIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("adminId")
	if id == nil {
		return 0, fmt.Errorf("adminId not found in context")
	}
	adminID, err := strconv.Atoi(fmt.Sprintf("%v", id))
	if err != nil {
		return 0, fmt.Errorf("adminId is not of type int in context")
	}
	return adminID, nil
}

func DoctorIdFromContext(c *gin.Context) (int, error) {
	id, exists := c.Get("doctorId")
	if !exists {
		// Handle the case where the value doesn't exist
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Value not found"})
		return 0, fmt.Errorf("Value not found")
	}
	userId, err := strconv.Atoi(fmt.Sprintf("%v", id))
	fmt.Println(c.Value("doctorId"))
	if err != nil {
		fmt.Println("1 test")
	}
	return userId, err

}
func UserIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("userId")
	userID, err := strconv.Atoi(fmt.Sprintf("%v", id))
	if err != nil {
		return 0, err
	}
	return userID, nil

}
