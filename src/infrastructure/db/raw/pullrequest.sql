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