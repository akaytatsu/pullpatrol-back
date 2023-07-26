package usecase_repository

import (
	"app/entity"
	"app/infrastructure/git"
)

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_repository.go -package=mocks app/usecase/repository IRepositoryRepository
type IRepositoryRepository interface {
	GetByID(id int) (repository *entity.EntityRepository, err error)
	GetByRepo(repo string) (repository *entity.EntityRepository, err error)
	CreateOrUpdateRepository(repository *entity.EntityRepository) error
	DeleteRepository(repository *entity.EntityRepository) error
	GetRepositories() (repositories []entity.EntityRepository, err error)
	CreateOrUpdatePullRequest(pullRequest *entity.EntityPullRequest) error
}

//go:generate mockgen -destination=../../mocks/mock_usecase_repository.go -package=mocks app/usecase/repository IUsecaseRepository
type IUsecaseRepository interface {
	Get(id int) (repository *entity.EntityRepository, err error)
	Create(repository *entity.EntityRepository) error
	Update(repository *entity.EntityRepository) error
	Delete(repository *entity.EntityRepository) error
	GetRepositories() (repositories []entity.EntityRepository, err error)
	ProccessPullRequest(git git.GitInterface, payload []byte) (err error)
}
