package service

import (
	"context"
	"errors"
	"log/slog"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/schedule"
	domain "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type ScheduleService struct {
	repo   repo.ScheduleRepo
	cache  repo.ScheduleCacheRepo
	logger *slog.Logger
}

func NewScheduleService(
	repo repo.ScheduleRepo,
	cache repo.ScheduleCacheRepo,
) *ScheduleService {
	sublogger := logger.Get().With(slog.String("domain", "schedule"), slog.String("layer", "svc"))

	return &ScheduleService{
		repo:   repo,
		cache:  cache,
		logger: sublogger,
	}
}

func (svc *ScheduleService) CreateSchedule(
	ctx context.Context,
	sch domain.ScheduleCommand,
) (string, error) {
	res, err := svc.repo.CreateSchedule(ctx, sch.Name, sch.RoomID)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (svc *ScheduleService) GetListScheduleByRoomID(
	ctx context.Context,
	roomID string,
) ([]domain.ScheduleQuery, error) {
	res, err := svc.repo.GetListScheduleByRoomID(ctx, roomID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (svc *ScheduleService) GetTaskDetail(
	ctx context.Context,
	taskID string,
) (domain.TaskDetailQuery, error) {
	res, err := svc.repo.GetTaskDetail(ctx, taskID)
	if err != nil {
		return domain.TaskDetailQuery{}, err
	}
	return res, nil
}

func (svc *ScheduleService) GetListTaskByScheduleID(
	ctx context.Context,
	scheduleID string,
	queryFilter domain.TaskQueryFilter,
) ([]domain.TaskQuery, error) {
	res, err := svc.repo.GetListTaskByScheduleID(ctx, scheduleID, queryFilter)
	if err != nil {
		return []domain.TaskQuery{}, err
	}
	return res, nil
}

func (svc *ScheduleService) AddTask(
	ctx context.Context,
	task domain.TaskWithDetailCommand,
) error {
	err := svc.repo.CreateTask(ctx, task)

	if err != nil {
		return err
	}

	return nil
}

func (svc *ScheduleService) EditTask(
	ctx context.Context,
	taskID string,
	task domain.TaskWithDetailCommand,
) error {
	err := svc.repo.EditTask(ctx, taskID, task)

	if err != nil {
		return err
	}

	return nil
}

func (svc *ScheduleService) AutoSchedulePreview(
	ctx context.Context,
	scheduleID string,
) ([]domain.TaskWithDetailQuery, error) {
	twds, err := svc.repo.GetListTaskWithDetailByScheduleID(ctx, scheduleID)
	graph := domain.CreateGraph(twds)
	sorted := domain.TopologicalSort(graph)

	if err != nil {
		return []domain.TaskWithDetailQuery{}, err
	}
	return sorted, nil
}

func (svc *ScheduleService) AutoScheduleSave(
	ctx context.Context,
	scheduleID string,
) error {
	_, err := svc.cache.GetAutoSchedule(ctx, scheduleID)
	if err != nil {
		return err
	}

	// TODO: Bulk Update
	return errors.New("Not Implemented")
}
