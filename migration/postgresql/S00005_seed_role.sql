-- +migrate Up
INSERT INTO roles(name)
VALUES ('admin');

-- +migrate Up
INSERT INTO roles(name)
VALUES ('member');
