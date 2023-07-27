package entity

import "time"

type EntityPullRequestReview struct {
	ID                int
	PullRequestID     int
	PullRequestRoleID int
	UserID            int
	Status            string
	Comment           string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
