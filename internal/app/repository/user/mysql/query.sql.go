// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package user

import (
	"context"
	"database/sql"
)

const createFriend = `-- name: CreateFriend :execresult
INSERT INTO friends(member1_id, member2_id)
VALUES (?, ?)
`

type CreateFriendParams struct {
	Member1ID int64
	Member2ID int64
}

func (q *Queries) CreateFriend(ctx context.Context, arg CreateFriendParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createFriend, arg.Member1ID, arg.Member2ID)
}

const createMember = `-- name: CreateMember :execresult
INSERT INTO members(name, status, user_id)
VALUES (?, ?, ?)
`

type CreateMemberParams struct {
	Name   string
	Status string
	UserID int64
}

func (q *Queries) CreateMember(ctx context.Context, arg CreateMemberParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createMember, arg.Name, arg.Status, arg.UserID)
}

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

const getFriendsByMemberId = `-- name: GetFriendsByMemberId :many
SELECT m.id, m.name, m.status
FROM members m
JOIN friends f ON (m.id = f.member1_id OR m.id = f.member2_id)
WHERE (f.member1_id = ? OR f.member2_id = ?) AND m.id != ?
LIMIT 10
`

type GetFriendsByMemberIdParams struct {
	Member1ID int64
	Member2ID int64
	ID        int64
}

type GetFriendsByMemberIdRow struct {
	ID     int64
	Name   string
	Status string
}

func (q *Queries) GetFriendsByMemberId(ctx context.Context, arg GetFriendsByMemberIdParams) ([]GetFriendsByMemberIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getFriendsByMemberId, arg.Member1ID, arg.Member2ID, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFriendsByMemberIdRow
	for rows.Next() {
		var i GetFriendsByMemberIdRow
		if err := rows.Scan(&i.ID, &i.Name, &i.Status); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMemberByName = `-- name: GetMemberByName :one
SELECT id, user_id, name, status, room_id
FROM members
WHERE name = ?
LIMIT 1
`

func (q *Queries) GetMemberByName(ctx context.Context, name string) (Member, error) {
	row := q.db.QueryRowContext(ctx, getMemberByName, name)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Status,
		&i.RoomID,
	)
	return i, err
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
