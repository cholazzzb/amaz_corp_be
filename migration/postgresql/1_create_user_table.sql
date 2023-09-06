-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id varchar(36) UNIQUE NOT NULL PRIMARY KEY UNIQUE,
  username text UNIQUE NOT NULL,
  password text NOT NULL,
  salt text NOT NULL
);
