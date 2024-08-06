package user

import (
	"context"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
)

type UserRepo interface {
	UserRepository
	ProductRepoQuery
}

type UserRepository interface {
	GetUser(
		ctx context.Context,
		params string,
	) (user.User, error)
	GetListUserByUsername(
		ctx context.Context,
		username string,
	) ([]user.UserQuery, error)
	GetUserExistance(
		ctx context.Context,
		username string,
	) (bool, error)
	CreateUser(
		ctx context.Context,
		params user.UserCommand,
	) (string, error)
}

type ProductRepoQuery interface {
	GetProductByUserID(
		ctx context.Context,
		userID string,
	) (user.ProductQuery, error)
	GetListProduct(
		ctx context.Context,
	) ([]user.ProductQuery, error)
	GetListFeatureByProductID(
		ctx context.Context,
		productID int32,
	) ([]user.FeatureQuery, error)
}
