-- name: GetUserProfile :one
SELECT id, email, first_name, last_name, timezone, created_at, updated_at, is_active
FROM users
WHERE id = $1 AND is_active = TRUE;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, first_name, last_name, timezone, is_active
FROM users
WHERE email = $1 AND is_active = TRUE;

-- admin
-- name: GetAllActiveUsers :many
SELECT id, email, first_name, last_name, timezone, created_at,
       (SELECT COUNT(*) FROM tasks WHERE user_id = users.id) as total_tasks,
       (SELECT COUNT(*) FROM schedules WHERE user_id = users.id) as total_schedules
FROM users
WHERE is_active = TRUE
ORDER BY created_at DESC;

-- name: GetUserScheduleEffectivenessMetrics :one
SELECT
    s.schedule_date,
    s.optimization_score,
    s.user_rating,
    COUNT(si.id) as total_items,
    COUNT(CASE WHEN si.is_completed = TRUE THEN 1 END) as completed_items,
    AVG(CASE WHEN si.actual_end_time IS NOT NULL AND si.actual_start_time IS NOT NULL
        THEN EXTRACT(EPOCH FROM (si.actual_end_time - si.actual_start_time))/60
        ELSE NULL END) as avg_actual_duration
FROM schedules s
LEFT JOIN schedule_items si ON s.id = si.schedule_id
WHERE s.user_id = $1 AND s.schedule_date >= CURRENT_DATE - INTERVAL '30 days'
GROUP BY s.id, s.schedule_date, s.optimization_score, s.user_rating
ORDER BY s.schedule_date DESC;
