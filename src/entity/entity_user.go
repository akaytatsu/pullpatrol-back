package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EntityUser struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=2,max=50" gorm:"not null"`
	Email    string `json:"email"      validate:"required,email" gorm:"unique;not null"`
	Password string `json:"password"   validate:"required,min=8,max=120" gorm:"not null"`
	IsAdmin  bool   `json:"is_admin"  validate:"required" gorm:"default:false"`
}

func NewUser(userParam EntityUser) (*EntityUser, error) {

	now := time.Now()

	password, err := GeneratePassword(userParam.Password)

	if err != nil {
		return nil, err
	}

	userParam.Password = password
	userParam.CreatedAt = now
	userParam.UpdatedAt = now

	u := &userParam

	return u, nil
}

func (u *EntityUser) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))

	if err != nil {
		return err
	}

	return nil
}

func (u *EntityUser) Validate() error {
	return validate.Struct(u)
}

func (u *EntityUser) GetValidated() error {
	err := u.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (u *EntityUser) UpdatePassword(newPassword string) error {
	hash, err := GeneratePassword(newPassword)
	if err != nil {
		return err
	}

	u.Password = hash

	return nil
}

func GeneratePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
