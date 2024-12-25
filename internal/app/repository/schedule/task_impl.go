package schedule

import (
	"context"
	"database/sql"
	"time"

	schedulepostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/schedule/postgresql"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/google/uuid"
)

func (r *PostgresScheduleRepository) GetTasksByRoomID(
	ctx context.Context,
	roomID string,
	page int32,
	pageSize int32,
	startTime *time.Time,
	endTime *time.Time,
) ([]schedule.TaskByRoomIDQuery, error) {
	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		r.logger.Error(err.Error())
		return []schedule.TaskByRoomIDQuery{}, err
	}

	stTime := sql.NullTime{}
	if startTime != nil {
		stTime.Scan(*startTime)
	}
	edTime := sql.NullTime{}
	if endTime != nil {
		edTime.Scan(*endTime)
	}

	queryResult, err := r.Postgres.GetTasksByRoomID(ctx,
		schedulepostgres.GetTasksByRoomIDParams{
			ID:        roomUUID,
			Limit:     pageSize,
			Offset:    (page - 1) * pageSize,
			StartTime: stTime,
			EndTime:   edTime,
		},
	)
	if err != nil {
		r.logger.Error(err.Error())
		return []schedule.TaskByRoomIDQuery{}, err
	}

	tasks := []schedule.TaskByRoomIDQuery{}

	for _, task := range queryResult {
		tasks = append(tasks, schedule.TaskByRoomIDQuery{
			ID:           task.ID.String(),
			TaskName:     task.TaskName.String,
			StartTime:    task.StartTime.Time.String(),
			EndTime:      task.EndTime.Time.String(),
			StatusID:     task.StatusID.UUID.String(),
			AssigneeID:   task.AssigneeID.UUID.String(),
			OwnerID:      task.OwnerID.UUID.String(),
			ScheduleName: task.ScheduleName,
			RoomName:     task.RoomName,
		})
	}

	return tasks, nil
}
