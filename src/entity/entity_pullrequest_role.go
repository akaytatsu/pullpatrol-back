package entity

import "time"

type EntityPullRequestRole struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	PullRequestID int
	RoleType      string
	Description   string
}
