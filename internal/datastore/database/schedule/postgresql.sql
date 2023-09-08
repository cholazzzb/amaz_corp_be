-- name: CreateScheduleByRoomID :one
INSERT INTO schedules(room_id)
VALUES($1)
RETURNING room_id;

-- name: GetScheduleIdByRoomID :one
SELECT * 
FROM schedules
WHERE schedules.room_id = $1
LIMIT 1;

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
	   task_details.owner_id,
	   task_details.assignee_id,
	   task_details.status,
	   ARRAY_AGG (TD.depended_task_id)
FROM tasks
INNER JOIN task_details 
ON tasks.task_detail_id = task_details.id
FULL OUTER JOIN tasks_dependencies TD
ON TD.task_id = tasks.id
WHERE tasks.schedule_id = $1
	AND tasks.start_time = $2
	AND tasks.end_time = $3
	AND task_details.owner_id = $4
	AND task_details.assignee_id = $5
GROUP BY tasks.id
LIMIT 100;

-- name: CreateTask :execresult
INSERT INTO tasks(name, start_time, end_time, schedule_id, task_detail_id)
VALUES ($1, $2, $3, $4, $5);

-- name: CreateTaskDetail :one
INSERT INTO task_details(owner_id, assignee_id, status)
VALUES ($1, $2, $3)
RETURNING id;

-- name: EditTask :execresult
UPDATE tasks
SET start_time = $2,
    end_time = $3,
    task_detail_id = $4
WHERE id = $1;