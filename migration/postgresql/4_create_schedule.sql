-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate Up
CREATE TABLE IF NOT EXISTS schedules (
    id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name text NOT NULL,
    room_id uuid NOT NULL,
    CONSTRAINT fk_room_id FOREIGN KEY(room_id) REFERENCES rooms(id) ON DELETE CASCADE
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS task_details (
    id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
    owner_id uuid,
    assignee_id uuid,
    status text,
    CONSTRAINT fk_owner_id FOREIGN KEY(owner_id) REFERENCES members(id) ON DELETE CASCADE,
    CONSTRAINT fk_assignee_id FOREIGN KEY(assignee_id) REFERENCES members(id) ON DELETE CASCADE
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS tasks (
    id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name text,
    start_time timestamp,
    end_time timestamp,
    schedule_id uuid NOT NULL,
    task_detail_id uuid UNIQUE NOT NULL,
    CONSTRAINT fk_schedule_id FOREIGN KEY(schedule_id) REFERENCES schedules(id) ON DELETE CASCADE,
    CONSTRAINT fk_task_detail_id FOREIGN KEY(task_detail_id) REFERENCES task_details(id) ON DELETE CASCADE
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS tasks_dependencies (
    task_id uuid NOT NULL,
    depended_task_id uuid NOT NULL,
    CONSTRAINT fk_task_id FOREIGN KEY(task_id) REFERENCES tasks(id),
    CONSTRAINT fk_depended_task_id FOREIGN KEY(depended_task_id) REFERENCES tasks(id)
);