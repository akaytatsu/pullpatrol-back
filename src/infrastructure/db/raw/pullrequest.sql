-- name: CheckPullRequestExists :one
select count(*) from pullrequest where number = $1 and repository_id = $2;

-- name: CreatePullRequest :exec
insert into pullrequest (number, action, repository_id, status, url, title, created_at, updated_at, closed_at, additions, deletions, changed_files, commits)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);

-- name: UpdatePullRequest :exec
update pullrequest set action = $1, status = $2, url = $3, title = $4, updated_at = $5, closed_at = $6, additions = $7, deletions = $8, changed_files = $9, commits = $10
where number = $11 and repository_id = $12;

-- name: GetPullRequest :one
select * from pullrequest where id = $1;

-- name: GetPullRequestRoles :many
select * from pullrequest_role where repository_id = $1 order by id asc;

-- name: GetPullRequestRole :one
select * from pullrequest_role where id = $1;

-- name: CreatePullRequestRole :one
insert into pullrequest_role (repository_id, role_type, description, created_at, updated_at)
values ($1, $2, $3, $4, $5) returning *;

-- name: UpdatePullRequestRole :exec
update pullrequest_role set role_type = $1, description = $2, updated_at = $3 where id = $4;

-- name: DeletePullRequestRole :exec
delete from pullrequest_role where id = $1;

-- name: GetPullRequestReviews :many
select * from pullrequest_review where pullrequest_id = $1 order by id asc;

-- name: GetPullRequestReview :one
select * from pullrequest_review where id = $1;

-- name: CreatePullRequestReview :one
insert into pullrequest_review (pullrequest_id, pullrequest_role_id, user_id, status, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6) returning *;

-- name: UpdatePullRequestReview :exec
update pullrequest_review set status = $1, comment = $2, updated_at = $3 where id = $4;

-- name: DeletePullRequestReview :exec
delete from pullrequest_review where id = $1;