package routers

import (
	"github.com/easy-health/pkg/api/handler"
	"github.com/easy-health/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminSetUpRoute(engine *gin.Engine, AdminHandler *handler.AdminHandler) {
	admin := engine.Group("/admin")
	{
		admin.POST("create", AdminHandler.AdminSignup)
		admin.POST("login", AdminHandler.AdminLogin)
		admin.GET("logout", AdminHandler.AdminLogout)

		category := admin.Group("category", middleware.AdminAuth)
		{
			category.GET("", AdminHandler.ListCategory)          // Handler for listing categories
			category.POST("create", AdminHandler.CreateCategory) // Handler for creating a category
			category.DELETE("/:id", AdminHandler.DeleteCategory)
		}

		Admindoctors := admin.Group("/doctors", middleware.AdminAuth)
		{
			Admindoctors.GET("list", AdminHandler.ListDoctorsNotApproved)
			Admindoctors.POST("/approve/:id", AdminHandler.ApproveDoctor)
		}

	}
}
