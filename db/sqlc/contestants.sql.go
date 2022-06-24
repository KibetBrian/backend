// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: contestants.sql

package db

import (
	"context"
)

const registerContestant = `-- name: RegisterContestant :one
INSERT INTO contestants(full_name,email,password,position, description)
VALUES ($1,$2,$3,$4, $5) RETURNING id, full_name, email, password, position, registered_at, description
`

type RegisterContestantParams struct {
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Position    string `json:"position"`
	Description string `json:"description"`
}

func (q *Queries) RegisterContestant(ctx context.Context, arg RegisterContestantParams) (Contestant, error) {
	row := q.queryRow(ctx, q.registerContestantStmt, registerContestant,
		arg.FullName,
		arg.Email,
		arg.Password,
		arg.Position,
		arg.Description,
	)
	var i Contestant
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.Position,
		&i.RegisteredAt,
		&i.Description,
	)
	return i, err
}
