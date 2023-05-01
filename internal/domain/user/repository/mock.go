package user

import (
	"context"
	"fmt"

	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/user/mysql"
)

type Username string

type MockUserRepository struct {
	BiggestId int64
	Users     map[Username]mysql.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		BiggestId: 0,
		Users:     map[Username]mysql.User{},
	}
}

func (mur *MockUserRepository) GetUser(ctx context.Context, params string) (mysql.User, error) {
	user, ok := mur.Users[Username(params)]
	if !ok {
		return mysql.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (mur *MockUserRepository) CreateUser(ctx context.Context, params mysql.CreateUserParams) error {
	id := mur.BiggestId + 1
	newUser := mysql.User{
		ID:       id,
		Username: params.Username,
		Password: params.Password,
		Salt:     params.Salt,
	}

	mur.BiggestId = id
	mur.Users[Username(params.Username)] = newUser
	return nil
}
