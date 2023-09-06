package user

import (
	"context"
	"log/slog"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"

	userpostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user/postgresql"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
)

type PostgresUserRepository struct {
	Postgres *userpostgres.Queries
	logger   *slog.Logger
}

func NewPostgresUserRepository(sqlRepo *database.SqlRepository) *PostgresUserRepository {
	sublogger := logger.Get().With("domain", "user", "layer", "repo")

	queries := userpostgres.New(sqlRepo.Db)
	return &PostgresUserRepository{Postgres: queries, logger: sublogger}
}

func (r *PostgresUserRepository) GetUser(
	ctx context.Context,
	params string,
) (user.User, error) {
	result, err := r.Postgres.GetUser(ctx, params)
	if err != nil {
		r.logger.Error(err.Error())
		return user.User{}, err
	}
	return user.User{
		ID:       result.ID,
		Username: result.Username,
		Password: result.Password,
		Salt:     result.Salt,
	}, nil
}

func (r *PostgresUserRepository) GetUserExistance(
	ctx context.Context,
	username string,
) (bool, error) {
	exist, err := r.Postgres.GetUserExistance(ctx, username)
	if err != nil {
		r.logger.Error(err.Error())
		return true, err
	}
	return exist, nil
}

func (r *PostgresUserRepository) CreateUser(
	ctx context.Context,
	params user.User,
) error {
	_, err := r.Postgres.CreateUser(ctx, userpostgres.CreateUserParams{
		ID:       params.ID,
		Username: params.Username,
		Password: params.Password,
		Salt:     params.Salt,
	})
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}
