package usecase_repository

import (
	"app/entity"
	"app/infrastructure/git"
)

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_repository.go -package=mocks app/usecase/repository IRepositoryRepository
type IRepositoryRepository interface {
	// repository
	GetByID(id int) (repository *entity.EntityRepository, err error)
	GetByRepo(repo string) (repository *entity.EntityRepository, err error)
	CreateOrUpdateRepository(repository *entity.EntityRepository) error
	DeleteRepository(repository *entity.EntityRepository) error
	GetRepositories() (repositories []entity.EntityRepository, err error)

	// pull request
	CreateOrUpdatePullRequest(pullRequest *entity.EntityPullRequest) error

	// pull request role
	GetPullRequestRoles(repositoryID int) (pullRequestRoles []entity.EntityPullRequestRole, err error)
	GetPullRequestRole(pullRequestRoleID int) (pullRequestRole *entity.EntityPullRequestRole, err error)
	CreatePullRequestRole(repositoryID int, pullRequestRole *entity.EntityPullRequestRole) error
	UpdatePullRequestRole(pullRequestRoleID int, pullRequestRole *entity.EntityPullRequestRole) error
	DeletePullRequestRole(pullRequestRoleID int) error
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
