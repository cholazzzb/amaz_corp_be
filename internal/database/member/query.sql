-- name: CreateMember :execresult
INSERT INTO members(name)
VALUES (?);
-- name: GetMember :one
SELECT *
FROM members
WHERE name = ?
LIMIT 1;