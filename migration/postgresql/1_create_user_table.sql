-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
  username text UNIQUE NOT NULL,
  password text NOT NULL,
  salt text NOT NULL
);
