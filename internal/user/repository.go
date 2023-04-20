package user

import (
	"context"
	"log"

	"github.com/cholazzzb/amaz_corp_be/internal/database"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/user/mysql"
)

type UserRepository struct {
	mysql *mysql.Queries
}

func NewUserRepository(mysqlRepo *database.MysqlRepository) *UserRepository {
	// create tables only if it not exists before (see schema.sql)
	ctx := context.Background()
	if _, err := mysqlRepo.Db.ExecContext(ctx, database.DdlUser); err != nil {
		log.Panicln("failed to create table users", err)
	}

	queries := mysql.New(mysqlRepo.Db)
	return &UserRepository{mysql: queries}
}

func (r *UserRepository) GetUser(ctx context.Context, params string) (mysql.User, error) {
	result, err := r.mysql.GetUser(ctx, params)
	if err != nil {
		return mysql.User{}, err
	}
	return result, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, params mysql.CreateUserParams) error {
	_, err := r.mysql.CreateUser(ctx, params)
	if err != nil {
		log.Println("repo/mysql/user CreateUserErr:", err)
		return err
	}
	return nil
}
