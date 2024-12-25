package service

import (
	"context"
	"log/slog"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/report"

	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type ReportService struct {
	ss     *ScheduleService
	logger *slog.Logger
}

func NewReportService(
	ss *ScheduleService,
) *ReportService {
	sublogger := logger.Get().With(slog.String("domain", "report"), slog.String("layer", "svc"))

	return &ReportService{
		ss:     ss,
		logger: sublogger,
	}
}

func (svc *ReportService) GetReportBySchedule(
	ctx context.Context,
	scheduleID string,
) (ent.ReportByScheduleQuery, error) {
	tasks, err := svc.ss.GetListTaskWithDetailByScheduleID(ctx, scheduleID)
	if err != nil {
		return ent.ReportByScheduleQuery{}, err
	}

	totalTask := int32(len(tasks))
	var totalTaskDur int32 = 0
	for _, task := range tasks {
		totalTaskDur += task.DurationDay
	}
	out := ent.ReportByScheduleQuery{
		ScheduleID:      scheduleID,
		TotalTask:       totalTask,
		AvgTaskDuration: totalTaskDur / totalTask,
	}

	return out, nil
}
