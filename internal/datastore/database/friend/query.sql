-- name: CreateFriend :execresult
INSERT INTO friends(member1_id, member2_id)
VALUES (?, ?);
-- name: GetFriendsByMemberId :many
SELECT m.id, m.name, m.status
FROM members m
JOIN friends f ON (m.id = f.member1_id OR m.id = f.member2_id)
WHERE (f.member1_id = ? OR f.member2_id = ?) AND m.id != ?
LIMIT 10;