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

type Repository interface {
	GetFriendByUserId(ctx context.Context, userId int64) ([]member.Member, error)
	CreateFriend(ctx context.Context, member1Id, member2Id int64) error
}

type MySQLFriendRepository struct {
	mysql  *mysql.Queries
	logger zerolog.Logger
}

func NewMySQLFriendRepository(mysqlRepo *database.MysqlRepository) *MySQLFriendRepository {
	sublogger := log.With().Str("layer", "repository").Str("package", "member").Logger()
	// create tables only if it not exists before (see schema.sql)
	ctx := context.Background()
	if _, err := mysqlRepo.Db.ExecContext(ctx, database.DdlMember); err != nil {
		sublogger.Panic().Err(err).Msg("failed to create table members")
	}

	queries := mysql.New(mysqlRepo.Db)
	return &MySQLFriendRepository{mysql: queries, logger: sublogger}
}

func (r *MySQLFriendRepository) GetFriendByUserId(ctx context.Context, userId int64) ([]member.Member, error) {
	return []member.Member{}, errors.New("")
}

func (r *MySQLFriendRepository) CreateFriend(ctx context.Context, member1Id, member2Id int64) error {
	return errors.New("")
}
