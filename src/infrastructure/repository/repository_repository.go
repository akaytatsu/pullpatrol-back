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

func (r *RepositoryRepository) GetByRepo(repo string) (repository *entity.EntityRepository, err error) {

	context := context.Background()

	qRepo, err := r.queries.GetRepositoryByRepository(context, repo)

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

func (r *RepositoryRepository) CreateOrUpdateRepository(repository *entity.EntityRepository) error {
	context := context.Background()

	check := r.CheckExistsRepo(repository.Repository)

	var data queries.Repository
	var err error

	if check {
		data, err = r.queries.UpdateRepository(context, queries.UpdateRepositoryParams{
			Repository: repository.Repository,
			Active:     true,
			ID:         int64(repository.ID),
			UpdatedAt:  time.Now(),
		})

		return err
	} else {
		data, err = r.queries.CreateRepository(context, queries.CreateRepositoryParams{
			Repository: repository.Repository,
			Active:     repository.Active,
			UpdatedAt:  time.Now(),
		})
	}

	if err != nil {
		return err
	}

	repository.ID = int(data.ID)
	repository.Repository = data.Repository
	repository.Active = data.Active

	return nil
}

func (r *RepositoryRepository) DeleteRepository(repository *entity.EntityRepository) error {

	context := context.Background()

	if check := r.CheckExistsRepo(repository.Repository); !check {
		return errors.New("repository not found")
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

func (r *RepositoryRepository) CreateOrUpdatePullRequest(pullRequest *entity.EntityPullRequest) error {
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

func (r *RepositoryRepository) CheckExistsRepo(repo string) bool {
	context := context.Background()

	counter, _ := r.queries.CheckRepositoryExists(context, repo)

	return counter != 0
}

func (r *RepositoryRepository) GetPullRequestRoles(repositoryID int) (pullRequestRoles []entity.EntityPullRequestRole, err error) {
	context := context.Background()

	results, err := r.queries.GetPullRequestRoles(context, int64(repositoryID))

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		pullRequestRoles = append(pullRequestRoles, entity.EntityPullRequestRole{
			ID: int(result.ID),
		})
	}

	return pullRequestRoles, err
}

func (r *RepositoryRepository) GetPullRequestRole(pullRequestRoleID int) (pullRequestRole *entity.EntityPullRequestRole, err error) {
	context := context.Background()

	result, err := r.queries.GetPullRequestRole(context, int64(pullRequestRoleID))

	if err != nil {
		return nil, err
	}

	pullRequestRole = &entity.EntityPullRequestRole{
		ID: int(result.ID),
	}

	return pullRequestRole, err
}

func (r *RepositoryRepository) CreatePullRequestRole(repositoryID int, pullRequestRole *entity.EntityPullRequestRole) error {
	context := context.Background()

	model, err := r.queries.CreatePullRequestRole(context, queries.CreatePullRequestRoleParams{
		RepositoryID: int64(repositoryID),
		RoleType:     pullRequestRole.RoleType,
		Description:  pullRequestRole.Description,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	pullRequestRole.ID = int(model.ID)

	return err
}

func (r *RepositoryRepository) UpdatePullRequestRole(pullRequestRoleID int, pullRequestRole *entity.EntityPullRequestRole) error {
	context := context.Background()

	err := r.queries.UpdatePullRequestRole(context, queries.UpdatePullRequestRoleParams{
		ID:          int64(pullRequestRoleID),
		RoleType:    pullRequestRole.RoleType,
		Description: pullRequestRole.Description,
		UpdatedAt:   time.Now(),
	})

	return err
}

func (r *RepositoryRepository) DeletePullRequestRole(pullRequestRoleID int) error {
	context := context.Background()

	err := r.queries.DeletePullRequestRole(context, int64(pullRequestRoleID))

	return err
}

func (r *RepositoryRepository) GetPullRequestReviews(pullRequestID int) (pullRequestReviews []entity.EntityPullRequestReview, err error) {
	context := context.Background()

	results, err := r.queries.GetPullRequestReviews(context, int64(pullRequestID))

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		pullRequestReviews = append(pullRequestReviews, entity.EntityPullRequestReview{
			ID: int(result.ID),
		})
	}

	return pullRequestReviews, err
}

func (r *RepositoryRepository) GetPullRequestReview(pullRequestReviewID int) (pullRequestReview *entity.EntityPullRequestReview, err error) {
	context := context.Background()

	result, err := r.queries.GetPullRequestReview(context, int64(pullRequestReviewID))

	if err != nil {
		return nil, err
	}

	pullRequestReview = &entity.EntityPullRequestReview{
		ID: int(result.ID),
	}

	return pullRequestReview, err
}

func (r *RepositoryRepository) CreatePullRequestReview(pullRequestID int, pullRequestReview *entity.EntityPullRequestReview) error {
	context := context.Background()

	model, err := r.queries.CreatePullRequestReview(context, queries.CreatePullRequestReviewParams{
		PullrequestID:     int64(pullRequestID),
		PullrequestRoleID: int64(pullRequestReview.PullRequestRoleID),
		UserID:            int64(pullRequestReview.UserID),
		Status:            pullRequestReview.Status,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	})

	pullRequestReview.ID = int(model.ID)

	return err
}

func (r *RepositoryRepository) UpdatePullRequestReview(pullRequestReviewID int, pullRequestReview *entity.EntityPullRequestReview) error {
	context := context.Background()

	err := r.queries.UpdatePullRequestReview(context, queries.UpdatePullRequestReviewParams{
		ID:        int64(pullRequestReviewID),
		Comment:   pullRequestReview.Comment,
		Status:    pullRequestReview.Status,
		UpdatedAt: time.Now(),
	})

	return err
}

func (r *RepositoryRepository) DeletePullRequestReview(pullRequestReviewID int) error {
	context := context.Background()

	err := r.queries.DeletePullRequestReview(context, int64(pullRequestReviewID))

	return err
}
