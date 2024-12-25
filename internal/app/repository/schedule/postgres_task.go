package schedule

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	schedulepostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/schedule/postgresql"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
)

func (r *PostgresScheduleRepository) GetListTaskStatus(ctx context.Context) ([]ent.TaskStatusQuery, error) {
	listStatus, err := r.Postgres.GetListTaskStatus(ctx)

	out := []ent.TaskStatusQuery{}
	if err != nil {
		r.logger.Error(err.Error())
		return out, err
	}

	for _, ls := range listStatus {
		out = append(out, ent.TaskStatusQuery{
			ID:   ls.ID.String(),
			Name: ls.Status,
		})
	}

	return out, nil
}

func (r *PostgresScheduleRepository) CreateTaskDependency(
	ctx context.Context,
	taskDep ent.TaskDependencyCommand,
) error {
	taskID, err := uuid.Parse(taskDep.TaskID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	taskDepID, err := uuid.Parse(taskDep.DependencyID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	_, err = r.Postgres.CreateTaskDependency(
		ctx,
		schedulepostgres.CreateTaskDependencyParams{
			TaskID:         taskID,
			DependedTaskID: taskDepID,
		},
	)

	if err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("failed to create task dependency")
	}

	return nil
}

func (r *PostgresScheduleRepository) EditTaskDependency(
	ctx context.Context,
	taskDep ent.TaskDependencyCommand,
) error {
	taskID, err := uuid.Parse(taskDep.TaskID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	taskDepID, err := uuid.Parse(taskDep.DependencyID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	_, err = r.Postgres.EditTaskDependency(
		ctx,
		schedulepostgres.EditTaskDependencyParams{
			TaskID:         taskID,
			DependedTaskID: taskDepID,
		},
	)
	if err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("failed to edit task dependency with id: %s", taskDep.TaskID)
	}

	return nil
}

func (r *PostgresScheduleRepository) DeleteTaskDependeny(
	ctx context.Context,
	taskDep ent.TaskDependencyCommand,
) error {
	taskID, err := uuid.Parse(taskDep.TaskID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	taskDepID, err := uuid.Parse(taskDep.DependencyID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	err = r.Postgres.DeleteTaskDependency(
		ctx,
		schedulepostgres.DeleteTaskDependencyParams{
			TaskID:         taskID,
			DependedTaskID: taskDepID,
		},
	)
	if err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("failed to delete task dependency with id: %s", taskDep.TaskID)
	}

	return errors.New("not implemented")
}
