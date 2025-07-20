-- name: GetUserEnergyLevels :one
SELECT id, level_name, level_value, description
FROM energy_levels
ORDER BY level_value DESC;

-- name: GetUserTimeSlots :one
SELECT id, slot_name, start_time, end_time, sort_order
FROM time_slots
ORDER BY sort_order;

-- name: GetUserTaskCategories :one
SELECT id, name, description, color_hex
FROM task_categories
ORDER BY name;
