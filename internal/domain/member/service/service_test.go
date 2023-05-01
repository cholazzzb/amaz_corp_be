package member_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cholazzzb/amaz_corp_be/internal/app/repository"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/member"
	memberSvc "github.com/cholazzzb/amaz_corp_be/internal/domain/member/service"
	user "github.com/cholazzzb/amaz_corp_be/internal/domain/user/service"
)

func CreateMockService() (*user.UserService, *memberSvc.MemberService) {
	repo := repository.CreateMockRepository()
	u := user.NewUserService(repo)
	m := memberSvc.NewMemberService(repo)
	return u, m
}

func TestGetMemberByName(t *testing.T) {
	us, ms := CreateMockService()
	m, err := ms.GetMemberByName(context.Background(), "member1")
	assert.Empty(t, m, "not exist member should return empty object")
	assert.Error(t, err, "not exist member should return error")

	err = us.RegisterUser(context.Background(), "username1", "password1")
	assert.NoError(t, err, "register user service with right params should success")
	m2, err := ms.CreateMember(context.Background(), member.Member{
		Name:   "name1",
		Status: "status1",
	}, "username1")
	assert.Empty(t, err, "create member with true params should not return error")
	assert.Equal(t, member.Member{
		Name:   "name1",
		Status: "status1",
	}, m2, "create member with true params should return the true member data")

	m3, err := ms.GetMemberByName(context.Background(), "name1")
	assert.Equal(t, member.Member{
		Name:   "name1",
		Status: "status1",
	}, m3, "exist member should return member object")
	assert.Empty(t, err, "exist member should not return error")
}

func TestCreateMember(t *testing.T) {
	us, ms := CreateMockService()

	err := us.RegisterUser(context.Background(), "username1", "password1")
	assert.NoError(t, err, "register user service with right params should success")
	m1, err := ms.CreateMember(context.Background(), member.Member{
		Name:   "name1",
		Status: "status1",
	}, "username1")
	assert.Empty(t, err, "create member with true params should not return error")
	assert.Equal(t, member.Member{
		Name:   "name1",
		Status: "status1",
	}, m1, "create member with true params should return the true member data")

}
