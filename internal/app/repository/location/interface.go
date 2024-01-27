package location

import (
	"context"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
)

type LocationRepo interface {
	BuildingRepository
	MemberRepository
	RoomRepository
	FriendRepository
}

type BuildingRepository interface {
	BuildingRepoCommand
	BuildingRepoQuery
}

type BuildingRepoCommand interface {
	CreateBuilding(
		ctx context.Context,
		name,
		owner_id string,
	) error

	CreateMemberBuilding(
		ctx context.Context,
		memberName,
		userID,
		buildingId string,
	) error

	EditMemberBuilding(
		ctx context.Context,
		memberID,
		buildingID string,
	) error

	DeleteBuilding(
		ctx context.Context,
		buildingId,
		memberId string,
	) error
}

type BuildingRepoQuery interface {
	GetBuildingByID(
		ctx context.Context,
		buildingID string,
	) (ent.BuildingQuery, error)

	GetAllBuildings(
		ctx context.Context,
	) ([]ent.BuildingQuery, error)

	GetMemberBuildingExist(
		ctx context.Context,
		userID,
		buildingID string,
	) (bool, error)

	GetInvitationByUserID(
		ctx context.Context,
		userID string,
	) ([]ent.BuildingMemberQuery, error)

	GetListMyOwnedBuilding(
		ctx context.Context,
		userID string,
	) ([]ent.BuildingQuery, error)

	GetListBuildingByUserID(
		ctx context.Context,
		userID string,
	) ([]ent.BuildingMemberQuery, error)

	GetListMemberByBuildingID(
		ctx context.Context,
		buildingID string,
	) ([]ent.MemberQuery, error)

	GetNumberOfBuildingOwned(
		ctx context.Context,
		ownerID string,
	) (int64, error)
}

type MemberRepository interface {
	MemberRepositoryCommand
	MemberRepositoryQuery
}

type MemberRepositoryCommand interface {
	EditMemberName(
		ctx context.Context,
		memberID,
		memberName string,
	) error
}

type MemberRepositoryQuery interface {
	GetMemberByName(
		ctx context.Context,
		memberName string,
	) (ent.MemberQuery, error)

	GetMemberByID(
		ctx context.Context,
		memberID string,
	) (ent.MemberQuery, error)
}

type RoomRepository interface {
	GetMembersByRoomId(
		ctx context.Context,
		roomId string,
	) ([]ent.Member, error)
	GetRoomsByMemberId(
		ctx context.Context,
		memberId string,
	) ([]ent.RoomQuery, error)
	GetRoomsByBuildingId(
		ctx context.Context,
		buildingId string,
	) ([]ent.RoomQuery, error)
}

type FriendRepository interface {
	GetFriendsByUserId(
		ctx context.Context,
		userId string,
	) ([]ent.MemberQuery, error)
	CreateFriend(
		ctx context.Context,
		member1Id,
		member2Id string,
	) error
}
