package friend

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/friend/mysql"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/member"
)

type MySQLFriendRepository struct {
	Mysql  *mysql.Queries
	logger zerolog.Logger
}

func NewMySQLFriendRepository(mysqlRepo *database.MysqlRepository) *MySQLFriendRepository {
	sublogger := log.With().Str("layer", "repository").Str("package", "member").Logger()
	// create tables only if it not exists before (see schema.sql)
	ctx := context.Background()
	// TODO: Change Migration to other options
	if _, err := mysqlRepo.Db.ExecContext(ctx, database.DdlMember); err != nil {
		sublogger.Panic().Err(err).Msg("failed to create table members")
	}

	queries := mysql.New(mysqlRepo.Db)
	return &MySQLFriendRepository{Mysql: queries, logger: sublogger}
}

func (r *MySQLFriendRepository) GetFriendsByUserId(ctx context.Context, userId int64) ([]member.Member, error) {
	fs, err := r.Mysql.GetFriendsByMemberId(ctx, mysql.GetFriendsByMemberIdParams{
		Member1ID: userId,
		Member2ID: userId,
		ID:        userId,
	})
	if err != nil {
		r.logger.Error().Err(err)
		return nil, err
	}
	result := make([]member.Member, len(fs))
	for i, friend := range fs {
		result[i] = member.Member{
			Name:   friend.Name,
			Status: friend.Status,
		}
	}
	return result, nil
}

func (r *MySQLFriendRepository) CreateFriend(ctx context.Context, member1Id, member2Id int64) error {
	return errors.New("")
}
