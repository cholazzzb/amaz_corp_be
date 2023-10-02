-- +migrate Up
INSERT INTO products(name)
VALUES ('Free');

-- +migrate Up
INSERT INTO features(name, max_limit)
VALUES ('building_ownership', 1);