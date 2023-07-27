package entity

import "time"

type EntityPullRequestRole struct {
	ID            int
	PullRequestID int
	RoleType      string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
