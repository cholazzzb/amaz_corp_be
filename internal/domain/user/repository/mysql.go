package user

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/user/mysql"
)

type MySQLUserRepository struct {
	Mysql  *mysql.Queries
	logger zerolog.Logger
}

func NewMySQLUserRepository(mysqlRepo *database.MysqlRepository) *MySQLUserRepository {
	sublogger := log.With().Str("layer", "repository").Str("package", "user").Logger()
	// create tables only if it not exists before (see schema.sql)
	ctx := context.Background()
	if _, err := mysqlRepo.Db.ExecContext(ctx, database.DdlUser); err != nil {
		sublogger.Panic().Err(err).Msg("failed to create table users")
	}

	queries := mysql.New(mysqlRepo.Db)
	return &MySQLUserRepository{Mysql: queries, logger: sublogger}
}

func (r *MySQLUserRepository) GetUser(ctx context.Context, params string) (mysql.User, error) {
	result, err := r.Mysql.GetUser(ctx, params)
	if err != nil {
		r.logger.Err(err)
		return mysql.User{}, err
	}
	return result, nil
}

func (r *MySQLUserRepository) CreateUser(ctx context.Context, params mysql.CreateUserParams) error {
	_, err := r.Mysql.CreateUser(ctx, params)
	if err != nil {
		r.logger.Error().Err(err)
		return err
	}
	return nil
}
