package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/user"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/domain/user/mysql"
	user "github.com/cholazzzb/amaz_corp_be/internal/domain/user/repository"
)

func TestCreateUser(t *testing.T) {
	mur := user.NewMockUserRepo()
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

	assert.Equal(t, expected1, mur.User.Users["new"], "mock data not same with request")

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
	assert.Equal(t, expected2, mur.User.Users["new2"], "mock data not same with request")
}

func TestGetUser(t *testing.T) {
	mur := user.NewMockUserRepo()
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

func TestCreateMember(t *testing.T) {
	mmr := user.NewMockUserRepo()

	mmr.CreateMember(context.Background(), ent.Member{
		Name:   "test1",
		Status: "new",
	}, 3)

	expected1 := mysql.Member{
		ID:     1,
		Name:   "test1",
		Status: "new",
		UserID: 3,
	}
	assert.Equal(t, expected1, mmr.Member.Members["test1"], "mock member data not same with request")

	mmr.CreateMember(context.Background(), ent.Member{
		Name:   "test2",
		Status: "new",
	}, 2)

	expected2 := mysql.Member{
		ID:     2,
		Name:   "test2",
		Status: "new",
		UserID: 2,
	}

	assert.Equal(t, expected2, mmr.Member.Members["test2"], "second member data not same with request")
}

func TestGetMemberByName(t *testing.T) {
	mmr := user.NewMockUserRepo()
	m1, err := mmr.GetMemberByName(context.Background(), "not exist")
	assert.Error(t, err, "not exist member not return error")
	assert.Empty(t, m1, "not exist user return user")

	mmr.CreateMember(context.Background(), ent.Member{
		Name:   "name1",
		Status: "status1",
	}, 2)

	m2, err := mmr.GetMemberByName(context.Background(), "name1")
	assert.Empty(t, err, "exist member return error")

	e1 := ent.Member{
		Name:   "name1",
		Status: "status1",
	}
	assert.Equal(t, e1, m2, "get member return different member")
}
