package repository

import (
	"app/entity"
	"errors"
	"time"

	"gorm.io/gorm"
)

type RepositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *RepositoryUser {
	return &RepositoryUser{db: db}
}

func (u *RepositoryUser) GetByID(id int) (user *entity.EntityUser, err error) {

	u.db.First(&user, id)

	return user, err
}

func (u *RepositoryUser) GetByMail(email string) (user *entity.EntityUser, err error) {

	err = u.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (u *RepositoryUser) CreateUser(user *entity.EntityUser) error {

	if err := u.checkExistsByMail(user.Email); err == nil {
		return err
	}

	err := u.db.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *RepositoryUser) UpdateUser(user *entity.EntityUser) error {
	if err := u.checkExistsByMail(user.Email); err != nil {
		return err
	}

	err := u.db.Where("id = ?", user.ID).Save(&user).Error

	return err
}

func (u *RepositoryUser) DeleteUser(user *entity.EntityUser) error {

	if err := u.checkExistsByMail(user.Email); err != nil {
		return err
	}

	err := u.db.Delete(user).Error

	return err
}

func (u *RepositoryUser) checkExistsByMail(email string) error {
	var exists bool

	u.db.Model(entity.EntityUser{}).Where("email = ?", email).Find(&exists)

	if !exists {
		return errors.New("user not found")
	}

	return nil
}

func (u *RepositoryUser) GetUsers() (users []entity.EntityUser, err error) {

	users = make([]entity.EntityUser, 0)

	err = u.db.Find(&users).Error

	return users, err
}

func (u *RepositoryUser) GetUser(id int) (user *entity.EntityUser, err error) {
	u.db.First(&user, id)

	return user, err
}

func (u *RepositoryUser) GetGroups() (groups []entity.EntityGroup, err error) {
	groups = make([]entity.EntityGroup, 0)

	u.db.Find(&groups)

	return groups, err
}

func (u *RepositoryUser) GetUsersByGroup(groupID int) (users []entity.EntityUser, err error) {
	users = make([]entity.EntityUser, 0)

	u.db.Joins("Group").Where("group.id = ?", groupID).Find(&users)

	return users, err
}

func (u *RepositoryUser) GetGroup(groupID int) (group *entity.EntityGroup, err error) {
	u.db.First(&group, groupID)

	return group, err
}

func (u *RepositoryUser) CreateGroup(group *entity.EntityGroup) error {
	return u.db.Create(&group).Error
}

func (u *RepositoryUser) UpdateGroup(group *entity.EntityGroup) error {
	return u.db.Where("id = ?", group.ID).Save(&group).Error
}

func (u *RepositoryUser) DeleteGroup(group *entity.EntityGroup) error {
	return u.db.Delete(&group, group.ID).Error
}

func (u *RepositoryUser) AddUserToGroup(userID, groupID int) error {
	return u.db.Model(&entity.EntityGroupUser{}).Create(&entity.EntityGroupUser{
		UserID:    userID,
		GroupID:   groupID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
}

func (u *RepositoryUser) RemoveUserFromGroup(userID, groupID int) error {

	err := u.db.Model(&entity.EntityGroupUser{}).Where("user_id = ? AND group_id = ?", userID, groupID).Delete(&entity.EntityGroupUser{}).Error

	return err
}
