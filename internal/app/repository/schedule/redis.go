package schedule

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-redis/cache/v9"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type RedisScheduleRepository struct {
	sc     *cache.Cache
	logger *slog.Logger
}

func NewRedisScheduleRepository(rds *database.RedisRepository) *RedisScheduleRepository {
	sc := cache.New(&cache.Options{
		Redis: rds.Rds,
	})
	sublogger := logger.Get().With(slog.String("domain", "schedule"), slog.String("layer", "repo-cache"))

	return &RedisScheduleRepository{
		sc:     sc,
		logger: sublogger,
	}
}

func (r *RedisScheduleRepository) SaveAutoSchedule(
	ctx context.Context,
	scheduleID string,
	tasks []ent.TaskWithDetailQuery,
) error {
	if err := r.sc.Set(&cache.Item{
		Ctx:   ctx,
		Key:   scheduleID,
		Value: tasks,
		TTL:   time.Hour,
	}); err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *RedisScheduleRepository) InvalidateAutoSchedule(
	ctx context.Context,
	scheduleID string,
) error {
	if err := r.sc.Delete(ctx, scheduleID); err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *RedisScheduleRepository) GetAutoSchedule(
	ctx context.Context,
	scheduleID string,
) ([]ent.TaskWithDetailQuery, error) {
	var tasks []ent.TaskWithDetailQuery
	if err := r.sc.Get(ctx, scheduleID, tasks); err != nil {
		r.logger.Error(err.Error())
		return []ent.TaskWithDetailQuery{}, err
	}

	return tasks, nil
}
