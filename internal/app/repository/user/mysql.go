package user

import (
	"context"
	"log/slog"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"

	mysql "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user/mysql"
)

type MySQLUserRepository struct {
	Mysql  *mysql.Queries
	logger *slog.Logger
}

func NewMySQLUserRepository(sqlRepo *database.SqlRepository) *MySQLUserRepository {
	sublogger := logger.Get().With(slog.String("domain", "user"), slog.String("layer", "repo"))

	queries := mysql.New(sqlRepo.Db)
	return &MySQLUserRepository{Mysql: queries, logger: sublogger}
}

func (r *MySQLUserRepository) GetUser(
	ctx context.Context,
	params string,
) (user.User, error) {
	result, err := r.Mysql.GetUser(ctx, params)
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

func (r *MySQLUserRepository) GetUserExistance(
	ctx context.Context,
	username string,
) (bool, error) {
	exist, err := r.Mysql.GetUserExistance(ctx, username)
	if err != nil {
		r.logger.Error(err.Error())
		return true, err
	}
	return exist, nil
}

func (r *MySQLUserRepository) CreateUser(
	ctx context.Context,
	params user.User,
) error {
	_, err := r.Mysql.CreateUser(ctx, mysql.CreateUserParams{
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
