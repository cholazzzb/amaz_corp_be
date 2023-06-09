-- +migrate Up
CREATE TABLE IF NOT EXISTS buildings (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name text NOT NULL
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS rooms (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name text NOT NULL,
    building_id BIGINT NOT NULL,
    CONSTRAINT fk_building_id FOREIGN KEY(building_id) REFERENCES buildings(id)
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS sessions (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    room_id BIGINT NOT NULL,
    start_time datetime,
    end_time datetime,
    CONSTRAINT fk_room_id FOREIGN KEY(room_id) REFERENCES rooms(id)
);