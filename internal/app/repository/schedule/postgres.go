package schedule

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"time"

	"github.com/google/uuid"

	schedulepostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/schedule/postgresql"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type PostgresScheduleRepository struct {
	db       *sql.DB
	Postgres *schedulepostgres.Queries
	logger   *slog.Logger
}

func NewPostgresLocationRepository(postgresRepo *database.SqlRepository) *PostgresScheduleRepository {
	sublogger := logger.Get().With(slog.String("domain", "schedule"), slog.String("layer", "repo"))
	queries := schedulepostgres.New(postgresRepo.Db)

	return &PostgresScheduleRepository{
		db:       postgresRepo.Db,
		Postgres: queries,
		logger:   sublogger,
	}
}

func (r *PostgresScheduleRepository) CreateSchedule(
	ctx context.Context,
	roomID string,
) (string, error) {
	res, err := r.Postgres.CreateScheduleByRoomID(ctx, roomID)

	if err != nil {
		r.logger.Error(err.Error())
		return "", fmt.Errorf("failed to create schedule with roomID: %s", roomID)
	}

	return res, nil
}

func (r *PostgresScheduleRepository) GetScheduleIDByRoomID(
	ctx context.Context,
	roomID string,
) (ent.ScheduleQuery, error) {
	res, err := r.Postgres.GetScheduleIdByRoomID(ctx, roomID)
	if err != nil {
		r.logger.Error(err.Error())
		return ent.ScheduleQuery{}, err
	}

	return ent.ScheduleQuery{
		ID:     res.ID.String(),
		RoomID: res.RoomID,
	}, nil
}

func (r *PostgresScheduleRepository) GetTaskDetail(
	ctx context.Context,
	taskID string,
) (ent.TaskDetailQuery, error) {
	tUUID, err := uuid.Parse(taskID)
	if err != nil {
		r.logger.Error(err.Error())
		return ent.TaskDetailQuery{}, err
	}

	res, err := r.Postgres.GetTaskDetailByID(ctx, tUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return ent.TaskDetailQuery{}, err
	}

	return ent.TaskDetailQuery{
		ID:         res.ID.String(),
		Name:       res.Name.String,
		OwnerID:    res.OwnerID.String,
		AssigneeID: res.AssigneeID.String,
		Status:     res.Status.String,
	}, nil
}

func (r *PostgresScheduleRepository) GetListTaskByScheduleID(
	ctx context.Context,
	scheduleID string,
	queryFilter ent.TaskQueryFilter,
) ([]ent.TaskQuery, error) {
	sUUID, err := uuid.Parse(scheduleID)
	tasks := []ent.TaskQuery{}
	if err != nil {
		r.logger.Error(err.Error())
		return tasks, err
	}
	startTime := sql.NullTime{}
	startTime.Scan(queryFilter.StartTime)
	endTime := sql.NullTime{}
	endTime.Scan(queryFilter.EndTime)

	arg := schedulepostgres.GetListTaskByScheduleIDParams{
		ScheduleID: sUUID,
		StartTime:  startTime,
		EndTime:    endTime,
	}
	res, err := r.Postgres.GetListTaskByScheduleID(ctx, arg)
	if err != nil {
		r.logger.Error(err.Error())
		return tasks, err
	}

	for _, task := range res {
		tasks = append(tasks, ent.TaskQuery{
			ID:           task.ID.String(),
			ScheduleID:   task.ScheduleID.String(),
			StartTime:    task.StartTime.Time,
			DurationDay:  calDurationDay(task.EndTime.Time, task.StartTime.Time),
			TaskDetailID: task.TaskDetailID.String(),
		})
	}
	return tasks, nil
}

func (r *PostgresScheduleRepository) GetListTaskWithDetailByScheduleID(
	ctx context.Context,
	scheduleID string,
) ([]ent.TaskWithDetailQuery, error) {
	scUUID, err := uuid.Parse(scheduleID)
	twds := []ent.TaskWithDetailQuery{}
	if err != nil {
		return twds, err
	}

	arg := schedulepostgres.GetListTaskAndDetailByScheduleIDParams{
		ScheduleID: scUUID,
	}
	res, err := r.Postgres.GetListTaskAndDetailByScheduleID(ctx, arg)
	if err != nil {
		return twds, err
	}

	for _, twd := range res {
		twds = append(twds, ent.TaskWithDetailQuery{
			TaskID:       twd.ID.String(),
			ScheduleID:   twd.ScheduleID.String(),
			StartTime:    twd.StartTime.Time,
			DurationDay:  calDurationDay(twd.EndTime.Time, twd.StartTime.Time),
			TaskDetailID: twd.TaskDetailID.String(),
			Name:         twd.Name.String,
			OwnerID:      twd.OwnerID.String,
			AssigneeID:   twd.AssigneeID.String,
			Status:       twd.Status.String,
		})
	}

	return twds, nil
}

func (r *PostgresScheduleRepository) CreateTask(
	ctx context.Context,
	task ent.TaskWithDetailCommand,
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	defer tx.Rollback()

	qtx := r.Postgres.WithTx(tx)
	tdUUID, err := qtx.CreateTaskDetail(ctx, schedulepostgres.CreateTaskDetailParams{
		Name: sql.NullString{
			String: task.Name,
			Valid:  len(task.Name) > 0,
		},
		OwnerID: sql.NullString{
			String: task.OwnerID,
			Valid:  len(task.OwnerID) > 0,
		},
		AssigneeID: sql.NullString{
			String: task.AssigneeID,
			Valid:  len(task.AssigneeID) > 0,
		},
		Status: sql.NullString{
			String: task.Status,
			Valid:  len(task.Status) > 0,
		},
	})

	if err != nil {
		r.logger.Error(err.Error())
		return errors.New("Failed to create task detail")
	}

	scheduleID, err := uuid.Parse(task.ScheduleID)
	if err != nil {
		r.logger.Error(err.Error())
		return errors.New("ScheduleID is in wrong format")
	}

	_, err = qtx.CreateTask(ctx, schedulepostgres.CreateTaskParams{
		ScheduleID: scheduleID,
		StartTime: sql.NullTime{
			Time:  task.StartTime,
			Valid: true,
		},
		EndTime: sql.NullTime{
			Time:  calEndTime(task.StartTime, task.DurationDay),
			Valid: true,
		},
		TaskDetailID: tdUUID,
	})

	if err != nil {
		r.logger.Error(err.Error())
		return errors.New("Failed to create task")
	}

	return tx.Commit()
}

func (r *PostgresScheduleRepository) EditTask(
	ctx context.Context,
	taskID string,
	task ent.TaskWithDetailCommand,
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	defer tx.Rollback()
	qtx := r.Postgres.WithTx(tx)

	tUUID, err := uuid.Parse(taskID)
	if err != nil {
		r.logger.Error(err.Error())
		return errors.New("Failed to parse task uuid")
	}

	td, err := qtx.GetTaskDetailByID(ctx, tUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return errors.New("Failed to get task detail")
	}

	_, err = r.Postgres.EditTask(ctx, schedulepostgres.EditTaskParams{
		ID: tUUID,
		StartTime: sql.NullTime{
			Time:  task.StartTime,
			Valid: true,
		},
		EndTime: sql.NullTime{
			Time:  calEndTime(task.StartTime, task.DurationDay),
			Valid: true,
		},
		TaskDetailID: td.ID,
	})

	if err != nil {
		r.logger.Error(err.Error())
		return errors.New("Failed to edit task")
	}

	return tx.Commit()
}

func calDurationDay(endTime time.Time, startTime time.Time) int32 {
	dif := endTime.Sub(startTime).Hours()
	return int32(math.Ceil(dif / 24))
}

func calEndTime(startTime time.Time, durationDay int32) time.Time {
	add := time.Duration(durationDay*24) * time.Hour
	return startTime.Add(add)
}
