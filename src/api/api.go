package api

import (
	"log"

	"app/api/handlers"
	"app/api/middleware"
	"app/config"
	"app/infrastructure/repository"
	usecase_user "app/usecase/user"

	"app/prisma/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {

	dbClient := db.NewClient()
	dbClient.Connect()
	defer dbClient.Disconnect()

	repositoryUser := repository.NewRepositoryUser(dbClient)
	usecaseUser := usecase_user.NewService(repositoryUser)

	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", handlers.HomeHandler)

	// Login do usuario
	r.POST("/api/login", func(gin *gin.Context) {
		handlers.LoginHandler(gin, usecaseUser)
	})

	authorized := r.Group("/api/user")
	authorized.Use(middleware.AuthenticatedMiddleware())
	authorized.GET("/me", func(gin *gin.Context) {
		handlers.GetMeHandler(gin, usecaseUser)
	})

	return r
}

func StartWebServer() {
	config.ReadEnvironmentVars()

	r := SetupRouters()

	// Bind to a port and pass our router in
	log.Fatal(r.Run())
}
