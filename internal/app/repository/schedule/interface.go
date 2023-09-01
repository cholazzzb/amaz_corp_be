package schedule

import (
	"context"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
)

type ScheduleRepo interface {
	ScheduleRepoCommand
	ScheduleRepoQuery
}

type ScheduleRepoCommand interface {
	CreateSchedule(
		ctx context.Context,
		roomID string,
	) (string, error)

	CreateTask(
		ctx context.Context,
		task ent.TaskWithDetailCommand,
	) error

	EditTask(
		ctx context.Context,
		taskID string,
		task ent.TaskWithDetailCommand,
	) error
}

type ScheduleRepoQuery interface {
	GetScheduleIDByRoomID(
		ctx context.Context,
		roomID string,
	) (ent.ScheduleQuery, error)

	GetTaskDetail(
		ctx context.Context,
		taskID string,
	) (ent.TaskDetailQuery, error)

	GetListTaskByScheduleID(
		ctx context.Context,
		scheduleID string,
		queryFilter ent.TaskQueryFilter,
	) ([]ent.TaskQuery, error)

	GetListTaskWithDetailByScheduleID(
		ctx context.Context,
		scheduleID string,
	) ([]ent.TaskWithDetailQuery, error)
}
