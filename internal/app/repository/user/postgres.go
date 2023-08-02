package user

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"

	userpostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user/postgresql"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
)

type PostgresUserRepository struct {
	Postgres *userpostgres.Queries
	logger   zerolog.Logger
}

func NewPostgresUserRepository(sqlRepo *database.SqlRepository) *PostgresUserRepository {
	sublogger := log.With().Str("layer", "repository").Str("package", "user").Logger()

	queries := userpostgres.New(sqlRepo.Db)
	return &PostgresUserRepository{Postgres: queries, logger: sublogger}
}

func (r *PostgresUserRepository) GetUser(
	ctx context.Context,
	params string,
) (user.User, error) {
	result, err := r.Postgres.GetUser(ctx, params)
	if err != nil {
		r.logger.Err(err)
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
		r.logger.Error().Err(err)
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
		r.logger.Error().Err(err)
		return err
	}
	return nil
}

func (r *PostgresUserRepository) GetMemberByName(
	ctx context.Context,
	memberName string,
) (user.Member, error) {
	result, err := r.Postgres.GetMemberByName(ctx, memberName)
	if err != nil {
		r.logger.Error().Err(err)
		return user.Member{}, err
	}
	return user.Member{
		Name:   result.Name,
		Status: result.Status,
	}, nil
}

func (r *PostgresUserRepository) CreateMemberParams(
	newMember user.Member,
	userID string,
) userpostgres.CreateMemberParams {
	return userpostgres.CreateMemberParams{
		ID:     newMember.ID,
		Name:   newMember.Name,
		Status: newMember.Status,
		UserID: userID,
	}
}

func (r *PostgresUserRepository) CreateMember(
	ctx context.Context,
	newMember user.Member,
	userID string,
) (user.Member, error) {
	params := r.CreateMemberParams(newMember, userID)
	_, err := r.Postgres.CreateMember(ctx, params)
	if err != nil {
		r.logger.Error().Err(err)
		return user.Member{}, err
	}
	return newMember, nil
}

func (r *PostgresUserRepository) GetFriendsByUserId(
	ctx context.Context,
	userId string,
) ([]user.Member, error) {
	fs, err := r.Postgres.GetFriendsByMemberId(ctx, userId)
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

func (r *PostgresUserRepository) CreateFriend(
	ctx context.Context,
	member1Id,
	member2Id string,
) error {
	return errors.New("")
}
