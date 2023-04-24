package member

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/member"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/member/mysql"
)

type Repository interface {
	GetMemberByName(ctx context.Context, memberName string) (member.Member, error)
	CreateMember(ctx context.Context, member member.Member) error
}

type MySQLMemberRepository struct {
	mysql  *mysql.Queries
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
	return &MySQLMemberRepository{mysql: queries, logger: sublogger}
}

func (r *MySQLMemberRepository) GetMemberByName(ctx context.Context, memberName string) (member.Member, error) {
	result, err := r.mysql.GetMemberByName(ctx, memberName)
	if err != nil {
		return member.Member{}, err
	}
	return member.Member{
		Name:   result.Name,
		Status: result.Status,
	}, nil

}

func (r *MySQLMemberRepository) CreateMember(ctx context.Context, member member.Member, userID int64) error {
	params := mysql.CreateMemberParams{
		Name:   member.Name,
		Status: member.Status,
		UserID: sql.NullInt64{
			Int64: userID,
			Valid: true,
		},
	}
	_, err := r.mysql.CreateMember(ctx, params)
	if err != nil {
		r.logger.Error().Err(err)
		return err
	}
	return nil
}
