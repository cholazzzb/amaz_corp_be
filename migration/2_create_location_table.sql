-- +migrate Up
CREATE TABLE IF NOT EXISTS buildings (
    id varchar(36) UNIQUE NOT NULL PRIMARY KEY,
    name text NOT NULL
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS rooms (
    id varchar(36) UNIQUE NOT NULL PRIMARY KEY,
    name text NOT NULL,
    building_id varchar(36) NOT NULL,
    CONSTRAINT fk_building_id FOREIGN KEY(building_id) REFERENCES buildings(id)
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS sessions (
    id varchar(36) UNIQUE NOT NULL PRIMARY KEY,
    room_id varchar(36) NOT NULL,
    start_time datetime,
    end_time datetime,
    CONSTRAINT fk_room_id FOREIGN KEY(room_id) REFERENCES rooms(id)
);