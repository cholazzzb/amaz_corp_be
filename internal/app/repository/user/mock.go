package user

import (
	"context"
	"fmt"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
)

type MockUserRepo struct {
	User *MockUserRepository
}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{
		User: newMockUserRepository(),
	}
}

type Username string

type MockUserRepository struct {
	BiggestId int64
	Users     map[Username]user.User
}

func newMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		BiggestId: 0,
		Users:     map[Username]user.User{},
	}
}

func (mur *MockUserRepo) GetUser(
	ctx context.Context,
	params string,
) (user.User, error) {
	res, ok := mur.User.Users[Username(params)]
	if !ok {
		return user.User{}, fmt.Errorf("user not found")
	}
	return res, nil
}

func (mur *MockUserRepo) GetUserExistance(
	ctx context.Context,
	username string,
) (bool, error) {
	_, exist := mur.User.Users[Username(username)]
	return exist, nil
}

func (mur *MockUserRepo) CreateUser(
	ctx context.Context,
	params user.UserCommand,
) error {
	id := mur.User.BiggestId + 1
	newUser := user.User{
		Username: params.Username,
		Password: params.Password,
		Salt:     params.Salt,
	}

	mur.User.BiggestId = id
	mur.User.Users[Username(params.Username)] = newUser
	return nil
}

// TODO: Implement this
func (mur *MockUserRepo) GetProductByUserID(
	ctx context.Context,
	userID string,
) (user.ProductQuery, error) {
	return user.ProductQuery{}, nil
}

func (mur *MockUserRepo) GetListProduct(
	ctx context.Context,
) ([]user.ProductQuery, error) {
	return []user.ProductQuery{}, nil
}

func (mur *MockUserRepo) GetListFeatureByProductID(
	ctx context.Context,
	productID int32,
) ([]user.FeatureQuery, error) {
	return []user.FeatureQuery{}, nil
}
