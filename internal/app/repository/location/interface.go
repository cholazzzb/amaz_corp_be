package location

import (
	"context"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
)

type LocationRepo interface {
	BuildingRepository
	RoomRepository
}

type BuildingRepository interface {
	GetAllBuildings(
		ctx context.Context,
	) ([]ent.Building, error)

	GetBuildingsByMemberId(
		ctx context.Context,
		memberId int64,
	) ([]ent.Building, error)

	CreateMemberBuilding(
		ctx context.Context,
		memberId int64,
		buildingId int64,
	) error
}

type RoomRepository interface {
	GetMembersByRoomId(
		ctx context.Context,
		roomId int64,
	) ([]ent.Member, error)
	GetRoomsByMemberId(
		ctx context.Context,
		memberId int64,
	) ([]ent.Room, error)
	GetRoomsByBuildingId(
		ctx context.Context,
		buildingId int64,
	) ([]ent.Room, error)
}
