-- name: CreateFriend :execresult
INSERT INTO friends(member1_id, member2_id)
VALUES (?, ?);
-- name: GetFriendsByMemberId :many
SELECT *
FROM friends
WHERE member1_id = ?
    OR member2_id = ?
LIMIT 10;