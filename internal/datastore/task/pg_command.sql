-- name: CreateTask :one
INSERT INTO tasks (user_id, title, description, energy_demand_id, category_id,
                   estimated_duration, priority, deadline)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, user_id, title, created_at;

-- name: UpdateTask :one
UPDATE tasks
SET title = $2,
    description = $3,
    energy_demand_id = $4,
    category_id = $5,
    estimated_duration = $6,
    priority = $7,
    deadline = $8,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND user_id = $9
RETURNING id, title, updated_at;

-- name: MarkTaskAsCompleted :one
UPDATE tasks
SET is_completed = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND user_id = $2
RETURNING id, title, is_completed, updated_at;

-- name: DeleteTask :one
DELETE FROM tasks
WHERE id = $1 AND user_id = $2
RETURNING id, title;
