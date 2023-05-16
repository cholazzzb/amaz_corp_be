package user

import (
	"context"

	"github.com/cholazzzb/amaz_corp_be/internal/app/user"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/app/user/mysql"
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
	) (mysql.User, error)
	CreateUser(
		ctx context.Context,
		params mysql.CreateUserParams,
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
		userID int64,
	) (user.Member, error)
}

type FriendRepository interface {
	GetFriendsByUserId(
		ctx context.Context,
		userId int64,
	) ([]user.Member, error)
	CreateFriend(
		ctx context.Context,
		member1Id,
		member2Id int64,
	) error
}
