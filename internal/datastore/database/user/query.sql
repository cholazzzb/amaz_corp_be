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

-- name: CreateFriend :execresult
INSERT INTO friends(member1_id, member2_id)
VALUES (?, ?);

-- name: GetFriendsByMemberId :many
SELECT m.id, m.name, m.status
FROM members m
JOIN friends f ON (m.id = f.member1_id OR m.id = f.member2_id)
WHERE (f.member1_id = ? OR f.member2_id = ?) AND m.id != ?
LIMIT 10;