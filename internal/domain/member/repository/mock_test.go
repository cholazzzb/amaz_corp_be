package member_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	member "github.com/cholazzzb/amaz_corp_be/internal/domain/member"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/member/mysql"
	memberRepository "github.com/cholazzzb/amaz_corp_be/internal/domain/member/repository"
)

func TestCreateMember(t *testing.T) {
	mmr := memberRepository.NewMockMemberRepository()

	mmr.CreateMember(context.Background(), member.Member{
		Name:   "test1",
		Status: "new",
	}, 3)

	expected1 := mysql.Member{
		ID:     1,
		Name:   "test1",
		Status: "new",
		UserID: 3,
	}
	assert.Equal(t, expected1, mmr.Members["test1"], "mock member data not same with request")

	mmr.CreateMember(context.Background(), member.Member{
		Name:   "test2",
		Status: "new",
	}, 2)

	expected2 := mysql.Member{
		ID:     2,
		Name:   "test2",
		Status: "new",
		UserID: 2,
	}

	assert.Equal(t, expected2, mmr.Members["test2"], "second member data not same with request")
}

func TestGetMemberByName(t *testing.T) {
	mmr := memberRepository.NewMockMemberRepository()
	m1, err := mmr.GetMemberByName(context.Background(), "not exist")
	assert.Error(t, err, "not exist member not return error")
	assert.Empty(t, m1, "not exist user return user")

	mmr.CreateMember(context.Background(), member.Member{
		Name:   "name1",
		Status: "status1",
	}, 2)

	m2, err := mmr.GetMemberByName(context.Background(), "name1")
	assert.Empty(t, err, "exist member return error")

	e1 := member.Member{
		Name:   "name1",
		Status: "status1",
	}
	assert.Equal(t, e1, m2, "get member return different member")
}
