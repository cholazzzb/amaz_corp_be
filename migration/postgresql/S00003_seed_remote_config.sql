-- +migrate Up
INSERT INTO remote_config(key, value)
VALUES ('apk-version', '0.1.0');
