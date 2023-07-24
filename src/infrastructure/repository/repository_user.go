package repository

import (
	"app/entity"
	"app/prisma/db"
	"context"
)

type RepositoryUser struct {
	db *db.PrismaClient
}

func NewRepositoryUser(db *db.PrismaClient) *RepositoryUser {
	return &RepositoryUser{db: db}
}

func (u *RepositoryUser) GetByID(id int) (user *entity.EntityUser, err error) {

	context := context.Background()

	model, err := u.db.User.FindFirst(
		db.User.ID.Equals(id),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	user = &entity.EntityUser{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	return user, err
}

func (u *RepositoryUser) GetByMail(email string) (user *entity.EntityUser, err error) {

	context := context.Background()

	model, err := u.db.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(context)

	if err != nil {
		return nil, err
	}

	user = &entity.EntityUser{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	return user, err
}

func (u *RepositoryUser) CreateUser(user *entity.EntityUser) error {

	_, err := u.GetByMail(user.Email)

	if err == nil {
		return err
	}

	context := context.Background()

	_, err = u.db.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Name.Set(user.Name),
		db.User.Password.Set(user.Password),
		db.User.IsAdmin.Set(user.IsAdmin),
	).Exec(context)

	return err
}

func (u *RepositoryUser) UpdateUser(user *entity.EntityUser) error {

	_, err := u.GetByMail(user.Email)

	if err != nil {
		return err
	}

	context := context.Background()

	_, err = u.db.User.FindUnique(
		db.User.ID.Equals(user.ID),
	).Update(
		db.User.Name.Set(user.Name),
	).Exec(context)

	return err
}

func (u *RepositoryUser) DeleteUser(user *entity.EntityUser) error {

	_, err := u.GetByMail(user.Email)

	if err != nil {
		return err
	}

	context := context.Background()

	_, err = u.db.User.FindUnique(
		db.User.ID.Equals(user.ID),
	).Delete().Exec(context)

	return err
}
