package schedule

import (
	"context"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
)

type ScheduleCacheRepo interface {
	AutoScheduleRepoCommand
	AutoScheduleRepoQuery
}

type AutoScheduleRepoCommand interface {
	SaveAutoSchedule(
		ctx context.Context,
		scheduleID string,
		tasks []ent.TaskWithDetailQuery,
	) error
	InvalidateAutoSchedule(
		ctx context.Context,
		scheduleID string,
	) error
}

type AutoScheduleRepoQuery interface {
	GetAutoSchedule(
		ctx context.Context,
		scheduleID string,
	) ([]ent.TaskWithDetailQuery, error)
}
