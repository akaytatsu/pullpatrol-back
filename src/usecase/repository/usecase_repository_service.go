package usecase_repository

import "app/entity"

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

	return u.repo.CreateRepository(repository)
}

func (u *UsecaseRepository) Update(repository *entity.EntityRepository) error {
	return u.repo.UpdateRepository(repository)
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
