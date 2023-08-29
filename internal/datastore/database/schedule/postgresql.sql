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
LIMIT 100;

-- name: GetListTaskAndDetailByScheduleID :many
SELECT tasks.id,
	   tasks.schedule_id,
	   tasks.start_time,
	   tasks.duration_day,
	   tasks.task_detail_id,
	   task_details.name,
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
GROUP BY tasks.id
LIMIT 100;

-- name: CreateTask :execresult
INSERT INTO tasks(schedule_id, start_time, duration_day, task_detail_id)
VALUES ($1, $2, $3, $4);

-- name: CreateTaskDetail :one
INSERT INTO task_details(name, owner_id, assignee_id, status)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: EditTask :execresult
UPDATE tasks
SET start_time = $2,
    duration_day = $3,
    task_detail_id = $4
WHERE id = $1;