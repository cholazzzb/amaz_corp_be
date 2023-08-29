package user

import (
	"context"
	"errors"
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

func (r *MySQLUserRepository) GetMemberByName(
	ctx context.Context,
	memberName string,
) (user.Member, error) {
	result, err := r.Mysql.GetMemberByName(ctx, memberName)
	if err != nil {
		r.logger.Error(err.Error())
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
		ID:     newMember.ID,
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
		r.logger.Error(err.Error())
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
		r.logger.Error(err.Error())
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
