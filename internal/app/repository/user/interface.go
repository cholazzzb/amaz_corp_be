package user

import (
	"context"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
)

type UserRepo interface {
	UserRepository
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
