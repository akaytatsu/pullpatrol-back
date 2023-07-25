-- name: GetRepositoryByID :one
select * from repositories where id = $1;

-- name: GetRepositoryByRepository :one
select * from repositories where repository = $1;

-- name: CheckRepositoryExists :one
select count(*) from repositories where repository = $1;

-- name: CreateRepository :exec
insert into repositories (repository, active) values ($1, $2);

-- name: UpdateRepository :exec
update repositories set repository = $1, active = $2 where id = $3;

-- name: DeleteRepository :exec
delete from repositories where id = $1;

-- name: GetRepositories :many
select * from repositories order by id asc;