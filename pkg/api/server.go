package http

import (
	_ "github.com/easy-health/cmd/api/docs"
	"github.com/easy-health/pkg/api/handler"
	"github.com/easy-health/pkg/api/middleware"
	routers "github.com/easy-health/pkg/api/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, doctorHandler *handler.DoctorHandler, adminHandler *handler.AdminHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// adminSIDE
	admin := engine.Group("admin")
	admin.POST("create", adminHandler.AdminSignup)
	admin.POST("login", adminHandler.AdminLogin)

	category := admin.Group("category", middleware.AdminAuth)
	{
		category.GET("", adminHandler.ListCategory)          // Handler for listing categories
		category.POST("create", adminHandler.CreateCategory) // Handler for creating a category
		category.DELETE("/:id", adminHandler.DeleteCategory)
	}
	Admindoctors := admin.Group("/doctors", middleware.AdminAuth)
	{
		Admindoctors.GET("list", adminHandler.ListDoctorsNotApproved)
		Admindoctors.POST("/approve/:id", adminHandler.ApproveDoctor)
	}

	//doctor
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
	//patient
	routers.SetUpUserRoutes(engine, userHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
