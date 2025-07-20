package repo_user

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/cholazzzb/amaz_corp_be/internal/app/model"
	repository_user "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user/postgresql"
	database "github.com/cholazzzb/amaz_corp_be/internal/datastore"

	custom_logger "github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type UserRepo interface {
	UserCommandRepo
	UserQueryRepo
}

type UserCommandRepo interface {
	Register(
		ctx context.Context,
		user *model.RegisterUserCommand,
	) error
	Update(
		ctx context.Context,
		id string,
		user *model.UserProfileCommand,
	) (*model.UserProfileCommandRes, error)
	Deactivate(
		ctx context.Context,
		id string,
	) error
}

type UserQueryRepo interface{}

type UserRepository struct {
	db     *sql.DB
	Pg     *repository_user.Queries
	logger *slog.Logger
}

func NewRepository(sqlRepo *database.SqlRepository) *UserRepository {
	sublogger := custom_logger.Get().With(slog.String("domain", "user"), slog.String("layer", "repo"))
	queries := repository_user.New(sqlRepo.Db)

	return &UserRepository{
		db:     sqlRepo.Db,
		Pg:     queries,
		logger: sublogger,
	}
}
