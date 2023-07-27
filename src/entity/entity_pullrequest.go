package entity

import "time"

type EntityPullRequest struct {
	Number        int              `json:"number"`
	Action        string           `json:"action"`
	Status        string           `json:"status"`
	Repository    EntityRepository `json:"repository"`
	RepositoryID  int              `json:"repository_id"`
	RepositoryURL string           `json:"repository_url"`
	URL           string           `json:"url"`
	Title         string           `json:"title"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	ClosedAt      time.Time        `json:"closed_at"`
	Additions     int              `json:"additions"`
	Deletions     int              `json:"deletions"`
	ChangedFiles  int              `json:"changed_files"`
	Commits       int              `json:"commits"`
}
