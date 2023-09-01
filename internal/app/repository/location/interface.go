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
	BuildingRepoCommand
	BuildingRepoQuery
}

type BuildingRepoCommand interface {
	CreateMemberBuilding(
		ctx context.Context,
		memberId,
		buildingId string,
	) error

	DeleteBuilding(
		ctx context.Context,
		buildingId,
		memberId string,
	) error
}

type BuildingRepoQuery interface {
	GetAllBuildings(
		ctx context.Context,
	) ([]ent.Building, error)

	GetBuildingsByMemberId(
		ctx context.Context,
		memberId string,
	) ([]ent.Building, error)

	GetListMemberByBuildingID(
		ctx context.Context,
		buildingID string,
	) ([]ent.MemberQuery, error)
}

type RoomRepository interface {
	GetMembersByRoomId(
		ctx context.Context,
		roomId string,
	) ([]ent.Member, error)
	GetRoomsByMemberId(
		ctx context.Context,
		memberId string,
	) ([]ent.Room, error)
	GetRoomsByBuildingId(
		ctx context.Context,
		buildingId string,
	) ([]ent.Room, error)
}
