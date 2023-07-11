package usecase_repository

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_repository.go -package=mocks app/usecase/repository IRepositoryRepository
type IRepositoryRepository interface {
	GetByID(id string) (repository *entity.EntityRepository, err error)
	GetByName(name string) (repository *entity.EntityRepository, err error)
	CreateRepository(repository *entity.EntityRepository) error
	UpdateRepository(repository *entity.EntityRepository) error
	DeleteRepository(repository *entity.EntityRepository) error
}

//go:generate mockgen -destination=../../mocks/mock_usecase_repository.go -package=mocks app/usecase/repository IUsecaseRepository
type IUsecaseRepository interface {
	Create(repository *entity.EntityRepository) error
	Update(repository *entity.EntityRepository) error
	Delete(repository *entity.EntityRepository) error
}
