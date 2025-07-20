-- name: RegisterNewUser :one
INSERT INTO users (email, password_hash, first_name, last_name, timezone)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, created_at;

-- name: UpdateUserProfile :one
UPDATE users
SET first_name = $2,
    last_name = $3,
    timezone = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND is_active = TRUE
RETURNING id, first_name, last_name, timezone, updated_at;

-- name: DeactivateUserAccount :one
UPDATE users
SET is_active = FALSE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, is_active;
