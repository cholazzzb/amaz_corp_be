-- +migrate Up
INSERT INTO task_status(status)
VALUES ('TODO');

-- +migrate Up
INSERT INTO task_status(status)
VALUES ('IN PROGRESS');

-- +migrate Up
INSERT INTO task_status(status)
VALUES ('DONE');