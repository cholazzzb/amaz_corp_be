-- +migrate Up
CREATE TABLE IF NOT EXISTS remote_config (
    id serial UNIQUE NOT NULL PRIMARY KEY,
    key text UNIQUE NOT NULL,
    value text UNIQUE NOT NULL
);