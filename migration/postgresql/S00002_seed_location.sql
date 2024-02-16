-- +migrate Up
INSERT INTO members_buildings_status(id, name)
VALUES (1, 'invited');

INSERT INTO members_buildings_status(id, name)
VALUES (2, 'joined');
