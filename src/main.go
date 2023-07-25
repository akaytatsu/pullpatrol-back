package main

import (
	"app/api"
	"app/cron"
	"app/infrastructure/repository"
	"app/prisma/db"
	usecase_user "app/usecase/user"
)

func main() {
	cron.StartCronJobs()

	dbClient := db.NewClient()
	if err := dbClient.Connect(); err != nil {
		panic(err)
	}

	// create default user
	repo := repository.NewRepositoryUser(dbClient)
	usecase := usecase_user.NewService(repo)
	usecase.CreateAdminUser()

	defer dbClient.Disconnect()

	api.StartWebServer()
}
