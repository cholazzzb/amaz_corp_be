-- name: CreateMember :execresult
INSERT INTO members(name, status, user_id)
VALUES (?, ?, ?);
-- name: GetMemberByName :one
SELECT *
FROM members
WHERE name = ?
LIMIT 1;