package location

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	locationpostgres "github.com/cholazzzb/amaz_corp_be/internal/app/repository/location/postgresql"
	"github.com/cholazzzb/amaz_corp_be/internal/datastore/database"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type PostgresLocationRepository struct {
	db       *sql.DB
	Postgres *locationpostgres.Queries
	logger   *slog.Logger
}

func NewPostgresLocationRepository(postgresRepo *database.SqlRepository) *PostgresLocationRepository {
	sublogger := logger.Get().With(slog.String("domain", "location"), slog.String("layer", "repo"))
	queries := locationpostgres.New(postgresRepo.Db)

	return &PostgresLocationRepository{
		db:       postgresRepo.Db,
		Postgres: queries,
		logger:   sublogger,
	}
}

func (r *PostgresLocationRepository) GetAllBuildings(
	ctx context.Context,
) ([]ent.BuildingQuery, error) {
	res, err := r.Postgres.GetAllBuildings(ctx)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.BuildingQuery{}, err
	}

	bs := []ent.BuildingQuery{}
	for _, mbs := range res {
		bs = append(bs, ent.BuildingQuery{
			ID:   mbs.ID.String(),
			Name: mbs.Name,
		})
	}
	return bs, nil
}

func (r *PostgresLocationRepository) GetMemberBuildingExist(
	ctx context.Context,
	userID,
	buildingID string,
) (bool, error) {
	buildingUUID, err := uuid.Parse(buildingID)
	if err != nil {
		r.logger.Error(err.Error())
		return true, err
	}

	exist, err := r.Postgres.GetUserBuildingExist(ctx, locationpostgres.GetUserBuildingExistParams{
		ID:         userID,
		BuildingID: buildingUUID,
	})

	if err != nil {
		r.logger.Error(err.Error())
		return true, err
	}

	if exist {
		return true, nil
	}

	return false, nil
}

func (r *PostgresLocationRepository) GetMemberByName(
	ctx context.Context,
	memberName string,
) (ent.MemberQuery, error) {
	result, err := r.Postgres.GetMemberByName(ctx, memberName)
	if err != nil {
		r.logger.Error(err.Error())
		return ent.MemberQuery{}, err
	}
	return ent.MemberQuery{
		ID:     result.ID.String(),
		UserID: result.UserID,
		Name:   result.Name,
		Status: result.Status,
		RoomID: result.RoomID.UUID.String(),
	}, nil
}

func (r *PostgresLocationRepository) GetFriendsByUserId(
	ctx context.Context,
	userId string,
) ([]ent.MemberQuery, error) {
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	fs, err := r.Postgres.GetFriendsByMemberId(ctx, userUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	result := make([]ent.MemberQuery, len(fs))
	for i, friend := range fs {
		result[i] = ent.MemberQuery{
			Name:   friend.Name,
			Status: friend.Status,
		}
	}
	return result, nil
}

func (r *PostgresLocationRepository) CreateFriend(
	ctx context.Context,
	member1Id,
	member2Id string,
) error {
	return errors.New("")
}

func (r *PostgresLocationRepository) GetBuildingsByUserID(
	ctx context.Context,
	userID string,
) ([]ent.BuildingQuery, error) {
	res, err := r.Postgres.GetBuildingsByUserID(ctx, userID)

	bs := []ent.BuildingQuery{}
	if err != nil {
		r.logger.Error(err.Error())
		return bs, err
	}

	for _, ubs := range res {
		bs = append(bs, ent.BuildingQuery{
			ID:   ubs.ID.String(),
			Name: ubs.Name,
		})
	}
	return bs, nil
}

func (r *PostgresLocationRepository) GetListMemberByBuildingID(
	ctx context.Context,
	buildingID string,
) ([]ent.MemberQuery, error) {
	buildingUUID, err := uuid.Parse(buildingID)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.MemberQuery{}, err
	}

	res, err := r.Postgres.GetListMemberByBuildingID(ctx, buildingUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.MemberQuery{}, err
	}

	ms := []ent.MemberQuery{}
	for _, m := range res {
		ms = append(ms, ent.MemberQuery{
			ID:     m.ID.String(),
			Name:   m.Name,
			Status: m.Status,
		})
	}

	return ms, nil
}

func (r *PostgresLocationRepository) CreateMemberBuilding(
	ctx context.Context,
	memberName,
	userID,
	buildingId string,
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	defer tx.Rollback()
	qtx := r.Postgres.WithTx(tx)

	memberUUID, err := qtx.CreateMember(ctx, locationpostgres.CreateMemberParams{
		Name:   memberName,
		Status: "new member",
		UserID: userID,
	})

	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	buildingUUID, err := uuid.Parse(buildingId)

	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	param := locationpostgres.CreateMemberBuildingParams{
		MemberID:   memberUUID,
		BuildingID: buildingUUID,
	}
	_, err = qtx.CreateMemberBuilding(ctx, param)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return tx.Commit()
}

func (r *PostgresLocationRepository) DeleteBuilding(
	ctx context.Context,
	memberId,
	buildingId string,
) error {
	memberUUID, err := uuid.Parse(memberId)

	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	buildingUUID, err := uuid.Parse(buildingId)

	fmt.Println("dbss", memberId, buildingId)

	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	params := locationpostgres.DeleteMemberBuildingParams{
		MemberID:   memberUUID,
		BuildingID: buildingUUID,
	}
	err = r.Postgres.DeleteMemberBuilding(ctx, params)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	return nil
}

func (r *PostgresLocationRepository) GetMembersByRoomId(
	ctx context.Context,
	roomId string,
) ([]ent.Member, error) {
	roomUUID, err := uuid.Parse(roomId)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.Member{}, nil
	}

	roomID := uuid.NullUUID{}
	roomID.Scan(roomUUID)

	res, err := r.Postgres.GetMembersByRoomId(ctx,
		roomID,
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

func (r *PostgresLocationRepository) GetRoomsByMemberId(
	ctx context.Context,
	memberId string,
) ([]ent.Room, error) {
	memberUUID, err := uuid.Parse(memberId)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.Room{}, nil
	}

	res, err := r.Postgres.GetRoomsByMemberId(ctx, memberUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.Room{}, nil
	}

	rs := []ent.Room{}
	for _, mrs := range res {
		rs = append(rs, ent.Room{
			Id:   mrs.ID.String(),
			Name: mrs.Name,
		})
	}
	return rs, nil
}

func (r *PostgresLocationRepository) GetRoomsByBuildingId(
	ctx context.Context,
	buildingId string,
) ([]ent.Room, error) {
	buildingUUID, err := uuid.Parse(buildingId)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.Room{}, nil
	}

	res, err := r.Postgres.GetRoomsByBuildingId(ctx, buildingUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.Room{}, nil
	}

	rs := []ent.Room{}
	for _, mrs := range res {
		rs = append(rs, ent.Room{
			Id:   mrs.ID.String(),
			Name: mrs.Name,
		})
	}
	return rs, nil
}
