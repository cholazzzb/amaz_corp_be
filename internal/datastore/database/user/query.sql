-- name: CreateUser :execresult
INSERT INTO users(username, password, salt)
VALUES (?, ?, ?);
-- name: GetUser :one
SELECT *
FROM users
WHERE username = ?
LIMIT 1;