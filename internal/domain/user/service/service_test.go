package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/user"
	repo "github.com/cholazzzb/amaz_corp_be/internal/domain/user/repository"
	svc "github.com/cholazzzb/amaz_corp_be/internal/domain/user/service"
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
	m2, err := s.CreateMember(context.Background(), ent.Member{
		Name:   "name1",
		Status: "status1",
	}, "username1")
	assert.Empty(t, err, "create member with true params should not return error")
	assert.Equal(t, ent.Member{
		Name:   "name1",
		Status: "status1",
	}, m2, "create member with true params should return the true member data")

	m3, err := s.GetMemberByName(context.Background(), "name1")
	assert.Equal(t, ent.Member{
		Name:   "name1",
		Status: "status1",
	}, m3, "exist member should return member object")
	assert.Empty(t, err, "exist member should not return error")
}

func TestCreateMember(t *testing.T) {
	s := CreateMockService()

	err := s.RegisterUser(context.Background(), "username1", "password1")
	assert.NoError(t, err, "register user service with right params should success")
	m1, err := s.CreateMember(context.Background(), ent.Member{
		Name:   "name1",
		Status: "status1",
	}, "username1")
	assert.Empty(t, err, "create member with true params should not return error")
	assert.Equal(t, ent.Member{
		Name:   "name1",
		Status: "status1",
	}, m1, "create member with true params should return the true member data")
}
