package repository

import (
	"app/entity"
	"errors"

	"gorm.io/gorm"
)

type RepositoryRepository struct {
	db *gorm.DB
}

func NewRepositoryRepository(db *gorm.DB) *RepositoryRepository {
	return &RepositoryRepository{db: db}
}

func (r *RepositoryRepository) GetByID(id int) (repository *entity.EntityRepository, err error) {
	r.db.First(&repository, id)

	return repository, err
}

func (r *RepositoryRepository) GetByRepo(repo string) (repository *entity.EntityRepository, err error) {

	r.db.Where("repository = ?", repo).First(&repository)

	return repository, err
}

func (r *RepositoryRepository) CreateOrUpdateRepository(repository *entity.EntityRepository) error {

	check := r.CheckExistsRepo(repository.Repository)

	if check {
		return r.db.Model(&repository).Where("repository = ?", repository.Repository).Updates(&repository).Error
	}

	return r.db.Create(&repository).Error
}

func (r *RepositoryRepository) DeleteRepository(repository *entity.EntityRepository) error {

	if check := r.CheckExistsRepo(repository.Repository); !check {
		return errors.New("repository not found")
	}

	err := r.db.Delete(&repository).Error

	return err
}

func (r *RepositoryRepository) GetRepositories() (repositories []entity.EntityRepository, err error) {

	repositories = make([]entity.EntityRepository, 0)

	r.db.Find(&repositories)

	return repositories, err
}

func (r *RepositoryRepository) CreateOrUpdatePullRequest(pullRequest *entity.EntityPullRequest) error {

	var exists bool

	r.db.Model(entity.EntityPullRequest{}).Where("number = ? AND repository_id = ?", pullRequest.Number, pullRequest.RepositoryID).Find(&exists)

	if exists {
		return r.db.Model(&pullRequest).Where("number = ? AND repository_id = ?", pullRequest.Number, pullRequest.RepositoryID).Updates(&pullRequest).Error
	}

	return r.db.Create(&pullRequest).Error
}

func (r *RepositoryRepository) CheckExistsRepo(repo string) bool {
	var exists bool

	r.db.Model(entity.EntityRepository{}).Where("repository = ?", repo).Find(&exists)

	return exists
}

func (r *RepositoryRepository) GetPullRequestRoles(repositoryID int) (pullRequestRoles []entity.EntityPullRequestRole, err error) {

	pullRequestRoles = make([]entity.EntityPullRequestRole, 0)

	r.db.Where("repository_id = ?", repositoryID).Find(&pullRequestRoles)

	return pullRequestRoles, err
}

func (r *RepositoryRepository) GetPullRequestRole(pullRequestRoleID int) (pullRequestRole *entity.EntityPullRequestRole, err error) {

	r.db.First(&pullRequestRole, pullRequestRoleID)

	return pullRequestRole, err
}

func (r *RepositoryRepository) CreatePullRequestRole(repositoryID int, pullRequestRole *entity.EntityPullRequestRole) error {

	err := r.db.Create(&pullRequestRole).Error

	return err
}

func (r *RepositoryRepository) UpdatePullRequestRole(pullRequestRoleID int, pullRequestRole *entity.EntityPullRequestRole) error {

	err := r.db.Model(&pullRequestRole).Where("id = ?", pullRequestRoleID).Updates(&pullRequestRole).Error

	return err
}

func (r *RepositoryRepository) DeletePullRequestRole(pullRequestRoleID int) error {

	err := r.db.Delete(&entity.EntityPullRequestRole{}, pullRequestRoleID).Error

	return err
}

func (r *RepositoryRepository) GetPullRequestReviews(pullRequestID int) (pullRequestReviews []entity.EntityPullRequestReview, err error) {
	pullRequestReviews = make([]entity.EntityPullRequestReview, 0)

	r.db.Where("pullrequest_id = ?", pullRequestID).Find(&pullRequestReviews)

	return pullRequestReviews, err
}

func (r *RepositoryRepository) GetPullRequestReview(pullRequestReviewID int) (pullRequestReview *entity.EntityPullRequestReview, err error) {

	r.db.First(&pullRequestReview, pullRequestReviewID)

	return pullRequestReview, err
}

func (r *RepositoryRepository) CreatePullRequestReview(pullRequestID int, pullRequestReview *entity.EntityPullRequestReview) error {

	err := r.db.Create(&pullRequestReview).Error

	return err
}

func (r *RepositoryRepository) UpdatePullRequestReview(pullRequestReviewID int, pullRequestReview *entity.EntityPullRequestReview) error {

	err := r.db.Model(&pullRequestReview).Where("id = ?", pullRequestReviewID).Updates(&pullRequestReview).Error

	return err
}

func (r *RepositoryRepository) DeletePullRequestReview(pullRequestReviewID int) error {

	err := r.db.Delete(&entity.EntityPullRequestReview{}, pullRequestReviewID).Error

	return err
}
