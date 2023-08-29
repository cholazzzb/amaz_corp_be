package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/schedule"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type ScheduleService struct {
	repo   repo.ScheduleRepo
	logger *slog.Logger
}

func NewScheduleService(repo repo.ScheduleRepo) *ScheduleService {
	sublogger := logger.Get().With(slog.String("domain", "schedule"), slog.String("layer", "svc"))

	return &ScheduleService{
		repo:   repo,
		logger: sublogger,
	}
}

func (svc *ScheduleService) CreateSchedule(
	ctx context.Context,
	roomID string,
) (string, error) {
	res, err := svc.repo.CreateSchedule(ctx, roomID)
	if err != nil {
		svc.logger.Error(err.Error())
		return "", err
	}
	return res, nil
}

func (svc *ScheduleService) GetScheduleIDByRoomID(
	ctx context.Context,
	roomID string,
) (ent.ScheduleQuery, error) {
	res, err := svc.repo.GetScheduleIDByRoomID(ctx, roomID)
	if err != nil {
		svc.logger.Error(err.Error())
		return ent.ScheduleQuery{}, fmt.Errorf("failed to get scheduleID by roomID: %s", roomID)
	}
	return res, nil
}

func (svc *ScheduleService) GetTaskDetail(
	ctx context.Context,
	taskID string,
) (ent.TaskDetailQuery, error) {
	res, err := svc.repo.GetTaskDetail(ctx, taskID)
	if err != nil {
		svc.logger.Error(err.Error())
		return ent.TaskDetailQuery{}, fmt.Errorf("failed to get task detail with taskID: %s", taskID)
	}
	return res, nil
}

func (svc *ScheduleService) GetListTaskByScheduleID(
	ctx context.Context,
	scheduleID string,
) ([]ent.TaskQuery, error) {
	res, err := svc.repo.GetListTaskByScheduleID(ctx, scheduleID)
	if err != nil {
		svc.logger.Error(err.Error())
		return []ent.TaskQuery{}, fmt.Errorf("failed to get list task with scheduleID: %s", scheduleID)
	}
	return res, nil
}

func (svc *ScheduleService) AddTask(
	ctx context.Context,
	task ent.TaskWithDetailCommand,
) error {
	err := svc.repo.CreateTask(ctx, task)

	if err != nil {
		svc.logger.Error(err.Error())
		return errors.New("failed to add task")
	}

	return nil
}

func (svc *ScheduleService) EditTask(
	ctx context.Context,
	taskID string,
	task ent.TaskWithDetailCommand,
) error {
	err := svc.repo.EditTask(ctx, taskID, task)

	if err != nil {
		svc.logger.Error(err.Error())
		return errors.New("failed to edit task")
	}

	return nil
}

// func (svc *ScheduleService) AutoSchedule(
// 	ctx context.Context,
// 	scheduleID string,
// ) []ent.TaskWithDetail {
// 	tasks, err := svc.repo.GetListTaskAndDetailByScheduleID(ctx, scheduleID)
// 	if err != nil {
// 		return []ent.Task{}, nil
// 	}
// 	return tasks, nil
// }
