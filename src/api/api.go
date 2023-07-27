package api

import (
	"database/sql"
	"log"

	"app/api/handlers"
	"app/config"
	"app/infrastructure/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupDatabase() *sql.DB {
	conn := db.Connect()
	return conn
}

func setupRouter(conn *sql.DB) *gin.Engine {
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	handlers.MountUsersHandlers(r, conn)
	handlers.MountRepositoryHandlers(r, conn)

	return r
}

func SetupRouters() *gin.Engine {
	conn := setupDatabase()
	return setupRouter(conn)
}

func StartWebServer() {
	config.ReadEnvironmentVars()
	log.Fatal(SetupRouters().Run())
}
