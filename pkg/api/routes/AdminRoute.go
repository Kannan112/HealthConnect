package routers

import (
	"github.com/easy-health/pkg/api/handler"
	"github.com/easy-health/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminSetUpRoute(engine *gin.Engine, adminHandler *handler.AdminHandler) {
	// Create an "admin" route group
	admin := engine.Group("/admin")
	{
		// Routes for admin authentication
		admin.POST("/create", adminHandler.AdminSignup)
		admin.POST("/login", adminHandler.AdminLogin)
		admin.GET("/logout", adminHandler.AdminLogout)
		admin.GET("/list-category", adminHandler.ListCategory)

		// Create a "category" route group under "admin" with admin authentication middleware
		category := admin.Group("/category", middleware.AdminAuth)
		{
			category.POST("/create", adminHandler.CreateCategory) // Handler for creating a category
			category.DELETE("/:id", adminHandler.DeleteCategory)
		}

		// Create a "middler" route group under "admin" with admin authentication middleware
		middler := admin.Group("", middleware.AdminAuth)
		{
			middler.GET("/list-doctors-not-approved", adminHandler.ListDoctorsNotApproved)
			middler.POST("/approve-doctor/:id", adminHandler.ApproveDoctor)
		}
	}
}
