CREATE TABLE IF NOT EXISTS members (
  user_id BIGINT UNSIGNED,
  name text NOT NULL,
  CONSTRAINT pk_members PRIMARY KEY (user_id, name),
  CONSTRAINT fk_members_user_id FOREIGN KEY (user_id) REFERENCES user (user_id)
);