package location

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	mysql "github.com/cholazzzb/amaz_corp_be/internal/app/repository/location/mysql"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type MySQLLocationRepository struct {
	db     *sql.DB
	Mysql  *mysql.Queries
	logger *slog.Logger
}

func NewMySQLLocationRepository(
	sqlRepo *database.SqlRepository,
) *MySQLLocationRepository {
	sublogger := logger.Get().With(slog.String("domain", "location"), slog.String("layer", "repo"))

	queries := mysql.New(sqlRepo.Db)
	return &MySQLLocationRepository{
		db:     sqlRepo.Db,
		Mysql:  queries,
		logger: sublogger,
	}
}

func (r *MySQLLocationRepository) GetAllBuildings(
	ctx context.Context,
) ([]ent.BuildingQuery, error) {
	res, err := r.Mysql.GetAllBuildings(ctx)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.BuildingQuery{}, err
	}

	bs := []ent.BuildingQuery{}
	for _, mbs := range res {
		bs = append(bs, ent.BuildingQuery{
			ID:   mbs.ID,
			Name: mbs.Name,
		})
	}
	return bs, nil
}

func (r *MySQLLocationRepository) GetBuildingsByMemberId(
	ctx context.Context,
	memberId string,
) ([]ent.BuildingQuery, error) {
	res, err := r.Mysql.GetBuildingsByMemberId(ctx, memberId)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.BuildingQuery{}, err
	}

	bs := []ent.BuildingQuery{}
	for _, mbs := range res {
		bs = append(bs, ent.BuildingQuery{
			ID:   mbs.ID,
			Name: mbs.Name,
		})
	}
	return bs, nil
}

func (r *MySQLLocationRepository) CreateMemberBuilding(
	ctx context.Context,
	memberId,
	buildingId string,
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	defer tx.Rollback()
	qtx := r.Mysql.WithTx(tx)
	exist, err := qtx.GetMemberBuildingById(ctx, mysql.GetMemberBuildingByIdParams{
		MemberID:   memberId,
		BuildingID: buildingId,
	})

	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if exist {
		return errors.New("member building already created")
	}

	param := mysql.CreateMemberBuildingParams{
		MemberID:   memberId,
		BuildingID: buildingId,
	}
	_, err = qtx.CreateMemberBuilding(ctx, param)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return tx.Commit()
}

func (r *MySQLLocationRepository) DeleteBuilding(
	ctx context.Context,
	memberId,
	buildingId string,
) error {
	params := mysql.DeleteMemberBuildingParams{
		MemberID:   memberId,
		BuildingID: buildingId,
	}
	err := r.Mysql.DeleteMemberBuilding(ctx, params)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}

func (r *MySQLLocationRepository) GetMembersByRoomId(
	ctx context.Context,
	roomId string,
) ([]ent.Member, error) {
	res, err := r.Mysql.GetMembersByRoomId(ctx,
		sql.NullString{String: roomId, Valid: true},
	)
	if err != nil {
		r.logger.Error(err.Error())
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
	memberId string,
) ([]ent.RoomQuery, error) {
	res, err := r.Mysql.GetRoomsByMemberId(ctx, memberId)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.RoomQuery{}, nil
	}

	rs := []ent.RoomQuery{}
	for _, mrs := range res {
		rs = append(rs, ent.RoomQuery{
			Id:   mrs.ID,
			Name: mrs.Name,
		})
	}
	return rs, nil
}

func (r *MySQLLocationRepository) GetRoomsByBuildingId(
	ctx context.Context,
	buildingId string,
) ([]ent.RoomQuery, error) {
	res, err := r.Mysql.GetRoomsByBuildingId(ctx, buildingId)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.RoomQuery{}, nil
	}

	rs := []ent.RoomQuery{}
	for _, mrs := range res {
		rs = append(rs, ent.RoomQuery{
			Id:   mrs.ID,
			Name: mrs.Name,
		})
	}
	return rs, nil
}
