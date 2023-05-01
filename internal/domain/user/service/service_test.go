package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cholazzzb/amaz_corp_be/internal/app/repository"
	user "github.com/cholazzzb/amaz_corp_be/internal/domain/user/service"
)

func CreateMockService() *user.UserService {
	repo := repository.CreateMockRepository()
	return user.NewUserService(repo)
}

func TestRegisterUser(t *testing.T) {
	s := CreateMockService()
	err := s.RegisterUser(context.Background(), "username", "password")

	assert.Empty(t, err, "failed to register user")
}

func TestLoginUser(t *testing.T) {
	s := CreateMockService()
	t1, err := s.Login(context.Background(), "not exist", "not exist")
	assert.Empty(t, t1, "not exist user when login return not empty token")
	assert.Error(t, err, "not exist user when login return empty error")

	err = s.RegisterUser(context.Background(), "user1", "password1")
	assert.Empty(t, err)
	t2, err := s.Login(context.Background(), "user1", "password1")
	assert.Empty(t, err, "successful login return error")
	assert.NotEmpty(t, t2, "successful login return empty token")
}
