package entity

import "gorm.io/gorm"

type EntityPullRequestReview struct {
	gorm.Model
	PullRequestID     int
	PullRequestRoleID int
	UserID            int
	Status            string
	Comment           string
}
