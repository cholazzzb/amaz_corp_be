package schedule

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"strings"
	"time"

	"github.com/google/uuid"

	schedulepostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/schedule/postgresql"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/parser"
)

type PostgresScheduleRepository struct {
	db       *sql.DB
	Postgres *schedulepostgres.Queries
	logger   *slog.Logger
}

func NewPostgresScheduleRepository(postgresRepo *database.SqlRepository) *PostgresScheduleRepository {
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
	name,
	roomID string,
) (string, error) {
	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		r.logger.Error(err.Error())
		return "", err
	}
	res, err := r.Postgres.CreateScheduleByRoomID(ctx, schedulepostgres.CreateScheduleByRoomIDParams{
		Name:   name,
		RoomID: roomUUID,
	})

	if err != nil {
		r.logger.Error(err.Error())
		return "", fmt.Errorf("failed to create schedule with roomID: %s", roomID)
	}

	return res.String(), nil
}

func (r *PostgresScheduleRepository) GetListScheduleByRoomID(
	ctx context.Context,
	roomID string,
) ([]ent.ScheduleQuery, error) {
	schs := []ent.ScheduleQuery{}
	roomUUID, err := uuid.Parse(roomID)

	if err != nil {
		r.logger.Error(err.Error())
		return schs, err
	}
	res, err := r.Postgres.GetListScheduleByRoomID(ctx, roomUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return schs, err
	}

	for _, sch := range res {
		schs = append(schs, ent.ScheduleQuery{
			ID:     sch.ID.String(),
			Name:   sch.Name,
			RoomID: sch.RoomID.String(),
		})
	}

	return schs, nil
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
		OwnerID:    res.OwnerID.UUID.String(),
		AssigneeID: res.AssigneeID.UUID.String(),
		StatusID:   res.StatusID.UUID.String(),
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
	if queryFilter.StartTime != nil {
		startTime.Scan(*queryFilter.StartTime)
	}
	endTime := sql.NullTime{}
	if queryFilter.EndTime != nil {
		endTime.Scan(*queryFilter.EndTime)
	}

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
			Name:         task.Name.String,
			StartTime:    task.StartTime.Time,
			DurationDay:  calDurationDay(task.EndTime.Time, task.StartTime.Time),
			EndTime:      task.EndTime.Time,
			ScheduleID:   task.ScheduleID.String(),
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
		r.logger.Error(err.Error())
		return twds, err
	}

	res, err := r.Postgres.GetListTaskAndDetailByScheduleID(ctx, scUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return twds, err
	}

	for _, twd := range res {
		ownerID := parser.PostgresInterfaceToString(twd.OwnerID)
		assigneeID := parser.PostgresInterfaceToString(twd.AssigneeID)
		statusID := parser.PostgresInterfaceToString(twd.StatusID)
		deps := parser.PostgresInterfaceToString(twd.DependedTaskID)

		depsId := strings.Split(strings.Trim(deps, "{}"), ",")

		twds = append(twds, ent.TaskWithDetailQuery{
			TaskID:       twd.ID.String(),
			ScheduleID:   twd.ScheduleID.String(),
			StartTime:    twd.StartTime.Time,
			DurationDay:  calDurationDay(twd.EndTime.Time, twd.StartTime.Time),
			TaskDetailID: twd.TaskDetailID.String(),
			Name:         twd.Name.String,
			OwnerID:      ownerID,
			AssigneeID:   assigneeID,
			StatusID:     statusID,
			Dependencies: depsId,
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

	ownerID := uuid.NullUUID{}
	ownerUUID, err := uuid.Parse(task.OwnerID)
	if len(task.OwnerID) > 0 && err != nil {
		r.logger.Error(err.Error())
		return err
	}
	ownerID.Scan(ownerUUID)

	assigneeID := uuid.NullUUID{}
	assigneeUUID, err := uuid.Parse(task.AssigneeID)
	if len(task.AssigneeID) > 0 && err != nil {
		r.logger.Error(err.Error())
		return err
	}
	assigneeID.Scan(assigneeUUID)

	statusID := uuid.NullUUID{}
	statusUUIID, err := uuid.Parse(task.StatusID)
	if len(task.AssigneeID) > 0 && err != nil {
		r.logger.Error(err.Error())
		return err
	}
	statusID.Scan((statusUUIID))

	tdUUID, err := qtx.CreateTaskDetail(ctx, schedulepostgres.CreateTaskDetailParams{
		OwnerID:    ownerID,
		AssigneeID: assigneeID,
		StatusID:   statusID,
	})

	if err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("failed to create task detail %w", err)
	}

	scheduleID, err := uuid.Parse(task.ScheduleID)
	if err != nil {
		r.logger.Error(err.Error())
		return fmt.Errorf("scheduleID is in wrong format %w", err)
	}

	name := sql.NullString{}
	name.Scan(task.Name)
	startTime := sql.NullTime{}
	startTime.Scan(task.StartTime)
	endTime := sql.NullTime{}
	endTime.Scan(calEndTime(task.StartTime, task.DurationDay))

	_, err = qtx.CreateTask(ctx, schedulepostgres.CreateTaskParams{
		Name:         name,
		StartTime:    startTime,
		EndTime:      endTime,
		ScheduleID:   scheduleID,
		TaskDetailID: tdUUID,
	})

	if err != nil {
		return fmt.Errorf("can't retrieve taskID %w", err)
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
		return errors.New("failed to parse task uuid")
	}

	td, err := qtx.GetTaskDetailByID(ctx, tUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return errors.New("failed to get task detail")
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
		return errors.New("failed to edit task")
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
