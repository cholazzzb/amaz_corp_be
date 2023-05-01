package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/user/mysql"
	user "github.com/cholazzzb/amaz_corp_be/internal/domain/user/repository"
)

func TestCreateUser(t *testing.T) {
	mur := user.NewMockUserRepository()
	mur.CreateUser(context.Background(), mysql.CreateUserParams{
		Username: "new",
		Password: "password",
		Salt:     "salt",
	})

	expected1 := mysql.User{
		ID:       1,
		Username: "new",
		Password: "password",
		Salt:     "salt",
	}

	assert.Equal(t, expected1, mur.Users["new"], "mock data not same with request")

	mur.CreateUser(context.Background(), mysql.CreateUserParams{
		Username: "new2",
		Password: "password2",
		Salt:     "salt2",
	})

	expected2 := mysql.User{
		ID:       2,
		Username: "new2",
		Password: "password2",
		Salt:     "salt2",
	}
	assert.Equal(t, expected2, mur.Users["new2"], "mock data not same with request")
}

func TestGetUser(t *testing.T) {
	mur := user.NewMockUserRepository()
	user, err := mur.GetUser(context.Background(), "not exist")
	assert.Error(t, err, "not exist user not return error")
	assert.Empty(t, user, "not exist user return user")

	mur.CreateUser(context.Background(), mysql.CreateUserParams{
		Username: "test1",
		Password: "password1",
		Salt:     "salt",
	})

	user, err = mur.GetUser(context.Background(), "test1")
	assert.Equal(t, mysql.User{
		ID:       1,
		Username: "test1",
		Password: "password1",
		Salt:     "salt",
	}, user, "exist user different with expected")
	assert.Empty(t, err, "error not empty when success Get User")
}
