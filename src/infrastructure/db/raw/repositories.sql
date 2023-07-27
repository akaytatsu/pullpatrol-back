-- name: GetRepositoryByID :one
select * from repositories where id = $1;

-- name: GetRepositoryByRepository :one
select * from repositories where repository = $1;

-- name: CheckRepositoryExists :one
select count(*) from repositories where repository = $1;

-- name: CreateRepository :one
insert into repositories (repository, active, updated_at) values ($1, $2, $3) RETURNING *;

-- name: UpdateRepository :one
update repositories set repository = $1, active = $2, updated_at = $3 where id = $4 RETURNING *;

-- name: DeleteRepository :exec
delete from repositories where id = $1;

-- name: GetRepositories :many
select * from repositories order by id asc;
