package user

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"

	mysql "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user/mysql"
)

type MySQLUserRepository struct {
	Mysql  *mysql.Queries
	logger zerolog.Logger
}

func NewMySQLUserRepository(mysqlRepo *database.MysqlRepository) *MySQLUserRepository {
	sublogger := log.With().Str("layer", "repository").Str("package", "user").Logger()

	queries := mysql.New(mysqlRepo.Db)
	return &MySQLUserRepository{Mysql: queries, logger: sublogger}
}

func (r *MySQLUserRepository) GetUser(
	ctx context.Context,
	params string,
) (mysql.User, error) {
	result, err := r.Mysql.GetUser(ctx, params)
	if err != nil {
		r.logger.Err(err)
		return mysql.User{}, err
	}
	return result, nil
}

func (r *MySQLUserRepository) GetUserExistance(
	ctx context.Context,
	username string,
) (bool, error) {
	exist, err := r.Mysql.GetUserExistance(ctx, username)
	if err != nil {
		r.logger.Error().Err(err)
		return true, err
	}
	return exist, nil
}

func (r *MySQLUserRepository) CreateUser(
	ctx context.Context,
	params mysql.CreateUserParams,
) error {
	_, err := r.Mysql.CreateUser(ctx, params)
	if err != nil {
		r.logger.Error().Err(err)
		return err
	}
	return nil
}

func (r *MySQLUserRepository) GetMemberByName(
	ctx context.Context,
	memberName string,
) (user.Member, error) {
	result, err := r.Mysql.GetMemberByName(ctx, memberName)
	if err != nil {
		r.logger.Error().Err(err)
		return user.Member{}, err
	}
	return user.Member{
		Name:   result.Name,
		Status: result.Status,
	}, nil
}

func (r *MySQLUserRepository) CreateMemberParams(
	newMember user.Member,
	userID string,
) mysql.CreateMemberParams {
	return mysql.CreateMemberParams{
		ID:     newMember.Id,
		Name:   newMember.Name,
		Status: newMember.Status,
		UserID: userID,
	}
}

func (r *MySQLUserRepository) CreateMember(
	ctx context.Context,
	newMember user.Member,
	userID string,
) (user.Member, error) {
	params := r.CreateMemberParams(newMember, userID)
	_, err := r.Mysql.CreateMember(ctx, params)
	if err != nil {
		r.logger.Error().Err(err)
		return user.Member{}, err
	}
	return newMember, nil
}

func (r *MySQLUserRepository) GetFriendsByUserId(
	ctx context.Context,
	userId string,
) ([]user.Member, error) {
	fs, err := r.Mysql.GetFriendsByMemberId(ctx, mysql.GetFriendsByMemberIdParams{
		Member1ID: userId,
		Member2ID: userId,
		ID:        userId,
	})
	if err != nil {
		r.logger.Error().Err(err)
		return nil, err
	}
	result := make([]user.Member, len(fs))
	for i, friend := range fs {
		result[i] = user.Member{
			Name:   friend.Name,
			Status: friend.Status,
		}
	}
	return result, nil
}

func (r *MySQLUserRepository) CreateFriend(
	ctx context.Context,
	member1Id,
	member2Id string,
) error {
	return errors.New("")
}
