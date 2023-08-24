package entity

import (
	"time"

	"gorm.io/gorm"
)

type EntityRepository struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Repository string         `json:"repository" validate:"required,min=2,max=80"`
	Label      string         `json:"label" validate:"required,min=2,max=80"`
	Driver     string         `json:"driver" validate:"required,min=2,max=80"`
	UserID     int            `json:"user_id"`
	Active     bool           `json:"active"`
	User       EntityUser     `gorm:"foreignKey:UserID" json:"user"`
}

func (e *EntityRepository) Validate() error {
	return validate.Struct(e)
}

func (e *EntityRepository) GetValidated() error {
	return e.Validate()
}
