-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate Up
CREATE TABLE IF NOT EXISTS buildings (
    id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name text NOT NULL,
    owner_id uuid,
    CONSTRAINT fk_owner_id FOREIGN KEY(owner_id) REFERENCES users(id)
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS rooms (
    id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name text NOT NULL,
    building_id uuid NOT NULL,
    CONSTRAINT fk_building_id FOREIGN KEY(building_id) REFERENCES buildings(id)
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS sessions (
    id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
    room_id uuid NOT NULL,
    start_time timestamp,
    end_time timestamp,
    CONSTRAINT fk_room_id FOREIGN KEY(room_id) REFERENCES rooms(id)
);