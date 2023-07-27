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

-- name: GetUsersByGroup :many
select u.* from users u inner join group_user gu on u.id = gu.user_id where gu.group_id = $1 order by u.id asc;

-- name: GetGroups :many
select * from groups order by id asc;

-- name: GetGroup :one
select * from groups where id = $1 LIMIT 1;

-- name: CreateGroup :one
insert into groups(name, description, updated_at) values ($1, $2, $3) RETURNING *;

-- name: UpdateGroup :one
update groups set name = $1, description = $2, updated_at = $3 where id = $4 RETURNING *;

-- name: DeleteGroup :exec
delete from groups where id = $1;

-- name: CheckGroupExists :one
select count(*) from groups where id = $1;

-- name: AddUserToGroup :one
insert into group_user(group_id, user_id, updated_at) values ($1, $2, $3) RETURNING *;

-- name: RemoveUserFromGroup :exec
delete from group_user where group_id = $1 and user_id = $2;