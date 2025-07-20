-- name: UpsertUserPreferences :one
INSERT INTO user_preferences (user_id, prefer_morning_focus, prefer_afternoon_focus,
                             prefer_evening_focus, break_frequency, break_duration,
                             work_start_time, work_end_time, buffer_time,
                             max_high_energy_blocks, notifications_enabled)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT (user_id)
DO UPDATE SET
    prefer_morning_focus = EXCLUDED.prefer_morning_focus,
    prefer_afternoon_focus = EXCLUDED.prefer_afternoon_focus,
    prefer_evening_focus = EXCLUDED.prefer_evening_focus,
    break_frequency = EXCLUDED.break_frequency,
    break_duration = EXCLUDED.break_duration,
    work_start_time = EXCLUDED.work_start_time,
    work_end_time = EXCLUDED.work_end_time,
    buffer_time = EXCLUDED.buffer_time,
    max_high_energy_blocks = EXCLUDED.max_high_energy_blocks,
    notifications_enabled = EXCLUDED.notifications_enabled,
    updated_at = CURRENT_TIMESTAMP
RETURNING id, user_id, updated_at;
