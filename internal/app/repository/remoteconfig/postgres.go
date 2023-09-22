package remoteconfig

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	remoteconfigpostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/remoteconfig/postgresql"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type PostgresRemoteConfigRepository struct {
	db       *sql.DB
	Postgres *remoteconfigpostgres.Queries
	logger   *slog.Logger
}

func NewPostgresRemoteConfigRepository(
	postgresRepo *database.SqlRepository,
) *PostgresRemoteConfigRepository {
	sublogger := logger.Get().With(
		slog.String("domain", "schedule"),
		slog.String("layer", "repo"),
	)
	queries := remoteconfigpostgres.New(postgresRepo.Db)

	return &PostgresRemoteConfigRepository{
		db:       postgresRepo.Db,
		Postgres: queries,
		logger:   sublogger,
	}
}

func (r *PostgresRemoteConfigRepository) CreateRemoteConfig(
	ctx context.Context,
	key string,
	value string,
) error {
	_, err := r.Postgres.CreateRemoteConfig(ctx, remoteconfigpostgres.CreateRemoteConfigParams{
		Key:   key,
		Value: value,
	})

	if err != nil {
		return fmt.Errorf("repo, remoteconfig. err:%w", err)
	}
	return nil
}

func (r *PostgresRemoteConfigRepository) UpdateRemoteConfig(
	ctx context.Context,
	key string,
	value string,
) error {
	_, err := r.Postgres.UpdateRemoteConfig(ctx, remoteconfigpostgres.UpdateRemoteConfigParams{
		Key:   key,
		Value: value,
	})

	if err != nil {
		return fmt.Errorf("repo, remoteconfig. err:%w", err)
	}
	return nil
}

func (r *PostgresRemoteConfigRepository) GetRemoteConfigByKey(
	ctx context.Context,
	key string,
) (string, error) {
	val, err := r.Postgres.GetRemoteConfigByKey(ctx, key)

	if err != nil {
		return "", fmt.Errorf("repo, remoteconfig. err:%w", err)
	}
	return val.Value, nil
}
