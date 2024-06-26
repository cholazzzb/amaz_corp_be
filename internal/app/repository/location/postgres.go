package location

import (
	"context"
	"database/sql"
	"errors"
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

func (r *PostgresLocationRepository) GetBuildingByID(
	ctx context.Context,
	buildingID string,
) (ent.BuildingQuery, error) {
	buildingUUID, err := uuid.Parse(buildingID)
	if err != nil {
		r.logger.Error(err.Error())
		return ent.BuildingQuery{}, err
	}

	res, err := r.Postgres.GetBuildingByID(ctx, buildingUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return ent.BuildingQuery{}, err
	}

	return ent.BuildingQuery{
		ID:   res.ID.String(),
		Name: res.Name,
	}, nil
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

func (r *PostgresLocationRepository) GetInvitationByUserID(
	ctx context.Context,
	userID string,
) ([]ent.BuildingMemberQuery, error) {
	res := []ent.BuildingMemberQuery{}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		r.logger.Error(err.Error())
		return res, err
	}

	bldngs, err := r.Postgres.GetInvitationByUserID(ctx, userUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return res, err
	}

	for _, bldng := range bldngs {
		res = append(res, ent.BuildingMemberQuery{
			BuildingID:   bldng.ID.String(),
			BuildingName: bldng.Name,
			MemberID:     bldng.MemberID.String(),
		})
	}

	return res, nil
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

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		r.logger.Error(err.Error())
		return true, err
	}

	exist, err := r.Postgres.GetUserBuildingExist(ctx, locationpostgres.GetUserBuildingExistParams{
		ID:         userUUID,
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

func (r *PostgresLocationRepository) GetListBuildingByUserID(
	ctx context.Context,
	userID string,
) ([]ent.BuildingMemberQuery, error) {
	bs := []ent.BuildingMemberQuery{}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		r.logger.Error(err.Error())
		return bs, err
	}
	res, err := r.Postgres.GetListBuildingByUserID(ctx, userUUID)

	if err != nil {
		r.logger.Error(err.Error())
		return bs, err
	}

	for _, ubs := range res {
		bs = append(bs, ent.BuildingMemberQuery{
			BuildingID:   ubs.BuildingID.String(),
			BuildingName: ubs.BuildingName,
			MemberID:     ubs.MemberID.String(),
		})
	}
	return bs, nil
}

func (r *PostgresLocationRepository) GetListMyOwnedBuilding(
	ctx context.Context,
	userID string,
) ([]ent.BuildingQuery, error) {
	res := []ent.BuildingQuery{}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		r.logger.Error(err.Error())
		return res, err
	}
	ownerUUID := uuid.NullUUID{
		UUID:  userUUID,
		Valid: true,
	}

	bs, err := r.Postgres.GetListMyOwnedBuilding(ctx, ownerUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return res, err
	}

	for _, bldng := range bs {
		res = append(res, ent.BuildingQuery{
			ID:   bldng.ID.String(),
			Name: bldng.Name,
		})
	}
	return res, nil
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

func (r *PostgresLocationRepository) GetNumberOfBuildingOwned(
	ctx context.Context,
	ownerID string,
) (int64, error) {
	ownerUUID := uuid.NullUUID{}
	err := ownerUUID.Scan(ownerID)

	if err != nil {
		r.logger.Error(err.Error())
		return -1, err
	}

	count, err := r.Postgres.GetNumberOfBuildingOwned(ctx, ownerUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return -1, err
	}
	return count, nil
}

func (r *PostgresLocationRepository) CreateBuilding(
	ctx context.Context,
	name,
	ownerID string,
) error {
	ownerUUID := uuid.NullUUID{}
	err := ownerUUID.Scan(ownerID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	_, err = r.Postgres.CreateBuilding(ctx, locationpostgres.CreateBuildingParams{
		Name:    name,
		OwnerID: ownerUUID,
	})
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *PostgresLocationRepository) CreateMemberBuilding(
	ctx context.Context,
	memberName,
	userID,
	buildingId string,
) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
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
		UserID: userUUID,
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

func (r *PostgresLocationRepository) EditMemberBuilding(
	ctx context.Context,
	memberID,
	buildingID string,
) error {
	memberUUID, err := uuid.Parse(memberID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	buildingUUID, err := uuid.Parse(buildingID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	_, err = r.Postgres.EditMemberBuildingStatus(ctx, locationpostgres.EditMemberBuildingStatusParams{
		MemberID:   memberUUID,
		BuildingID: buildingUUID,
		StatusID:   2,
	})
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
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
		return []ent.Member{}, err
	}

	roomID := uuid.NullUUID{}
	roomID.Scan(roomUUID)

	res, err := r.Postgres.GetMembersByRoomId(ctx,
		roomID,
	)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.Member{}, err
	}

	ms := []ent.Member{}

	for _, mms := range res {
		ms = append(ms, ent.Member{
			Name:   mms.Name,
			Status: mms.Status,
			UserId: mms.UserID.String(),
		})
	}
	return ms, nil
}

func (r *PostgresLocationRepository) GetRoomsByMemberId(
	ctx context.Context,
	memberId string,
) ([]ent.RoomQuery, error) {
	memberUUID, err := uuid.Parse(memberId)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.RoomQuery{}, err
	}

	res, err := r.Postgres.GetRoomsByMemberId(ctx, memberUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.RoomQuery{}, err
	}

	rs := []ent.RoomQuery{}
	for _, mrs := range res {
		rs = append(rs, ent.RoomQuery{
			Id:   mrs.ID.String(),
			Name: mrs.Name,
		})
	}
	return rs, nil
}

func (r *PostgresLocationRepository) GetRoomsByBuildingId(
	ctx context.Context,
	buildingId string,
) ([]ent.RoomQuery, error) {
	buildingUUID, err := uuid.Parse(buildingId)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.RoomQuery{}, err
	}

	res, err := r.Postgres.GetRoomsByBuildingId(ctx, buildingUUID)
	if err != nil {
		r.logger.Error(err.Error())
		return []ent.RoomQuery{}, err
	}

	rs := []ent.RoomQuery{}
	for _, mrs := range res {
		rs = append(rs, ent.RoomQuery{
			Id:   mrs.ID.String(),
			Name: mrs.Name,
		})
	}
	return rs, nil
}
