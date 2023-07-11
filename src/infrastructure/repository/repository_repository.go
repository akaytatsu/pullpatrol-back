package repository

import (
	"app/entity"

	"gorm.io/gorm"
)

type RepositoryRepository struct {
	db *gorm.DB
}

func NewRepositoryPostgres(db *gorm.DB) *RepositoryRepository {
	return &RepositoryRepository{db: db}
}

func (r *RepositoryRepository) GetByID(id string) (repository *entity.EntityRepository, err error) {
	err = r.db.Where("id = ?", id).First(&repository).Error

	return repository, err
}

func (r *RepositoryRepository) GetByName(name string) (repository *entity.EntityRepository, err error) {
	err = r.db.Where("name = ?", name).First(&repository).Error

	return repository, err
}

func (r *RepositoryRepository) CreateRepository(repository *entity.EntityRepository) error {
	return r.db.Create(&repository).Error
}

func (r *RepositoryRepository) UpdateRepository(repository *entity.EntityRepository) error {

	_, err := r.GetByName(repository.Name)

	if err != nil {
		return err
	}

	return r.db.Save(&repository).Error
}

func (r *RepositoryRepository) DeleteRepository(repository *entity.EntityRepository) error {

	_, err := r.GetByName(repository.Name)

	if err != nil {
		return err
	}

	return r.db.Delete(&repository).Error
}
