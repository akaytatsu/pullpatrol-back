-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CheckUserByEmail :one
select count(*) from users where email = $1;

-- name: CheckUserByID :one
select count(*) from users where id = $1;

-- name: CreateUser :exec
insert into users(name, email, password, is_admin) values ($1, $2, $3, $4);

-- name: UpdateUser :exec
update users set name = $1, email = $2, is_admin = $3, git_name = $4 where id = $5;

-- name: DeleteUser :exec
delete from users where id = $1;