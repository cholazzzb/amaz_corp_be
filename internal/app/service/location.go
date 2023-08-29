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
) ([]ent.Building, error) {
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

func (svc *LocationService) GetBuildingsByMemberId(
	ctx context.Context,
	memberId string,
) ([]ent.Building, error) {
	bs, err := svc.repo.GetBuildingsByMemberId(ctx, memberId)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, fmt.Errorf("cannot get buildings with memberId %s", memberId)
	}
	return bs, nil
}

func (svc *LocationService) JoinBuilding(
	ctx context.Context,
	memberId,
	buildingId string,
) error {
	err := svc.repo.CreateMemberBuilding(ctx, memberId, buildingId)
	if err != nil {
		svc.logger.Error(err.Error())
		return fmt.Errorf("cannot join member with id %s to building id %s", memberId, buildingId)
	}
	return nil
}

func (svc *LocationService) GetRoomsByBuildingId(
	ctx context.Context,
	buildingId string,
) ([]ent.Room, error) {
	bs, err := svc.repo.GetRoomsByBuildingId(ctx, buildingId)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, fmt.Errorf("cannot get rooms with buildingID %s", buildingId)
	}
	return bs, nil
}
