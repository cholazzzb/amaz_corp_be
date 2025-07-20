-- name: GetUserActiveTasks :many
SELECT t.id, t.title, t.description, el.level_name as energy_demand,
       tc.name as category, tc.color_hex, t.estimated_duration,
       t.priority, t.deadline, t.is_completed, t.created_at
FROM tasks t
JOIN energy_levels el ON t.energy_demand_id = el.id
LEFT JOIN task_categories tc ON t.category_id = tc.id
WHERE t.user_id = $1 AND t.is_completed = FALSE
ORDER BY t.priority DESC, t.deadline ASC NULLS LAST;

-- name: GetUserTasksByEnergyDemand :many
SELECT t.id, t.title, t.description, t.estimated_duration,
       t.priority, t.deadline, tc.name as category
FROM tasks t
LEFT JOIN task_categories tc ON t.category_id = tc.id
WHERE t.user_id = $1 AND t.energy_demand_id = $2 AND t.is_completed = FALSE
ORDER BY t.priority DESC, t.deadline ASC NULLS LAST;

-- name: GetUserOverdueTasks :many
SELECT t.id, t.title, t.description, el.level_name as energy_demand,
       t.deadline, t.priority,
       CURRENT_DATE - t.deadline as days_overdue
FROM tasks t
JOIN energy_levels el ON t.energy_demand_id = el.id
WHERE t.user_id = $1
  AND t.deadline < CURRENT_DATE
  AND t.is_completed = FALSE
ORDER BY days_overdue DESC;

-- name: GetUserTaskCompletionStats :one
SELECT
    COUNT(*) as total_tasks,
    COUNT(CASE WHEN is_completed = TRUE THEN 1 END) as completed_tasks,
    COUNT(CASE WHEN deadline < CURRENT_DATE AND is_completed = FALSE THEN 1 END) as overdue_tasks,
    ROUND(
        COUNT(CASE WHEN is_completed = TRUE THEN 1 END)::decimal /
        NULLIF(COUNT(*), 0) * 100, 2
    ) as completion_rate
FROM tasks
WHERE user_id = $1;
