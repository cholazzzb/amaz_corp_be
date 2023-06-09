-- +migrate Up
CREATE TABLE IF NOT EXISTS members_buildings (
    member_id BIGINT NOT NULL,
    building_id BIGINT NOT NULL,
    CONSTRAINT fk_mb_member_id FOREIGN KEY(member_id) REFERENCES members(id),
    CONSTRAINT fk_mb_building_id FOREIGN KEY(building_id) REFERENCES buildings(id)
);

-- +migrate Up
ALTER TABLE members
ADD COLUMN room_id BIGINT AFTER status;

-- +migrate Up
ALTER TABLE members
ADD CONSTRAINT fk_members_room_id FOREIGN KEY(room_id) REFERENCES rooms(id);
