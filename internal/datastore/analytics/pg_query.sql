-- name: GetUserDashboardSummary :one
SELECT
    u.first_name, u.last_name, u.email,
    (SELECT COUNT(*) FROM tasks WHERE user_id = u.id AND is_completed = FALSE) as pending_tasks,
    (SELECT COUNT(*) FROM tasks WHERE user_id = u.id AND is_completed = TRUE) as completed_tasks,
    (SELECT COUNT(*) FROM schedules WHERE user_id = u.id AND schedule_date >= CURRENT_DATE - INTERVAL '7 days') as recent_schedules,
    (SELECT AVG(user_rating) FROM schedules WHERE user_id = u.id AND user_rating IS NOT NULL) as avg_schedule_rating,
    (SELECT COUNT(*) FROM user_energy_patterns WHERE user_id = u.id AND date >= CURRENT_DATE - INTERVAL '7 days') as recent_energy_entries
FROM users u
WHERE u.id = $1 AND u.is_active = TRUE;

-- name: GetUserEnergyTrends :one
SELECT
    DATE_TRUNC('week', uep.date) as week_start,
    ts.slot_name,
    AVG(el.level_value) as avg_energy,
    COUNT(*) as entries
FROM user_energy_patterns uep
JOIN time_slots ts ON uep.time_slot_id = ts.id
JOIN energy_levels el ON uep.energy_level_id = el.id
WHERE uep.user_id = $1
  AND uep.date >= CURRENT_DATE - INTERVAL '30 days'
GROUP BY DATE_TRUNC('week', uep.date), ts.id, ts.slot_name, ts.sort_order
ORDER BY week_start, ts.sort_order;
