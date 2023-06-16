package location

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	mysql "github.com/cholazzzb/amaz_corp_be/internal/app/repository/location/mysql"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
)

type MySQLLocationRepository struct {
	Mysql  *mysql.Queries
	logger zerolog.Logger
}

func NewMySQLLocationRepository(
	mysqlRepo *database.MysqlRepository,
) *MySQLLocationRepository {
	sublogger := log.With().
		Str("layer", "repository").
		Str("package", "location").Logger()

	queries := mysql.New(mysqlRepo.Db)
	return &MySQLLocationRepository{
		Mysql:  queries,
		logger: sublogger,
	}
}

func (r *MySQLLocationRepository) GetAllBuildings(
	ctx context.Context,
) ([]ent.Building, error) {
	res, err := r.Mysql.GetAllBuildings(ctx)
	if err != nil {
		r.logger.Error().Err(err)
		return []ent.Building{}, err
	}

	bs := []ent.Building{}
	for _, mbs := range res {
		bs = append(bs, ent.Building{
			Id:   mbs.ID,
			Name: mbs.Name,
		})
	}
	return bs, nil
}

func (r *MySQLLocationRepository) GetBuildingsByMemberId(
	ctx context.Context,
	memberId int64,
) ([]ent.Building, error) {
	res, err := r.Mysql.GetBuildingsByMemberId(ctx, memberId)
	if err != nil {
		r.logger.Error().Err(err)
		return []ent.Building{}, nil
	}

	bs := []ent.Building{}
	for _, mbs := range res {
		bs = append(bs, ent.Building{
			Id:   mbs.ID,
			Name: mbs.Name,
		})
	}
	return bs, nil
}

func (r *MySQLLocationRepository) CreateMemberBuilding(
	ctx context.Context,
	memberId int64,
	buildingId int64,
) error {
	param := mysql.CreateMemberBuildingParams{
		MemberID:   memberId,
		BuildingID: buildingId,
	}
	_, err := r.Mysql.CreateMemberBuilding(ctx, param)
	if err != nil {
		r.logger.Error().Err(err)
		return err
	}
	return nil
}

func (r *MySQLLocationRepository) GetMembersByRoomId(
	ctx context.Context,
	roomId int64,
) ([]ent.Member, error) {
	res, err := r.Mysql.GetMembersByRoomId(ctx, sql.NullInt64{
		Int64: roomId, Valid: true,
	})
	if err != nil {
		r.logger.Error().Err(err)
		return []ent.Member{}, nil
	}

	ms := []ent.Member{}

	for _, mms := range res {
		ms = append(ms, ent.Member{
			Name:   mms.Name,
			Status: mms.Status,
			UserId: mms.UserID,
		})
	}
	return ms, nil
}

func (r *MySQLLocationRepository) GetRoomsByMemberId(
	ctx context.Context,
	memberId int64,
) ([]ent.Room, error) {
	res, err := r.Mysql.GetRoomsByMemberId(ctx, memberId)
	if err != nil {
		r.logger.Error().Err(err)
		return []ent.Room{}, nil
	}

	rs := []ent.Room{}
	for _, mrs := range res {
		rs = append(rs, ent.Room{
			Id:   mrs.ID,
			Name: mrs.Name,
		})
	}
	return rs, nil
}

func (r *MySQLLocationRepository) GetRoomsByBuildingId(
	ctx context.Context,
	buildingId int64,
) ([]ent.Room, error) {
	res, err := r.Mysql.GetRoomsByBuildingId(ctx, buildingId)
	if err != nil {
		r.logger.Error().Err(err)
		return []ent.Room{}, nil
	}

	rs := []ent.Room{}
	for _, mrs := range res {
		rs = append(rs, ent.Room{
			Id:   mrs.ID,
			Name: mrs.Name,
		})
	}
	return rs, nil
}
