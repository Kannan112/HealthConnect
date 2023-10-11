package routers

import (
	"github.com/easy-health/pkg/api/handler"
	"github.com/easy-health/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func DoctorSetUpRoute(engine *gin.Engine, doctorHandler *handler.DoctorHandler) {
	doctor := engine.Group("doctor")
	{
		doctor.POST("login", doctorHandler.Login) // get block because its not approved by admin
		doctor.POST("signup/:categoryid", doctorHandler.DoctorRegistration)
		doctor.GET("/categorylist", doctorHandler.ListCategory)
		test := doctor.Group("/profile", middleware.DoctorAuth)
		{
			test.GET("/", doctorHandler.Profile)
		}
	}
}
