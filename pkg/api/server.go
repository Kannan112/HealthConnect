package http

import (
	_ "github.com/easy-health/cmd/api/docs"
	"github.com/easy-health/pkg/api/handler"
	routers "github.com/easy-health/pkg/api/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, doctorHandler *handler.DoctorHandler, adminHandler *handler.AdminHandler) *ServerHTTP {
	// Create a new Gin engine
	engine := gin.New()

	// Use Gin's built-in logger middleware
	engine.Use(gin.Logger())

	// Serve Swagger API documentation
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Set up routes for admin users
	routers.AdminSetUpRoute(engine, adminHandler)

	// Set up routes for doctors
	routers.DoctorSetUpRoute(engine, doctorHandler)

	// Set up routes for patients
	routers.SetUpUserRoutes(engine, userHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
