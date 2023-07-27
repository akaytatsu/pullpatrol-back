package entity

import "time"

type EntityGroupUser struct {
	ID        int
	GroupID   int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
