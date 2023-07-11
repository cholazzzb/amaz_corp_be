package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/location"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
)

type LocationService struct {
	hrs    *HeartbeatService
	us     *UserService
	repo   repo.LocationRepo
	logger zerolog.Logger
}

func NewLocationService(
	hrs *HeartbeatService,
	us *UserService,
	repo repo.LocationRepo,
) *LocationService {
	sublogger := log.With().Str("layer", "service").Str("package", "location").Logger()

	return &LocationService{
		hrs:    hrs,
		us:     us,
		repo:   repo,
		logger: sublogger,
	}
}

func (svc *LocationService) GetListOnlineMembers(
	ctx context.Context,
	roomId int64,
) ([]ent.Member, error) {
	membersInRoom, err := svc.repo.GetMembersByRoomId(ctx, roomId)
	if err != nil {
		svc.logger.Error().Err(err)
		return []ent.Member{}, errors.New("failed to get rooms")
	}

	onlineMap, err := svc.hrs.repo.GetHeartbeatMap(ctx)
	if err != nil {
		svc.logger.Error().Err(err)
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
		svc.logger.Error().Err(err)
		return nil, fmt.Errorf("cannot get all buildings")
	}
	return bs, nil
}

func (svc *LocationService) DeleteBuilding(
	ctx context.Context,
	buildingId,
	memberId int64,
) error {
	err := svc.repo.DeleteBuilding(ctx, buildingId, memberId)
	if err != nil {
		svc.logger.Error().Err(err)
		return fmt.Errorf("cannot delete with id %d", buildingId)
	}
	return nil
}

func (svc *LocationService) GetBuildingsByMemberId(
	ctx context.Context,
	memberId int64,
) ([]ent.Building, error) {
	bs, err := svc.repo.GetBuildingsByMemberId(ctx, memberId)
	if err != nil {
		svc.logger.Error().Err(err)
		return nil, fmt.Errorf("cannot get buildings with memberId %d", memberId)
	}
	return bs, nil
}

func (svc *LocationService) JoinBuilding(
	ctx context.Context,
	memberId int64,
	buildingId int64,
) error {
	err := svc.repo.CreateMemberBuilding(ctx, memberId, buildingId)
	if err != nil {
		svc.logger.Error().Err(err)
		return fmt.Errorf("cannot join member with id %d to building id %d", memberId, buildingId)
	}
	return nil
}

func (svc *LocationService) GetRoomsByBuildingId(
	ctx context.Context,
	buildingId int64,
) ([]ent.Room, error) {
	bs, err := svc.repo.GetRoomsByBuildingId(ctx, buildingId)
	if err != nil {
		svc.logger.Error().Err(err)
		return nil, fmt.Errorf("cannot get rooms with buildingID %d", buildingId)
	}
	return bs, nil
}
