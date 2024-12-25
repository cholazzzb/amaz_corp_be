-- name: CreateScheduleByRoomID :one
INSERT INTO schedules(name, room_id)
VALUES($1, $2)
RETURNING id;

-- name: GetListScheduleByRoomID :many
SELECT * 
FROM schedules
WHERE schedules.room_id = $1
LIMIT 10;

-- name: GetListTaskStatus :many
SELECT *
FROM task_status
LIMIT 10;

-- name: GetTaskDetailByID :one
SELECT *
FROM task_details
WHERE task_details.id = $1
LIMIT 1;

-- name: GetListTaskByScheduleID :many
SELECT *
FROM tasks
WHERE tasks.schedule_id = $1
	AND (tasks.start_time >= sqlc.narg('start_time') OR sqlc.narg('start_time') IS NULL) 
	AND (tasks.end_time <= sqlc.narg('end_time') OR tasks.end_time <= sqlc.narg('end_time') + interval '1 day' OR sqlc.narg('end_time') IS NULL)
LIMIT 100;

-- name: GetListTaskAndDetailByScheduleID :many
SELECT tasks.id,
	   tasks.name,
	   tasks.start_time,
	   tasks.end_time,
	   tasks.schedule_id,
	   tasks.task_detail_id,
	   ARRAY_AGG (DISTINCT task_details.owner_id) AS owner_id,
	   ARRAY_AGG (DISTINCT task_details.assignee_id) AS assignee_id,
       ARRAY_AGG (DISTINCT task_details.status_id) AS status_id,
	   ARRAY_AGG (TD.depended_task_id) AS depended_task_id
FROM tasks
INNER JOIN task_details 
ON tasks.task_detail_id = task_details.id
FULL OUTER JOIN tasks_dependencies TD
ON TD.task_id = tasks.id
WHERE tasks.schedule_id = $1
	-- TODO: Add Filter
	-- AND tasks.start_time = $2
	-- AND tasks.end_time = $3
	-- AND task_details.owner_id = $4
	-- AND task_details.assignee_id = $5
GROUP BY tasks.id
LIMIT 100;

-- name: CreateTask :execresult
INSERT INTO tasks(name, start_time, end_time, schedule_id, task_detail_id)
VALUES ($1, $2, $3, $4, $5);

-- name: CreateTaskDetail :one
INSERT INTO task_details(owner_id, assignee_id, status_id)
VALUES ($1, $2, $3)
RETURNING id;

-- name: EditTask :execresult
UPDATE tasks
SET start_time = $2,
    end_time = $3,
    task_detail_id = $4
WHERE id = $1;

-- name: CreateTaskDependency :one
INSERT INTO tasks_dependencies(task_id, depended_task_id)
VALUES ($1, $2)
RETURNING task_id;

-- name: EditTaskDependency :execresult
UPDATE tasks_dependencies
SET depended_task_id = $2
WHERE task_id = $1;

-- name: DeleteTaskDependency :exec
DELETE FROM tasks_dependencies
WHERE task_id = $1 AND depended_task_id = $2;

-- name: GetTasksByRoomID :many
SELECT tasks.id, tasks.name AS task_name, tasks.start_time, tasks.end_time, task_details.status_id, task_details.assignee_id, task_details.owner_id, schedules.name AS schedule_name, rooms.name AS room_name
FROM tasks 
INNER JOIN schedules
ON schedules.id = tasks.schedule_id
INNER JOIN rooms
ON schedules.room_id = rooms.id
INNER JOIN task_details
ON tasks.task_detail_id = task_details.id
WHERE rooms.id = $1 
	AND (tasks.start_time >= sqlc.narg('start_time') OR sqlc.narg('start_time') IS NULL) 
	AND (tasks.end_time <= sqlc.narg('end_time') OR tasks.end_time <= sqlc.narg('end_time') + interval '1 day' OR sqlc.narg('end_time') IS NULL)
LIMIT $2
OFFSET $3;