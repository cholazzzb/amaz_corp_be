package notification

import (
	"database/sql"
	"log/slog"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type PostgresNotificationRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewPostgresNotificationRepository(
	postgresRepo *database.SqlRepository,
) *PostgresNotificationRepository {
	sublogger := logger.Get().With(
		slog.String("domain", "notification"),
		slog.String("layer", "repo"),
	)

	return &PostgresNotificationRepository{
		db:     postgresRepo.Db,
		logger: sublogger,
	}
}
