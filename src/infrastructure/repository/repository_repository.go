package repository

import (
	"app/entity"
	"app/infrastructure/db/queries"
	"context"
	"database/sql"
	"errors"
	"time"
)

type RepositoryRepository struct {
	db      *sql.DB
	queries *queries.Queries
}

func NewRepositoryRepository(db *sql.DB) *RepositoryRepository {
	return &RepositoryRepository{db: db, queries: queries.New(db)}
}

func (r *RepositoryRepository) GetByID(id int) (repository *entity.EntityRepository, err error) {

	context := context.Background()

	qRepo, err := r.queries.GetRepositoryByID(context, int64(id))

	if err != nil {
		return nil, err
	}

	repository = &entity.EntityRepository{
		ID:         int(qRepo.ID),
		Repository: qRepo.Repository,
		Active:     qRepo.Active,
	}

	return repository, err
}

func (r *RepositoryRepository) CreateRepository(repository *entity.EntityRepository) error {
	context := context.Background()

	if err := r.checkExistsRepo(repository.Repository); err != nil {
		return err
	}

	err := r.queries.CreateRepository(context, queries.CreateRepositoryParams{
		Repository: repository.Repository,
		Active:     repository.Active,
	})

	return err
}

func (r *RepositoryRepository) UpdateRepository(repository *entity.EntityRepository) error {

	context := context.Background()

	if err := r.checkExistsRepo(repository.Repository); err != nil {
		return err
	}

	err := r.queries.UpdateRepository(context, queries.UpdateRepositoryParams{
		Repository: repository.Repository,
		Active:     repository.Active,
		ID:         int64(repository.ID),
	})

	return err
}

func (r *RepositoryRepository) DeleteRepository(repository *entity.EntityRepository) error {

	context := context.Background()

	if err := r.checkExistsRepo(repository.Repository); err != nil {
		return err
	}

	err := r.queries.DeleteRepository(context, int64(repository.ID))

	return err
}

func (r *RepositoryRepository) GetRepositories() (repositories []entity.EntityRepository, err error) {

	context := context.Background()

	results, err := r.queries.GetRepositories(context)

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		repositories = append(repositories, entity.EntityRepository{
			ID:         int(result.ID),
			Repository: result.Repository,
			Active:     result.Active,
		})
	}

	return repositories, err
}

func (r *RepositoryRepository) CreatePullRequest(pullRequest *entity.EntityPullRequest) error {
	context := context.Background()

	counter, _ := r.queries.CheckPullRequestExists(context, queries.CheckPullRequestExistsParams{
		Number:       int64(pullRequest.Number),
		RepositoryID: int64(pullRequest.RepositoryID),
	})

	if counter > 0 {
		err := r.queries.UpdatePullRequest(context, queries.UpdatePullRequestParams{
			Number:       int64(pullRequest.Number),
			Action:       pullRequest.Action,
			Status:       pullRequest.Status,
			RepositoryID: int64(pullRequest.RepositoryID),
			Url:          pullRequest.URL,
			Title:        pullRequest.Title,
			UpdatedAt:    time.Now(),
			ClosedAt:     sql.NullTime{Time: pullRequest.ClosedAt},
			Additions:    int32(pullRequest.Additions),
			Deletions:    int32(pullRequest.Deletions),
			ChangedFiles: int32(pullRequest.ChangedFiles),
			Commits:      int32(pullRequest.Commits),
		})

		return err
	} else {
		err := r.queries.CreatePullRequest(context, queries.CreatePullRequestParams{
			Number:       int64(pullRequest.Number),
			Action:       pullRequest.Action,
			Status:       pullRequest.Status,
			RepositoryID: int64(pullRequest.RepositoryID),
			Url:          pullRequest.URL,
			Title:        pullRequest.Title,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			ClosedAt:     sql.NullTime{Time: pullRequest.ClosedAt},
			Additions:    int32(pullRequest.Additions),
			Deletions:    int32(pullRequest.Deletions),
			ChangedFiles: int32(pullRequest.ChangedFiles),
			Commits:      int32(pullRequest.Commits),
		})

		return err
	}
}

func (r *RepositoryRepository) checkExistsRepo(repo string) error {
	context := context.Background()

	counter, _ := r.queries.CheckRepositoryExists(context, repo)

	if counter == 0 {
		return errors.New("Repository not found")
	}

	return nil
}
