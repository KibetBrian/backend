// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const checkEmail = `-- name: CheckEmail :one
SELECT email, count(*) FROM users WHERE email = $1 GROUP BY email
`

type CheckEmailRow struct {
	Email string `json:"email"`
	Count int64  `json:"count"`
}

func (q *Queries) CheckEmail(ctx context.Context, email string) (CheckEmailRow, error) {
	row := q.queryRow(ctx, q.checkEmailStmt, checkEmail, email)
	var i CheckEmailRow
	err := row.Scan(&i.Email, &i.Count)
	return i, err
}

const getTotalUsersNum = `-- name: GetTotalUsersNum :one
SELECT COUNT(email) FROM users
`

func (q *Queries) GetTotalUsersNum(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getTotalUsersNumStmt, getTotalUsersNum)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUser = `-- name: GetUser :one
SELECT id, full_name, email, password, is_admin FROM users WHERE id=$1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.IsAdmin,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, full_name, email, password, is_admin FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.IsAdmin,
	)
	return i, err
}

const registerUser = `-- name: RegisterUser :one
INSERT INTO users (full_name, email, password)
VALUES($1, $2, $3) RETURNING id, full_name, email, password, is_admin
`

type RegisterUserParams struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) RegisterUser(ctx context.Context, arg RegisterUserParams) (User, error) {
	row := q.queryRow(ctx, q.registerUserStmt, registerUser, arg.FullName, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.IsAdmin,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET email = $1, password = $2 WHERE email = $3 RETURNING id, full_name, email, password, is_admin
`

type UpdateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Email_2  string `json:"email2"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserStmt, updateUser, arg.Email, arg.Password, arg.Email_2)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.IsAdmin,
	)
	return i, err
}
