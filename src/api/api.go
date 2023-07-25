package api

import (
	"log"

	"app/api/handlers"
	"app/api/middleware"
	"app/config"
	"app/infrastructure/repository"
	usecase_repository "app/usecase/repository"
	usecase_user "app/usecase/user"

	"app/prisma/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {

	dbClient := db.NewClient()
	dbClient.Connect()
	// defer dbClient.Disconnect()

	repositoryUser := repository.NewRepositoryUser(dbClient)
	usecaseUser := usecase_user.NewService(repositoryUser)

	repositoryRepo := repository.NewRepositoryRepository(dbClient)
	usecaseRepo := usecase_repository.NewService(repositoryRepo)

	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/routers", func(c *gin.Context) {
		type Router struct {
			Method string `json:"method"`
			Path   string `json:"path"`
		}

		var routers []Router = make([]Router, 0)

		for _, route := range r.Routes() {
			routers = append(routers, Router{
				Method: route.Method,
				Path:   route.Path,
			})
		}

		if gin.Mode() == gin.DebugMode {
			c.JSON(200, routers)
		}
	})

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

	authorized = r.Group("/api/repository")
	// authorized.Use(middleware.AuthenticatedMiddleware())
	authorized.GET("/list", func(gin *gin.Context) {
		handlers.GetRepositoriesHandle(gin, usecaseRepo)
	})
	authorized.POST("/create", func(gin *gin.Context) {
		handlers.CreateRepositoryHandle(gin, usecaseRepo)
	})
	authorized.DELETE("/delete/:id", func(gin *gin.Context) {
		handlers.DeleteRepositoryHandle(gin, usecaseRepo)
	})

	return r
}

func StartWebServer() {
	config.ReadEnvironmentVars()

	r := SetupRouters()

	// Bind to a port and pass our router in
	log.Fatal(r.Run())
}
