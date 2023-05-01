package member

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/member"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/member/mysql"
)

type MySQLMemberRepository struct {
	Mysql  *mysql.Queries
	logger zerolog.Logger
}

func NewMySQLMemberRepository(mysqlRepo *database.MysqlRepository) *MySQLMemberRepository {
	sublogger := log.With().Str("layer", "repository").Str("package", "member").Logger()
	// create tables only if it not exists before (see schema.sql)
	ctx := context.Background()
	if _, err := mysqlRepo.Db.ExecContext(ctx, database.DdlMember); err != nil {
		sublogger.Panic().Err(err).Msg("failed to create table members")
	}

	queries := mysql.New(mysqlRepo.Db)
	return &MySQLMemberRepository{Mysql: queries, logger: sublogger}
}

func (r *MySQLMemberRepository) GetMemberByName(ctx context.Context, memberName string) (member.Member, error) {
	result, err := r.Mysql.GetMemberByName(ctx, memberName)
	if err != nil {
		r.logger.Error().Err(err)
		return member.Member{}, err
	}
	return member.Member{
		Name:   result.Name,
		Status: result.Status,
	}, nil

}

func (r *MySQLMemberRepository) CreateMemberParams(newMember member.Member, userID int64) mysql.CreateMemberParams {
	return mysql.CreateMemberParams{
		Name:   newMember.Name,
		Status: newMember.Status,
		UserID: userID,
	}
}

func (r *MySQLMemberRepository) CreateMember(ctx context.Context, newMember member.Member, userID int64) (member.Member, error) {
	params := r.CreateMemberParams(newMember, userID)
	_, err := r.Mysql.CreateMember(ctx, params)
	if err != nil {
		r.logger.Error().Err(err)
		return member.Member{}, err
	}
	return newMember, nil
}
