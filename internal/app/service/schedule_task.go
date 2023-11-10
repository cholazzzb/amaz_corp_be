package service

import (
	"context"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
)

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
