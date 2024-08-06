-- +migrate Up
INSERT INTO roles(name)
VALUES ('admin');

INSERT INTO roles(name)
VALUES ('member');
