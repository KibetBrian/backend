// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleRegistration Role = "Registration"
	RoleSupport      Role = "Support"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type Admin struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     Role      `json:"role"`
}

type Contestant struct {
	ID           uuid.UUID    `json:"id"`
	FullName     string       `json:"fullName"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
	Position     string       `json:"position"`
	RegisteredAt sql.NullTime `json:"registeredAt"`
	Description  string       `json:"description"`
}

type User struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type Voter struct {
	ID                  uuid.UUID    `json:"id"`
	FullName            string       `json:"fullName"`
	Email               string       `json:"email"`
	RegisteredAt        time.Time    `json:"registeredAt"`
	VotedAt             sql.NullTime `json:"votedAt"`
	VotersPublicAddress string       `json:"votersPublicAddress"`
}
