-- name: CreateBuilding :execresult
INSERT INTO buildings(name, owner_id)
VALUES ($1, $2);

-- name: GetNumberOfBuildingOwned :one
SELECT COUNT(*)
FROM buildings
WHERE buildings.owner_id = $1;

-- name: GetBuildingByID :one
SELECT * 
FROM buildings
WHERE buildings.id = $1;

-- name: GetAllBuildings :many
SELECT *
FROM buildings
WHERE buildings.owner_id IS NULL
LIMIT 10;

-- name: GetListMyOwnedBuilding :many
SELECT * 
FROM buildings
WHERE owner_id = $1
LIMIT 10;

-- name: GetListBuildingByUserID :many
SELECT b.id as building_id, b.name as building_name, members.id as member_id
FROM members
INNER JOIN members_buildings
ON members.id = members_buildings.member_id
INNER JOIN buildings b
ON b.id = members_buildings.building_id
INNER JOIN users
ON users.id = members.user_id
WHERE users.id = $1 AND members_buildings.status_id = 2
LIMIT 10;

-- name: GetInvitationByUserID :many
SELECT b.id as building_id, b.name as building_name, members.id as member_id
FROM members
INNER JOIN members_buildings
ON members.id = members_buildings.member_id
INNER JOIN buildings b
ON b.id = members_buildings.building_id
INNER JOIN users
ON users.id = members.user_id
WHERE users.id = $1 AND members_buildings.status_id = 1
LIMIT 10;

-- name: CreateMember :one
INSERT INTO members(name, status, user_id)
VALUES ($1, $2, $3)
RETURNING id;

-- name: EditMemberName :execresult
UPDATE members
SET name = $2
WHERE id = $1;

-- name: GetMemberByName :one
SELECT *
FROM members
WHERE name = $1
LIMIT 1;

-- name: GetMemberByID :one
SELECT *
FROM members
WHERE id = $1
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

-- name: CreateRoom :one
INSERT INTO rooms(name, building_id)
VALUES ($1, $2)
RETURNING id;

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
SELECT EXISTS(
    SELECT mb.member_id, mb.building_id 
    FROM members_buildings mb
    WHERE mb.member_id = $1 AND mb.building_id = $2
    LIMIT 1
);

-- name: GetUserBuildingExist :one
SELECT EXISTS(
    SELECT * 
    FROM users
    INNER JOIN members
    ON users.id = members.user_id
    INNER JOIN members_buildings
    ON members_buildings.member_id = members.id
    WHERE users.id = $1 AND members_buildings.building_id = $2
    LIMIT 1
);

-- name: GetListMemberByBuildingID :many
SELECT members.id, members.name, members.status
FROM members
INNER JOIN members_buildings
ON members.id = members_buildings.member_id
WHERE members_buildings.building_id = $1
LIMIT 20;

-- name: CreateMemberBuilding :execresult
INSERT INTO members_buildings(member_id, building_id, status_id)
VALUES ($1, $2, 1);

-- name: EditMemberBuildingStatus :execresult
UPDATE members_buildings
SET status_id = $3
WHERE member_id = $1 AND building_id = $2;

-- name: DeleteMemberBuilding :exec
DELETE FROM members_buildings
WHERE member_id = $1 AND building_id = $2;
