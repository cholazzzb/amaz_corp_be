package schedule

import (
	"context"
	"time"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
)

type TaskRepo interface {
	TaskRepoQuery
}

type TaskRepoQuery interface {
	GetListTaskStatus(
		ctx context.Context,
	) ([]schedule.TaskStatusQuery, error)

	GetTasksByRoomID(
		ctx context.Context,
		roomID string,
		page int32,
		pageSize int32,
		startTime *time.Time,
		endTime *time.Time,
	) ([]schedule.TaskByRoomIDQuery, error)
}
