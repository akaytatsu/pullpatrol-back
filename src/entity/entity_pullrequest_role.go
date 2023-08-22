package entity

import (
	"gorm.io/gorm"
)

type EntityPullRequestRole struct {
	gorm.Model
	PullRequestID int
	RoleType      string
	Description   string
}
