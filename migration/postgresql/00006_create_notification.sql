-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate Up
CREATE TABLE IF NOT EXISTS notifications (
    id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id uuid,
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);