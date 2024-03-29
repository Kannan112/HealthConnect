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

		// Create a "category" route group under "admin" with admin authentication middleware
		category := admin.Group("/categories", middleware.AdminAuthMiddleware)
		{
			category.GET("/", adminHandler.ListCategory)
			category.POST("/", adminHandler.CreateCategory) // Handler for creating a category
			//	category.PUT("/",adminHandler.)
			category.DELETE("/:id", adminHandler.DeleteCategory)
		}
		//	admin.GET("/list-doctors-not-approved", adminHandler.ListDoctorsNotApproved, middleware.AdminAuth)

		// Create a "middler" route group under "admin" with admin authentication middleware
		doctor := admin.Group("doctors", middleware.AdminAuthMiddleware)
		{
			doctor.GET("/")
			doctor.GET("/verified", adminHandler.VerifiedDoctors)
			doctor.GET("/not-approved", adminHandler.ListDoctorsNotApproved, middleware.AdminAuthMiddleware)
			doctor.PATCH("/approve/:id", adminHandler.ApproveDoctor)
		}
	}
}
