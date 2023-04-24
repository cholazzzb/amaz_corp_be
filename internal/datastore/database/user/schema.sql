CREATE TABLE IF NOT EXISTS users (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username text NOT NULL,
  password text NOT NULL,
  salt text NOT NULL
);