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
) ([]ent.Building, error) {
	res, err := r.Postgres.GetAllBuildings(ctx)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.Building{}, err
	}

	bs := []ent.Building{}
	for _, mbs := range res {
		bs = append(bs, ent.Building{
			Id:   mbs.ID.String(),
			Name: mbs.Name,
		})
	}
	return bs, nil
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

func (r *PostgresLocationRepository) GetBuildingsByMemberId(
	ctx context.Context,
	memberID string,
) ([]ent.Building, error) {
	memberUUID, err := uuid.Parse(memberID)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.Building{}, err
	}

	res, err := r.Postgres.GetBuildingsByMemberId(ctx, memberUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.Building{}, err
	}

	bs := []ent.Building{}
	for _, mbs := range res {
		bs = append(bs, ent.Building{
			Id:   mbs.ID.String(),
			Name: mbs.Name,
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

	exist, err := qtx.GetUserBuildingExist(ctx, userID)

	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if exist {
		return errors.New("member building already exist")
	}
	fmt.Println("db", userID, buildingId)

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
	exist, err = qtx.GetMemberBuildingById(ctx, locationpostgres.GetMemberBuildingByIdParams{
		MemberID:   memberUUID,
		BuildingID: buildingUUID,
	})

	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if exist {
		return errors.New("member building already created")
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
