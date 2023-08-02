package location

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	locationpostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/location/postgresql"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
)

type PostgresLocationRepository struct {
	db       *sql.DB
	Postgres *locationpostgres.Queries
	logger   zerolog.Logger
}

func NewPostgresLocationRepository(postgresRepo *database.SqlRepository) *PostgresLocationRepository {
	sublogger := log.With().Str("layer", "repository").Str("package", "location").Logger()

	queries := locationpostgres.New(postgresRepo.Db)
	return &PostgresLocationRepository{
		db:       postgresRepo.Db,
		Postgres: queries,
		logger:   sublogger,
	}
}

func (r *PostgresLocationRepository) GetAllBuildings(
	ctx context.Context,
) ([]ent.Building, error) {
	res, err := r.Postgres.GetAllBuildings(ctx)
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

func (r *PostgresLocationRepository) GetBuildingsByMemberId(
	ctx context.Context,
	memberId string,
) ([]ent.Building, error) {
	res, err := r.Postgres.GetBuildingsByMemberId(ctx, memberId)
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

func (r *PostgresLocationRepository) CreateMemberBuilding(
	ctx context.Context,
	memberId,
	buildingId string,
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error().Err(err)
		return err
	}
	defer tx.Rollback()
	qtx := r.Postgres.WithTx(tx)
	exist, err := qtx.GetMemberBuildingById(ctx, locationpostgres.GetMemberBuildingByIdParams{
		MemberID:   memberId,
		BuildingID: buildingId,
	})

	if err != nil {
		r.logger.Error().Err(err)
		return err
	}

	if exist {
		return errors.New("member building already created")
	}

	param := locationpostgres.CreateMemberBuildingParams{
		MemberID:   memberId,
		BuildingID: buildingId,
	}
	_, err = qtx.CreateMemberBuilding(ctx, param)
	if err != nil {
		r.logger.Error().Err(err)
		return err
	}

	return tx.Commit()
}

func (r *PostgresLocationRepository) DeleteBuilding(
	ctx context.Context,
	memberId,
	buildingId string,
) error {
	params := locationpostgres.DeleteMemberBuildingParams{
		MemberID:   memberId,
		BuildingID: buildingId,
	}
	err := r.Postgres.DeleteMemberBuilding(ctx, params)
	if err != nil {
		r.logger.Error().Err(err)
		return err
	}
	return nil
}

func (r *PostgresLocationRepository) GetMembersByRoomId(
	ctx context.Context,
	roomId string,
) ([]ent.Member, error) {
	res, err := r.Postgres.GetMembersByRoomId(ctx,
		sql.NullString{String: roomId, Valid: true},
	)
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

func (r *PostgresLocationRepository) GetRoomsByMemberId(
	ctx context.Context,
	memberId string,
) ([]ent.Room, error) {
	res, err := r.Postgres.GetRoomsByMemberId(ctx, memberId)
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

func (r *PostgresLocationRepository) GetRoomsByBuildingId(
	ctx context.Context,
	buildingId string,
) ([]ent.Room, error) {
	res, err := r.Postgres.GetRoomsByBuildingId(ctx, buildingId)
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
