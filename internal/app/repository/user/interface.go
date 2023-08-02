package user

import (
	"context"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
)

type UserRepo interface {
	UserRepository
	MemberRepository
	FriendRepository
}

type UserRepository interface {
	GetUser(
		ctx context.Context,
		params string,
	) (user.User, error)
	GetUserExistance(
		ctx context.Context,
		username string,
	) (bool, error)
	CreateUser(
		ctx context.Context,
		params user.User,
	) error
}

type MemberRepository interface {
	GetMemberByName(
		ctx context.Context,
		memberName string,
	) (user.Member, error)
	CreateMember(
		ctx context.Context,
		newMember user.Member,
		userID string,
	) (user.Member, error)
}

type FriendRepository interface {
	GetFriendsByUserId(
		ctx context.Context,
		userId string,
	) ([]user.Member, error)
	CreateFriend(
		ctx context.Context,
		member1Id,
		member2Id string,
	) error
}
