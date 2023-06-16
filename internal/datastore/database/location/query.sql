-- name: GetAllBuildings :many
SELECT *
FROM buildings
LIMIT 10;

-- name: GetBuildingsByMemberId :many
SELECT b.id, b.name
FROM buildings b
JOIN members_buildings mb ON (b.id = mb.member_id)
WHERE mb.member_id = ?
LIMIT 10;

-- name: GetMembersByRoomId :many
SELECT m.name, m.status, m.user_id
FROM members m
WHERE m.room_id = ?
LIMIT 10;

-- name: GetRoomsByMemberId :many
SELECT r.id, r.name
FROM rooms r
JOIN members_buildings mb ON (mb.building_id = r.building_id)
WHERE mb.member_id = ?
LIMIT 10;

-- name: GetRoomsByBuildingId :many
Select r.id, r.name
FROM  rooms r
WHERE r.building_id = ?
LIMIT 10;

-- name: CreateMemberBuilding :execresult
INSERT INTO members_buildings(member_id, building_id)
VALUES (?, ?);