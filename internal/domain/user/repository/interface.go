package user

import (
	"context"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/user/mysql"
)

type UserRepo interface {
	UserRepository
	MemberRepository
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
