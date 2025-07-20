-- name: GetUserPreferences :one
SELECT user_id, prefer_morning_focus, prefer_afternoon_focus, prefer_evening_focus,
       break_frequency, break_duration, work_start_time, work_end_time,
       buffer_time, max_high_energy_blocks, notifications_enabled
FROM user_preferences
WHERE user_id = $1;
