package main

import (
	"app/api"
	"app/cron"
	"app/entity"
	"app/infrastructure/postgres"
	"app/infrastructure/repository"
	usecase_user "app/usecase/user"
)

func main() {
	cron.StartCronJobs()

	db, err := postgres.Connect()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.EntityUser{})

	// create default user
	repo := repository.NewUserPostgres(db)
	usecase := usecase_user.NewService(repo)
	usecase.CreateAdminUser()

	api.StartWebServer()
}
