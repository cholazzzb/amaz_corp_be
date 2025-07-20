-- name: CreateNewSchedule :one
INSERT INTO schedules (user_id, schedule_date, ai_model_version, optimization_score, raw_ai_response)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, schedule_date, generated_at;

-- name: AddScheduleItem :one
INSERT INTO schedule_items (schedule_id, task_id, start_time, end_time, title, description,
                           item_type, sort_order)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, schedule_id, start_time, end_time, title;

-- name: UpdateScheduleItemCompletion :one
UPDATE schedule_items
SET is_completed = $2,
    actual_start_time = $3,
    actual_end_time = $4
WHERE id = $1
RETURNING id, is_completed, actual_start_time, actual_end_time;

-- name: SubmitScheduleFeedback :one
UPDATE schedules
SET user_rating = $2,
    user_feedback = $3
WHERE id = $1 AND user_id = $4
RETURNING id, user_rating, user_feedback;

-- name: InsertScheduleFeedback :one
INSERT INTO schedule_feedback (schedule_id, user_id, feedback_date, energy_accuracy_rating,
                              task_completion_rate, overall_satisfaction, suggestions, would_use_again)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, schedule_id, feedback_date;
