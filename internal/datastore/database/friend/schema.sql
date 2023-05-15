CREATE TABLE IF NOT EXISTS members (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  name varchar(255) NOT NULL,
  status text NOT NULL,
  CONSTRAINT fk_members_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS friends (
  member1_id BIGINT NOT NULL,
  member2_id BIGINT NOT NULl,
  CONSTRAINT fk_member1_id FOREIGN KEY(member1_id) REFERENCES members(id),
  CONSTRAINT fk_member2_id FOREIGN KEY(member2_id) REFERENCES members(id)
);