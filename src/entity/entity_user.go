package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type EntityUser struct {
	ID        int
	FirstName string    `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string    `json:"last_name"  validate:"required,min=2,max=50"`
	Email     string    `json:"email"      validate:"required,email"`
	Password  string    `json:"password"   validate:"required,min=8,max=120"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(userParam EntityUser) (*EntityUser, error) {

	now := time.Now()

	password, err := GeneratePassword(userParam.Password)

	if err != nil {
		return nil, err
	}

	u := &EntityUser{
		FirstName: userParam.FirstName,
		LastName:  userParam.LastName,
		Email:     userParam.Email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}

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

	pwd, err := GeneratePassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pwd

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
