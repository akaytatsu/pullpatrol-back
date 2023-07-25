// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: users.sql

package queries

import (
	"context"
	"database/sql"
)

const checkUserByEmail = `-- name: CheckUserByEmail :one
select count(*) from users where email = $1
`

func (q *Queries) CheckUserByEmail(ctx context.Context, email string) (int64, error) {
	row := q.queryRow(ctx, q.checkUserByEmailStmt, checkUserByEmail, email)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const checkUserByID = `-- name: CheckUserByID :one
select count(*) from users where id = $1
`

func (q *Queries) CheckUserByID(ctx context.Context, id int64) (int64, error) {
	row := q.queryRow(ctx, q.checkUserByIDStmt, checkUserByID, id)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :exec
insert into users(name, email, password, is_admin) values ($1, $2, $3, $4)
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.exec(ctx, q.createUserStmt, createUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.IsAdmin,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
delete from users where id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, password, is_admin, git_name, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.IsAdmin,
		&i.GitName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, is_admin, git_name, created_at, updated_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.IsAdmin,
		&i.GitName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
update users set name = $1, email = $2, is_admin = $3, git_name = $4 where id = $5
`

type UpdateUserParams struct {
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	IsAdmin bool           `json:"is_admin"`
	GitName sql.NullString `json:"git_name"`
	ID      int64          `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.exec(ctx, q.updateUserStmt, updateUser,
		arg.Name,
		arg.Email,
		arg.IsAdmin,
		arg.GitName,
		arg.ID,
	)
	return err
}
