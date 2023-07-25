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

func setupDatabase() *db.PrismaClient {
	dbClient := db.NewClient()
	dbClient.Connect()
	return dbClient
}

func setupHandlers(dbClient *db.PrismaClient) (*handlers.UserHandlers, *handlers.RepositoryHandlers) {
	userHandlers := handlers.NewUserHandler(usecase_user.NewService(repository.NewRepositoryUser(dbClient)))
	repoHandlers := handlers.NewRepositoryHandler(usecase_repository.NewService(repository.NewRepositoryRepository(dbClient)))
	return userHandlers, repoHandlers
}

func setupRouter(userHandlers handlers.UserHandlers, repoHandlers handlers.RepositoryHandlers) *gin.Engine {
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/routers", func(ctx *gin.Context) { handlers.RoutersHandler(ctx, r) })
	r.GET("/", handlers.HomeHandler)
	r.POST("/api/login", userHandlers.LoginHandler)

	userGroup := r.Group("/api/user")
	userGroup.Use(middleware.AuthenticatedMiddleware())
	userGroup.GET("/me", userHandlers.GetMeHandler)

	repGroup := r.Group("/api/repository")
	repGroup.GET("/list", repoHandlers.GetRepositoriesHandle)
	repGroup.POST("/create", repoHandlers.CreateRepositoryHandle)
	repGroup.DELETE("/delete/:id", repoHandlers.DeleteRepositoryHandle)

	return r
}

func SetupRouters() *gin.Engine {
	dbClient := setupDatabase()
	userHandlers, repoHandlers := setupHandlers(dbClient)
	return setupRouter(*userHandlers, *repoHandlers)
}

func StartWebServer() {
	config.ReadEnvironmentVars()
	log.Fatal(SetupRouters().Run())
}
