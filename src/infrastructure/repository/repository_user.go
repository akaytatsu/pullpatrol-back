package repository

import (
	"app/entity"
	"app/infrastructure/db/queries"
	"context"
	"database/sql"
	"errors"
)

type RepositoryUser struct {
	db      *sql.DB
	queries *queries.Queries
}

func NewRepositoryUser(db *sql.DB) *RepositoryUser {
	return &RepositoryUser{db: db, queries: queries.New(db)}
}

func (u *RepositoryUser) GetByID(id int) (user *entity.EntityUser, err error) {

	ctx := context.Background()

	qUser, err := u.queries.GetUser(ctx, int64(id))

	if err != nil {
		return nil, err
	}

	user = &entity.EntityUser{
		ID:        int(qUser.ID),
		Name:      qUser.Name,
		Email:     qUser.Email,
		Password:  qUser.Password,
		CreatedAt: qUser.CreatedAt,
		UpdatedAt: qUser.CreatedAt,
	}

	return user, err
}

func (u *RepositoryUser) GetByMail(email string) (user *entity.EntityUser, err error) {

	context := context.Background()

	qUser, err := u.queries.GetUserByEmail(context, email)

	user = &entity.EntityUser{
		ID:        int(qUser.ID),
		Name:      qUser.Name,
		Email:     qUser.Email,
		Password:  qUser.Password,
		CreatedAt: qUser.CreatedAt,
		UpdatedAt: qUser.UpdatedAt,
	}

	return user, err
}

func (u *RepositoryUser) CreateUser(user *entity.EntityUser) error {

	if err := u.checkExistsByMail(user.Email); err != nil {
		return err
	}

	context := context.Background()

	err := u.queries.CreateUser(context, queries.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	})

	return err
}

func (u *RepositoryUser) UpdateUser(user *entity.EntityUser) error {
	context := context.Background()

	if err := u.checkExistsByMail(user.Email); err != nil {
		return err
	}

	err := u.queries.UpdateUser(context, queries.UpdateUserParams{
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
		ID:      int64(user.ID),
	})

	return err
}

func (u *RepositoryUser) DeleteUser(user *entity.EntityUser) error {

	if err := u.checkExistsByMail(user.Email); err != nil {
		return err
	}

	context := context.Background()

	err := u.queries.DeleteUser(context, int64(user.ID))

	return err
}

func (u *RepositoryUser) checkExistsByMail(email string) error {
	context := context.Background()

	exists, _ := u.queries.CheckUserByEmail(context, email)

	if exists == 0 {
		return errors.New("user not found")
	}

	return nil
}
