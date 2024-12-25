package service

import (
	"context"
	"time"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
)

func (svc *ScheduleService) GetListTaskStatus(ctx context.Context) ([]ent.TaskStatusQuery, error) {
	out, err := svc.repo.GetListTaskStatus(ctx)
	if err != nil {
		return []ent.TaskStatusQuery{}, err
	}
	return out, nil
}

func (svc *ScheduleService) CreateTaskDependency(
	ctx context.Context,
	taskDep ent.TaskDependencyCommand,
) error {
	err := svc.repo.CreateTaskDependency(ctx, taskDep)
	if err != nil {
		return err
	}
	return nil
}

func (svc *ScheduleService) EditTaskDependency(
	ctx context.Context,
	taskDep ent.TaskDependencyCommand,
) error {
	err := svc.repo.EditTaskDependency(ctx, taskDep)
	if err != nil {
		return err
	}
	return nil
}

func (svc *ScheduleService) DeleteTaskDependency(
	ctx context.Context,
	taskDep ent.TaskDependencyCommand,
) error {
	err := svc.repo.DeleteTaskDependeny(ctx, taskDep)
	if err != nil {
		return nil
	}
	return nil
}

func (svc *ScheduleService) GetTasksByRoomID(
	ctx context.Context,
	roomID string,
	page int32,
	pageSize int32,
	startTime *time.Time,
	endtime *time.Time,
) ([]ent.TaskByRoomIDQuery, error) {
	tasks, err := svc.repo.GetTasksByRoomID(
		ctx,
		roomID,
		page,
		pageSize,
		startTime,
		endtime,
	)
	if err != nil {
		return []ent.TaskByRoomIDQuery{}, err
	}

	return tasks, nil
}
