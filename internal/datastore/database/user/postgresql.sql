-- name: CreateUser :execresult
INSERT INTO users(id, username, password, salt)
VALUES ($1, $2, $3, $4);

-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;

-- name: GetUserExistance :one
SELECT EXISTS(
    SELECT *
    FROM users
    WHERE username = $1
    LIMIT 1
);
