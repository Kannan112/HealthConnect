package http

import (
	_ "github.com/easy-health/cmd/api/docs"
	"github.com/easy-health/pkg/api/handler"
	routers "github.com/easy-health/pkg/api/routes"
	"github.com/easy-health/pkg/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

// @title						E-commerce Application Backend API
// @description				Backend API built with Golang using Clean Code architecture. \nGithub: [https://github.com/kannan112/easy-health].
// @contact.name				For API Support
// @contact.email				abhinandarun369@gmail.com
// @license.name				MIT
// @license.url				https://opensource.org/licenses/MIT
// @SecurityDefinitions.apikey	BearerAuth
// @Name						Authorization
// @In							header
// @Description				Add prefix of Bearer before  token Ex: "Bearer token"
// @Query.collection.format	multi
func NewServerHTTP(userHandler *handler.UserHandler, doctorHandler *handler.DoctorHandler, adminHandler *handler.AdminHandler) *ServerHTTP {
	// Create a new Gin engine
	engine := gin.New()

	// Use Gin's built-in logger middleware
	engine.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Allow requests from your React app's origin
	config.AllowHeaders = []string{"*"}                     // Allow any headers
	engine.Use(cors.New(config))                            // Use the CORS middleware

	services.AllRooms.Init()

	// Serve Swagger API documentation
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Set up routes for admin users
	routers.AdminSetUpRoute(engine, adminHandler)

	// Set up routes for doctors
	routers.DoctorSetUpRoute(engine, doctorHandler)

	// Set up routes for patients
	engine.GET("/create", services.CreateRoomRequestHandler)

	engine.GET("/join", services.JoinRoomRequestHandler)

	routers.AuthSetUpRoute(engine, userHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8000")
}
