package routers

import (
	"github.com/easy-health/pkg/api/handler"
	"github.com/easy-health/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AuthSetUpRoute(engine *gin.Engine, UserHandler *handler.UserHandler) {
	//auth := engine.Group("/auth")

	// Routes for user authentication
	//auth.GET("/login", UserHandler.UserGoogleAuthLoginPage)
	// Add more authentication routes as needed
	user := engine.Group("/user")
	{
		user.POST("/login", UserHandler.Login)
		user.POST("/register", UserHandler.Register)
		user.GET("/logout", UserHandler.Logout)

		{
			category := user.Group("/category", middleware.UserAuth)
			category.GET("/", UserHandler.ListCategory)
		}

	}
}
