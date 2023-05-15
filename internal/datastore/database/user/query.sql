-- name: CreateUser :execresult
INSERT INTO users(username, password, salt)
VALUES (?, ?, ?);
-- name: GetUser :one
SELECT *
FROM users
WHERE username = ?
LIMIT 1;

-- name: CreateMember :execresult
INSERT INTO members(name, status, user_id)
VALUES (?, ?, ?);
-- name: GetMemberByName :one
SELECT *
FROM members
WHERE name = ?
LIMIT 1;