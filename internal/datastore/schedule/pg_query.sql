-- name: GetUserScheduleForSpecificDate :many
SELECT s.id as schedule_id, s.schedule_date, s.generated_at, s.optimization_score,
       si.id as item_id, si.start_time, si.end_time, si.title, si.description,
       si.item_type, si.is_completed, t.id as task_id
FROM schedules s
LEFT JOIN schedule_items si ON s.id = si.schedule_id
LEFT JOIN tasks t ON si.task_id = t.id
WHERE s.user_id = $1 AND s.schedule_date = $2 AND s.is_active = TRUE
ORDER BY si.sort_order;

-- name: GetRecentSchedulesWithCompletionRates :many
SELECT s.id, s.schedule_date, s.generated_at, s.optimization_score, s.user_rating,
       COUNT(si.id) as total_items,
       COUNT(CASE WHEN si.is_completed = TRUE THEN 1 END) as completed_items,
       ROUND(
           COUNT(CASE WHEN si.is_completed = TRUE THEN 1 END)::decimal /
           NULLIF(COUNT(si.id), 0) * 100, 2
       ) as completion_rate
FROM schedules s
LEFT JOIN schedule_items si ON s.id = si.schedule_id
WHERE s.user_id = $1 AND s.is_active = TRUE
GROUP BY s.id, s.schedule_date, s.generated_at, s.optimization_score, s.user_rating
ORDER BY s.schedule_date DESC
LIMIT 10;

-- name: GetScheduleItemsForAIOptimizationInput :many
SELECT t.id, t.title, t.description, el.level_value as energy_demand,
       t.estimated_duration, t.priority, t.deadline,
       tc.name as category
FROM tasks t
JOIN energy_levels el ON t.energy_demand_id = el.id
LEFT JOIN task_categories tc ON t.category_id = tc.id
WHERE t.user_id = $1 AND t.is_completed = FALSE
ORDER BY t.priority DESC, t.deadline ASC NULLS LAST;
