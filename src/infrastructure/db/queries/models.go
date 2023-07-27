// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package queries

import (
	"database/sql"
	"time"
)

type Group struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GroupUser struct {
	ID        int64     `json:"id"`
	GroupID   int64     `json:"group_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Pullrequest struct {
	ID           int64        `json:"id"`
	Number       int64        `json:"number"`
	Action       string       `json:"action"`
	RepositoryID int64        `json:"repository_id"`
	Status       string       `json:"status"`
	Url          string       `json:"url"`
	Title        string       `json:"title"`
	ClosedAt     sql.NullTime `json:"closed_at"`
	Additions    int32        `json:"additions"`
	Deletions    int32        `json:"deletions"`
	ChangedFiles int32        `json:"changed_files"`
	Commits      int32        `json:"commits"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type PullrequestReview struct {
	ID                int64     `json:"id"`
	PullrequestID     int64     `json:"pullrequest_id"`
	PullrequestRoleID int64     `json:"pullrequest_role_id"`
	UserID            int64     `json:"user_id"`
	Status            string    `json:"status"`
	Comment           string    `json:"comment"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type PullrequestRole struct {
	ID           int64     `json:"id"`
	RepositoryID int64     `json:"repository_id"`
	RoleType     string    `json:"role_type"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Repository struct {
	ID         int64     `json:"id"`
	Repository string    `json:"repository"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type User struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	IsAdmin   bool           `json:"is_admin"`
	GitName   sql.NullString `json:"git_name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
