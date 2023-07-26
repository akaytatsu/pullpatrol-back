package usecase_repository

import (
	"app/entity"
	"app/infrastructure/git"
	"app/infrastructure/git/github"
	"errors"
)

type UsecaseRepository struct {
	repo IRepositoryRepository
}

func NewService(repository IRepositoryRepository) *UsecaseRepository {
	return &UsecaseRepository{repo: repository}
}

func (u *UsecaseRepository) Create(repository *entity.EntityRepository) error {

	err := repository.GetValidated()

	if err != nil {
		return err
	}

	repository.Active = true

	return u.repo.CreateOrUpdateRepository(repository)
}

func (u *UsecaseRepository) Update(repository *entity.EntityRepository) error {
	return u.repo.CreateOrUpdateRepository(repository)
}

func (u *UsecaseRepository) Delete(repository *entity.EntityRepository) error {
	return u.repo.DeleteRepository(repository)
}

func (u *UsecaseRepository) GetRepositories() (repositories []entity.EntityRepository, err error) {
	return u.repo.GetRepositories()
}

func (u *UsecaseRepository) Get(id int) (repository *entity.EntityRepository, err error) {
	return u.repo.GetByID(id)
}

func (u *UsecaseRepository) ProccessPullRequest(git git.GitInterface, payload []byte) (err error) {
	var repo entity.EntityRepository
	var entityPR entity.EntityPullRequest

	data, err := git.ProccessWebhook(payload)

	if err != nil {
		return err
	}

	if git.Driver() == "github" {
		structuredData := data.(github.GitHubWebhookPullRequest)

		repo = entity.EntityRepository{
			Repository: structuredData.Repository.CloneURL,
		}

		err := u.repo.CreateOrUpdateRepository(&repo)

		if err != nil {
			return err
		}

		entityPR = entity.EntityPullRequest{
			Number:        structuredData.PullRequest.Number,
			Repository:    repo,
			RepositoryID:  repo.ID,
			Title:         structuredData.PullRequest.Title,
			Action:        structuredData.Action,
			Status:        structuredData.PullRequest.State,
			URL:           structuredData.PullRequest.HTMLURL,
			RepositoryURL: structuredData.Repository.CloneURL,
			CreatedAt:     structuredData.PullRequest.CreatedAt,
			UpdatedAt:     structuredData.PullRequest.UpdatedAt,
			Additions:     structuredData.PullRequest.Additions,
			Deletions:     structuredData.PullRequest.Deletions,
			ChangedFiles:  structuredData.PullRequest.ChangedFiles,
			Commits:       structuredData.PullRequest.Commits,
		}

		err = u.repo.CreateOrUpdatePullRequest(&entityPR)

		if err != nil {
			return err
		}

	} else {
		return errors.New("driver not found")
	}

	return nil
}
