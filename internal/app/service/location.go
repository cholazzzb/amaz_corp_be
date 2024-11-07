package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/location"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type LocationService struct {
	hrs    *HeartbeatService
	us     *UserService
	repo   repo.LocationRepo
	logger *slog.Logger
}

func NewLocationService(
	hrs *HeartbeatService,
	us *UserService,
	repo repo.LocationRepo,
) *LocationService {
	sublogger := logger.Get().With(slog.String("domain", "location"), slog.String("layer", "svc"))

	return &LocationService{
		hrs:    hrs,
		us:     us,
		repo:   repo,
		logger: sublogger,
	}
}

func (svc *LocationService) GetListOnlineMembers(
	ctx context.Context,
	roomId string,
) ([]ent.Member, error) {
	membersInRoom, err := svc.repo.GetMembersByRoomId(ctx, roomId)
	if err != nil {
		svc.logger.Error(err.Error())
		return []ent.Member{}, errors.New("failed to get rooms")
	}

	onlineMap, err := svc.hrs.repo.GetHeartbeatMap(ctx)
	if err != nil {
		svc.logger.Error(err.Error())
		return []ent.Member{}, errors.New("failed to get heartBeatMap")
	}

	ms := []ent.Member{}
	for _, mir := range membersInRoom {
		if _, ok := onlineMap[mir.UserId]; ok {
			ms = append(ms, mir)
		}
	}

	return ms, nil
}

func (svc *LocationService) GetBuildings(
	ctx context.Context,
) ([]ent.BuildingQuery, error) {
	bs, err := svc.repo.GetAllBuildings(ctx)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, fmt.Errorf("cannot get all buildings")
	}
	return bs, nil
}

func (svc *LocationService) DeleteBuilding(
	ctx context.Context,
	buildingId,
	memberId string,
) error {
	err := svc.repo.DeleteBuilding(ctx, buildingId, memberId)
	if err != nil {
		svc.logger.Error(err.Error())
		return fmt.Errorf("cannot delete with id %s", buildingId)
	}
	return nil
}

func (svc *LocationService) GetBuildingByID(
	ctx context.Context,
	buildingID string,
) (ent.BuildingQuery, error) {
	building, err := svc.repo.GetBuildingByID(ctx, buildingID)
	if err != nil {
		return ent.BuildingQuery{}, err
	}
	return building, nil
}

func (svc *LocationService) GetBuildingsByUserID(
	ctx context.Context,
	userID string,
) ([]ent.BuildingMemberQuery, error) {
	bs, err := svc.repo.GetListBuildingByUserID(ctx, userID)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, fmt.Errorf("cannot get buildings with userID %s", userID)
	}
	return bs, nil
}

func (svc *LocationService) GetMyInvitation(
	ctx context.Context,
	userID string,
) ([]ent.BuildingMemberQuery, error) {
	bldngs, err := svc.repo.GetInvitationByUserID(ctx, userID)
	if err != nil {
		return []ent.BuildingMemberQuery{}, err
	}
	return bldngs, err
}

func (svc *LocationService) GetListMyOwnedBuilding(
	ctx context.Context,
	userID string,
) ([]ent.BuildingQuery, error) {
	bldngs, err := svc.repo.GetListMyOwnedBuilding(ctx, userID)
	if err != nil {
		return []ent.BuildingQuery{}, err
	}

	return bldngs, nil
}

func (svc *LocationService) GetListMemberByBuildingID(
	ctx context.Context,
	buildingID string,
) ([]ent.MemberQuery, error) {
	ms, err := svc.repo.GetListMemberByBuildingID(ctx, buildingID)
	if err != nil {
		svc.logger.Error(err.Error())
		return []ent.MemberQuery{}, nil
	}

	return ms, nil
}

func (svc *LocationService) GetNumberOfBuildingOwned(
	ctx context.Context,
	ownerID string,
) (int64, error) {
	count, err := svc.repo.GetNumberOfBuildingOwned(ctx, ownerID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (svc *LocationService) CheckMemberBuildingExist(
	ctx context.Context,
	userID,
	buildingID string,
) (bool, error) {
	return svc.repo.GetMemberBuildingExist(ctx, userID, buildingID)
}

func (svc *LocationService) CreateBuilding(
	ctx context.Context,
	name,
	userID string,
) error {
	product, err := svc.us.GetProductByUserID(ctx, userID)
	if err != nil {
		svc.logger.Error(err.Error())
		return err
	}

	owned, err := svc.repo.GetNumberOfBuildingOwned(ctx, userID)
	if err != nil {
		return err
	}

	memberBuilding, _ := svc.repo.GetListBuildingByUserID(ctx, userID)

	if product.ID == 1 && (owned >= 1 || len(memberBuilding) > 0) {
		err = errors.New("free user already reach limit create building")
		svc.logger.Error(err.Error())
		return err
	}

	err = svc.repo.CreateBuilding(ctx, name, userID)
	if err != nil {
		return err
	}

	return nil
}

func (svc *LocationService) InviteMemberToBuilding(
	ctx context.Context,
	memberName,
	userID,
	buildingId string,
) error {
	err := svc.repo.CreateMemberBuilding(ctx, memberName, userID, buildingId)
	if err != nil {
		return fmt.Errorf("cannot join member with userID %s to building id %s", userID, buildingId)
	}
	return nil
}

func (svc *LocationService) JoinBuilding(
	ctx context.Context,
	memberID,
	buildingID string,
) error {
	err := svc.repo.EditMemberBuilding(ctx, memberID, buildingID)
	if err != nil {
		return fmt.Errorf("join building err:%w", err)
	}
	return nil
}

func (svc *LocationService) EditMemberName(
	ctx context.Context,
	memberID,
	name string,
) error {
	err := svc.repo.EditMemberName(ctx, memberID, name)
	if err != nil {
		return err
	}
	return nil
}

func (svc *LocationService) GetMemberByName(ctx context.Context, name string) (ent.MemberQuery, error) {
	member, err := svc.repo.GetMemberByName(ctx, name)
	if err != nil {
		return member, err
	}
	return member, nil
}

func (svc *LocationService) GetMemberByID(
	ctx context.Context,
	memberID string,
) (ent.MemberQuery, error) {
	member, err := svc.repo.GetMemberByID(ctx, memberID)
	if err != nil {
		return member, err
	}
	return member, nil
}

func (svc *LocationService) GetFriendsByMemberId(ctx context.Context, userId string) ([]ent.MemberQuery, error) {
	fs, err := svc.repo.GetFriendsByUserId(ctx, userId)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, fmt.Errorf("cannot find friends with name %s", fs)
	}
	return fs, nil
}

func (svc *LocationService) GetRoomsByBuildingId(
	ctx context.Context,
	buildingId string,
) ([]ent.RoomQuery, error) {
	bs, err := svc.repo.GetRoomsByBuildingId(ctx, buildingId)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, fmt.Errorf("cannot get rooms with buildingID %s", buildingId)
	}
	return bs, nil
}

func (svc *LocationService) CreateRoom(
	ctx context.Context,
	name, buildingID, userID string,
) error {
	myBuildings, _ := svc.repo.GetListMyOwnedBuilding(ctx, userID)

	owned := false
	for _, bld := range myBuildings {
		if bld.ID == buildingID {
			owned = true
		}
	}

	if !owned {
		return errors.New("cannot create room for the building not owned")
	}

	err := svc.repo.CreateRoom(ctx, name, buildingID)
	if err != nil {
		return err
	}
	return nil
}
