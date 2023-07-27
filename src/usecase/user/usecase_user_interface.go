package usecase_user

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_user.go -package=mocks app/usecase/user IRepositoryUser
type IRepositoryUser interface {
	GetByID(id int) (user *entity.EntityUser, err error)
	GetByMail(email string) (user *entity.EntityUser, err error)
	CreateUser(user *entity.EntityUser) error
	UpdateUser(user *entity.EntityUser) error
	DeleteUser(user *entity.EntityUser) error
	GetUsers() (users []entity.EntityUser, err error)
	GetUser(id int) (user *entity.EntityUser, err error)
	GetGroups() (groups []entity.EntityGroup, err error)
	GetUsersByGroup(groupID int) (users []entity.EntityUser, err error)
	GetGroup(groupID int) (group *entity.EntityGroup, err error)
	CreateGroup(group *entity.EntityGroup) error
	UpdateGroup(group *entity.EntityGroup) error
	DeleteGroup(group *entity.EntityGroup) error
	AddUserToGroup(userID, groupID int) error
	RemoveUserFromGroup(userID, groupID int) error
}

//go:generate mockgen -destination=../../mocks/mock_usecase_user.go -package=mocks app/usecase/user IUsecaseUser
type IUsecaseUser interface {
	LoginUser(email string, password string) (*entity.EntityUser, error)
	GetUserByToken(token string) (*entity.EntityUser, error)
	Create(user *entity.EntityUser) error
	Update(user *entity.EntityUser) error
	Delete(user *entity.EntityUser) error
	UpdatePassword(id int, oldPassword, newPassword, confirmPassword string) error
	GetUsers() (users []entity.EntityUser, err error)
	GetUser(id int) (user *entity.EntityUser, err error)
	GetGroups() (groups []entity.EntityGroup, err error)
	GetUsersByGroup(groupID int) (users []entity.EntityUser, err error)
	GetGroup(groupID int) (group *entity.EntityGroup, err error)
	CreateGroup(group *entity.EntityGroup) error
	UpdateGroup(group *entity.EntityGroup) error
	DeleteGroup(group *entity.EntityGroup) error
	AddUserToGroup(userID, groupID int) error
	RemoveUserFromGroup(userID, groupID int) error
}
