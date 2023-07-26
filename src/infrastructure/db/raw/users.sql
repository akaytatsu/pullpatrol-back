-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CheckUserByEmail :one
select count(*) from users where email = $1;

-- name: CheckUserByID :one
select count(*) from users where id = $1;

-- name: CreateUser :one
insert into users(name, email, password, is_admin, updated_at) values ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateUser :one
update users set name = $1, email = $2, is_admin = $3, git_name = $4, updated_at = $5 where id = $6 RETURNING *;

-- name: DeleteUser :exec
delete from users where id = $1;

-- name: GetUsers :many
select * from users order by id asc;