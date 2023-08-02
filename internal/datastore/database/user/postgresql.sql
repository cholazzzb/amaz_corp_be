-- name: CreateUser :execresult
INSERT INTO users(id, username, password, salt)
VALUES ($1, $2, $3, $4);

-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;

-- name: GetUserExistance :one
SELECT EXISTS(SELECT *
FROM users
WHERE username = $1
LIMIT 1);

-- name: CreateMember :execresult
INSERT INTO members(id, name, status, user_id)
VALUES ($1, $2, $3, $4);

-- name: GetMemberByName :one
SELECT *
FROM members
WHERE name = $1
LIMIT 1;

-- name: CreateFriend :execresult
INSERT INTO friends(member1_id, member2_id)
VALUES ($1, $2);

-- name: GetFriendsByMemberId :many
SELECT m.id, m.name, m.status
FROM members m
JOIN friends f ON (m.id = f.member1_id OR m.id = f.member2_id)
WHERE (f.member1_id = $1 OR f.member2_id = $1) AND m.id != $1
LIMIT 10;