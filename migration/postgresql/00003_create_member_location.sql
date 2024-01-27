-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate Up
CREATE TABLE IF NOT EXISTS members (
  id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
  user_id uuid NOT NULL,
  name varchar(255) NOT NULL,
  status text NOT NULL,
  room_id uuid,
  CONSTRAINT fk_members_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT fk_members_room_id FOREIGN KEY(room_id) REFERENCES rooms(id) ON DELETE CASCADE
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS friends (
  member1_id uuid NOT NULL,
  member2_id uuid NOT NULl,
  CONSTRAINT fk_member1_id FOREIGN KEY(member1_id) REFERENCES members(id),
  CONSTRAINT fk_member2_id FOREIGN KEY(member2_id) REFERENCES members(id)
);

CREATE TYPE mb_status AS ENUM ('invited', 'joined');
-- +migrate Up
CREATE TABLE IF NOT EXISTS members_buildings (
    member_id uuid NOT NULL,
    building_id uuid NOT NULL,
    status mb_status NOT NULL,
    CONSTRAINT fk_mb_member_id FOREIGN KEY(member_id) REFERENCES members(id),
    CONSTRAINT fk_mb_building_id FOREIGN KEY(building_id) REFERENCES buildings(id)
);

-- Notes for study
-- -- +migrate Up
-- ALTER TABLE members
-- ADD COLUMN room_id varchar(36);

-- -- +migrate Up
-- ALTER TABLE members
-- ADD CONSTRAINT fk_members_room_id FOREIGN KEY(room_id) REFERENCES rooms(id);
