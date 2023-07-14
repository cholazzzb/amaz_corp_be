package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/user"
	svc "github.com/cholazzzb/amaz_corp_be/internal/app/service"
)

func CreateMockService() *svc.UserService {
	repo := repo.NewMockUserRepo()
	return svc.NewUserService(repo)
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

func TestGetMemberByName(t *testing.T) {
	s := CreateMockService()
	m, err := s.GetMemberByName(context.Background(), "member1")
	assert.Empty(t, m, "not exist member should return empty object")
	assert.Error(t, err, "not exist member should return error")

	err = s.RegisterUser(context.Background(), "username1", "password1")
	assert.NoError(t, err, "register user service with right params should success")
	m2, err := s.CreateMember(context.Background(), "name1", "username1")
	assert.Empty(t, err, "create member with true params should not return error")
	assert.Equal(t, "name1", m2.Name, "create member with true params should return the true member name")
	assert.Equal(t, "new member", m2.Status, "create member with true params should return the true member status")

	m3, err := s.GetMemberByName(context.Background(), "name1")
	assert.Equal(t, "name1", m3.Name, "exist member should return the true member name")
	assert.Equal(t, "new member", m3.Status, "exist member with true params should return the true member name")

	assert.Empty(t, err, "exist member should not return error")
}

func TestCreateMember(t *testing.T) {
	s := CreateMockService()

	err := s.RegisterUser(context.Background(), "username1", "password1")
	assert.NoError(t, err, "register user service with right params should success")
	m1, err := s.CreateMember(context.Background(), "name1", "username1")
	assert.Empty(t, err, "create member with true params should not return error")
	assert.Equal(t, "name1", m1.Name, "create member with true params should return the true member name")
	assert.Equal(t, "new member", m1.Status, "create member with true params should return the true member status")
}
