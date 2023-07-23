package repository

import (
	"app/entity"
	"app/prisma/db"
	"context"
)

type RepositoryRepository struct {
	db *db.PrismaClient
}

func NewRepositoryRepository(db *db.PrismaClient) *RepositoryRepository {
	return &RepositoryRepository{db: db}
}

func (r *RepositoryRepository) GetByID(id int) (repository *entity.EntityRepository, err error) {

	context := context.Background()

	model, err := r.db.Repository.FindUnique(
		db.Repository.ID.Equals(id),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	repository = &entity.EntityRepository{
		ID:   model.ID,
		Name: model.Name,
		Url:  model.URL,
	}

	return repository, err
}

func (r *RepositoryRepository) GetByName(name string) (repository *entity.EntityRepository, err error) {
	context := context.Background()

	model, err := r.db.Repository.FindFirst(
		db.Repository.Name.Equals(name),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	repository = &entity.EntityRepository{
		ID:   model.ID,
		Name: model.Name,
		Url:  model.URL,
	}

	return repository, err
}

func (r *RepositoryRepository) CreateRepository(repository *entity.EntityRepository) error {
	context := context.Background()

	_, err := r.db.Repository.FindUnique(
		db.Repository.ID.Equals(repository.ID),
	).Exec(context)

	if err != nil {
		return err
	}

	_, err = r.db.Repository.CreateOne(
		db.Repository.Name.Equals(repository.Name),
		db.Repository.URL.Equals(repository.Url),
	).Exec(context)

	return err
}

func (r *RepositoryRepository) UpdateRepository(repository *entity.EntityRepository) error {

	context := context.Background()

	_, err := r.db.Repository.FindUnique(
		db.Repository.ID.Equals(repository.ID),
	).Exec(context)

	if err != nil {
		return err
	}

	_, err = r.db.Repository.FindUnique(
		db.Repository.ID.Equals(repository.ID),
	).Update(
		db.Repository.Name.Set(repository.Name),
	).Exec(context)

	return err
}

func (r *RepositoryRepository) DeleteRepository(repository *entity.EntityRepository) error {

	context := context.Background()

	_, err := r.db.Repository.FindUnique(
		db.Repository.ID.Equals(repository.ID),
	).Exec(context)

	if err != nil {
		return err
	}

	_, err = r.db.Repository.FindUnique(
		db.Repository.ID.Equals(repository.ID),
	).Delete().Exec(context)

	return err
}
