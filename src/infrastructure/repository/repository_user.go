package repository

import (
	"app/entity"
	"app/infrastructure/db/queries"
	"context"
	"database/sql"
	"errors"
	"time"
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
		CreatedAt: qUser.CreatedAt.Time,
		UpdatedAt: qUser.UpdatedAt,
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
		CreatedAt: qUser.CreatedAt.Time,
		UpdatedAt: qUser.UpdatedAt,
	}

	return user, err
}

func (u *RepositoryUser) CreateUser(user *entity.EntityUser) error {

	if err := u.checkExistsByMail(user.Email); err == nil {
		return err
	}

	context := context.Background()

	data, err := u.queries.CreateUser(context, queries.CreateUserParams{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	user.ID = int(data.ID)
	user.IsAdmin = data.IsAdmin
	user.CreatedAt = data.CreatedAt.Time
	user.UpdatedAt = data.UpdatedAt
	user.Email = data.Email
	user.Name = data.Name

	return nil
}

func (u *RepositoryUser) UpdateUser(user *entity.EntityUser) error {
	context := context.Background()

	if err := u.checkExistsByMail(user.Email); err != nil {
		return err
	}

	_, err := u.queries.UpdateUser(context, queries.UpdateUserParams{
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

func (u *RepositoryUser) GetUsers() (users []entity.EntityUser, err error) {

	context := context.Background()

	users = make([]entity.EntityUser, 0)

	qUsers, err := u.queries.GetUsers(context)

	for _, qUser := range qUsers {
		user := entity.EntityUser{
			ID:        int(qUser.ID),
			Name:      qUser.Name,
			Email:     qUser.Email,
			IsAdmin:   qUser.IsAdmin,
			CreatedAt: qUser.CreatedAt.Time,
			UpdatedAt: qUser.UpdatedAt,
		}

		users = append(users, user)
	}

	return users, err
}

func (u *RepositoryUser) GetUser(id int) (user *entity.EntityUser, err error) {
	context := context.Background()

	qUser, err := u.queries.GetUser(context, int64(id))

	if err != nil {
		return nil, err
	}

	user = &entity.EntityUser{
		ID:        int(qUser.ID),
		Name:      qUser.Name,
		Email:     qUser.Email,
		CreatedAt: qUser.CreatedAt.Time,
		UpdatedAt: qUser.UpdatedAt,
	}

	return user, err
}

func (u *RepositoryUser) GetGroups() (groups []entity.EntityGroup, err error) {

	context := context.Background()

	groups = make([]entity.EntityGroup, 0)

	qGroups, err := u.queries.GetGroups(context)

	for _, qGroup := range qGroups {
		group := entity.EntityGroup{
			ID:        int(qGroup.ID),
			Name:      qGroup.Name,
			CreatedAt: qGroup.CreatedAt.Time,
			UpdatedAt: qGroup.UpdatedAt,
		}

		groups = append(groups, group)
	}

	return groups, err
}

func (u *RepositoryUser) GetUsersByGroup(groupID int) (users []entity.EntityUser, err error) {

	context := context.Background()

	users = make([]entity.EntityUser, 0)

	qUsers, err := u.queries.GetUsersByGroup(context, int64(groupID))

	for _, qUser := range qUsers {
		user := entity.EntityUser{
			ID:        int(qUser.ID),
			Name:      qUser.Name,
			Email:     qUser.Email,
			IsAdmin:   qUser.IsAdmin,
			CreatedAt: qUser.CreatedAt.Time,
			UpdatedAt: qUser.UpdatedAt,
		}

		users = append(users, user)
	}

	return users, err
}

func (u *RepositoryUser) GetGroup(groupID int) (group *entity.EntityGroup, err error) {
	context := context.Background()

	qGroup, err := u.queries.GetGroup(context, int64(groupID))

	if err != nil {
		return nil, err
	}

	group = &entity.EntityGroup{
		ID:        int(qGroup.ID),
		Name:      qGroup.Name,
		CreatedAt: qGroup.CreatedAt.Time,
		UpdatedAt: qGroup.UpdatedAt,
	}

	return group, err
}

func (u *RepositoryUser) CreateGroup(group *entity.EntityGroup) error {

	context := context.Background()

	data, err := u.queries.CreateGroup(context, queries.CreateGroupParams{
		Name:      group.Name,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	group.ID = int(data.ID)
	group.CreatedAt = data.CreatedAt.Time
	group.UpdatedAt = data.UpdatedAt

	return nil
}

func (u *RepositoryUser) UpdateGroup(group *entity.EntityGroup) error {
	context := context.Background()

	_, err := u.queries.UpdateGroup(context, queries.UpdateGroupParams{
		Name:      group.Name,
		UpdatedAt: time.Now(),
		ID:        int64(group.ID),
	})

	return err
}

func (u *RepositoryUser) DeleteGroup(group *entity.EntityGroup) error {

	context := context.Background()

	err := u.queries.DeleteGroup(context, int64(group.ID))

	return err
}

func (u *RepositoryUser) AddUserToGroup(userID, groupID int) error {
	context := context.Background()

	_, err := u.queries.AddUserToGroup(context, queries.AddUserToGroupParams{
		UserID:  int64(userID),
		GroupID: int64(groupID),
	})

	return err
}

func (u *RepositoryUser) RemoveUserFromGroup(userID, groupID int) error {

	context := context.Background()

	err := u.queries.RemoveUserFromGroup(context, queries.RemoveUserFromGroupParams{
		UserID:  int64(userID),
		GroupID: int64(groupID),
	})

	return err
}
