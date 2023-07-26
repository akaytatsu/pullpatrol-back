package api

import (
	"database/sql"
	"log"

	"app/api/handlers"
	"app/config"
	"app/infrastructure/db"
	"app/infrastructure/repository"
	usecase_repository "app/usecase/repository"
	usecase_user "app/usecase/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupDatabase() *sql.DB {
	conn := db.Connect()
	return conn
}

func setupHandlers(conn *sql.DB) (*handlers.UserHandlers, *handlers.RepositoryHandlers) {
	userHandlers := handlers.NewUserHandler(
		usecase_user.NewService(
			repository.NewRepositoryUser(conn),
		),
	)

	repoHandlers := handlers.NewRepositoryHandler(
		usecase_repository.NewService(
			repository.NewRepositoryRepository(conn),
		),
	)
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
	// userGroup.Use(middleware.AuthenticatedMiddleware())
	userGroup.GET("/me", userHandlers.GetMeHandler)
	userGroup.GET("/", userHandlers.GetUsersHandler)
	userGroup.GET("/:id", userHandlers.GetUserHandler)

	repGroup := r.Group("/api/repository")
	repGroup.GET("", repoHandlers.GetRepositoriesHandle)
	repGroup.POST("", repoHandlers.CreateRepositoryHandle)
	repGroup.DELETE("/:id", repoHandlers.DeleteRepositoryHandle)

	r.POST("/git-webhook", repoHandlers.GitWebhookHandler)

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
