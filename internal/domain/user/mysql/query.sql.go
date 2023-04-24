// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package user

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users(username, password, salt)
VALUES (?, ?, ?)
`

type CreateUserParams struct {
	Username string
	Password string
	Salt     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.Username, arg.Password, arg.Salt)
}

const getUser = `-- name: GetUser :one
SELECT id, username, password, salt
FROM users
WHERE username = ?
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Salt,
	)
	return i, err
}