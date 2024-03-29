package routers

import (
	"github.com/easy-health/pkg/api/handler"
	"github.com/easy-health/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AuthSetUpRoute(engine *gin.Engine, UserHandler *handler.UserHandler) {
	auth := engine.Group("/auth")
	auth.GET("google/login", UserHandler.GoogleLogin)
	auth.GET("google/callback", UserHandler.GoogleAuthCallback)

	auth.GET("/google-auth", UserHandler.UserGoogleAuthLoginPage)
	auth.GET("google-auth/callback", UserHandler.UserGoogleAuthCallBack)
	auth.GET("/google-auth/initialize", UserHandler.UserGoogleAuthInitialize)

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
