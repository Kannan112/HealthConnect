package routers

import (
	"github.com/easy-health/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func SetUpUserRoutes(engine *gin.Engine, Handler *handler.UserHandler) {
	user := engine.Group("/user")
	{
		user.POST("/login", Handler.Login)
		user.POST("/register", Handler.Register)
		user.GET("/logout", Handler.Logout)
	}
}
