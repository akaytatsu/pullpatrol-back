package main

import (
	"app/api"
	"app/cron"
	"app/infrastructure/db"
	"app/infrastructure/repository"
	usecase_user "app/usecase/user"
)

func main() {
	cron.StartCronJobs()

	conn := db.Connect()
	db.Migrations()

	// create default user
	repo := repository.NewRepositoryUser(conn)
	usecase := usecase_user.NewService(repo)
	usecase.CreateAdminUser()

	api.StartWebServer()
}
