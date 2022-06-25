// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.checkEmailStmt, err = db.PrepareContext(ctx, checkEmail); err != nil {
		return nil, fmt.Errorf("error preparing query CheckEmail: %w", err)
	}
	if q.registerContestantStmt, err = db.PrepareContext(ctx, registerContestant); err != nil {
		return nil, fmt.Errorf("error preparing query RegisterContestant: %w", err)
	}
	if q.registerUserStmt, err = db.PrepareContext(ctx, registerUser); err != nil {
		return nil, fmt.Errorf("error preparing query RegisterUser: %w", err)
	}
	if q.registerVoterStmt, err = db.PrepareContext(ctx, registerVoter); err != nil {
		return nil, fmt.Errorf("error preparing query RegisterVoter: %w", err)
	}
	if q.seedAdminStmt, err = db.PrepareContext(ctx, seedAdmin); err != nil {
		return nil, fmt.Errorf("error preparing query SeedAdmin: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	if q.updateVoterStmt, err = db.PrepareContext(ctx, updateVoter); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateVoter: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.checkEmailStmt != nil {
		if cerr := q.checkEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkEmailStmt: %w", cerr)
		}
	}
	if q.registerContestantStmt != nil {
		if cerr := q.registerContestantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing registerContestantStmt: %w", cerr)
		}
	}
	if q.registerUserStmt != nil {
		if cerr := q.registerUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing registerUserStmt: %w", cerr)
		}
	}
	if q.registerVoterStmt != nil {
		if cerr := q.registerVoterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing registerVoterStmt: %w", cerr)
		}
	}
	if q.seedAdminStmt != nil {
		if cerr := q.seedAdminStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing seedAdminStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	if q.updateVoterStmt != nil {
		if cerr := q.updateVoterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateVoterStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                     DBTX
	tx                     *sql.Tx
	checkEmailStmt         *sql.Stmt
	registerContestantStmt *sql.Stmt
	registerUserStmt       *sql.Stmt
	registerVoterStmt      *sql.Stmt
	seedAdminStmt          *sql.Stmt
	updateUserStmt         *sql.Stmt
	updateVoterStmt        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                     tx,
		tx:                     tx,
		checkEmailStmt:         q.checkEmailStmt,
		registerContestantStmt: q.registerContestantStmt,
		registerUserStmt:       q.registerUserStmt,
		registerVoterStmt:      q.registerVoterStmt,
		seedAdminStmt:          q.seedAdminStmt,
		updateUserStmt:         q.updateUserStmt,
		updateVoterStmt:        q.updateVoterStmt,
	}
}
