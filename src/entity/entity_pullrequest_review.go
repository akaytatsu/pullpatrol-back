package entity

import "time"

type EntityPullRequestReview struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	PullRequestID     int
	PullRequestRoleID int
	UserID            int
	Status            string
	Comment           string
}
