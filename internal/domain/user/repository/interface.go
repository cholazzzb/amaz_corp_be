package user

import (
	"context"

	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/user/mysql"
)

type UserRepository interface {
	GetUser(ctx context.Context, params string) (mysql.User, error)
	CreateUser(ctx context.Context, params mysql.CreateUserParams) error
}
