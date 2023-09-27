package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/easy-health/cmd/api/docs"
	handler "github.com/easy-health/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, doctorHandler *handler.DoctorHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	//	engine.POST("/login", middleware.LoginHandler)

	// Auth middleware
	// api := engine.Group("/api", middleware.AuthorizationMiddleware)

	// api.GET("users", userHandler.FindAll)
	// api.GET("users/:id", userHandler.FindByID)
	// api.POST("users", userHandler.Save)
	// api.DELETE("users/:id", userHandler.Delete)
	engine.POST("reg", doctorHandler.DoctorRegistration)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
