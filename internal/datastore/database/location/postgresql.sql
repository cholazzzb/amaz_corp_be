-- name: GetAllBuildings :many
SELECT *
FROM buildings
LIMIT 10;

-- name: GetBuildingsByMemberId :many
SELECT b.id, b.name
FROM buildings b
JOIN members_buildings mb ON (b.id = mb.building_id)
WHERE mb.member_id = $1
LIMIT 10;

-- name: GetMembersByRoomId :many
SELECT m.name, m.status, m.user_id
FROM members m
WHERE m.room_id = $1
LIMIT 10;

-- name: GetRoomsByMemberId :many
SELECT r.id, r.name
FROM rooms r
JOIN members_buildings mb ON (mb.building_id = r.building_id)
WHERE mb.member_id = $1
LIMIT 10;

-- name: GetRoomsByBuildingId :many
SELECT r.id, r.name
FROM rooms r
WHERE r.building_id = $1
LIMIT 10;

-- name: GetMemberBuildingById :one
SELECT EXISTS(SELECT mb.member_id, mb.building_id 
FROM members_buildings mb
WHERE mb.member_id = $1 AND mb.building_id = $2
LIMIT 1);

-- name: CreateMemberBuilding :execresult
INSERT INTO members_buildings(member_id, building_id)
VALUES ($1, $2);

-- name: DeleteMemberBuilding :exec
DELETE FROM members_buildings
WHERE member_id = $1 AND building_id = $2;