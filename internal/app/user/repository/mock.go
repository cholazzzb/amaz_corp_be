package user

import (
	"context"
	"fmt"

	ent "github.com/cholazzzb/amaz_corp_be/internal/app/user"
	mysql "github.com/cholazzzb/amaz_corp_be/internal/app/user/mysql"
)

type MockUserRepo struct {
	User   *MockUserRepository
	Member *MockMemberRepository
	Friend *MockFriendRepository
}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{
		User:   newMockUserRepository(),
		Member: newMockMemberRepository(),
		Friend: newMockFriendRepository(),
	}
}

type Username string

type MockUserRepository struct {
	BiggestId int64
	Users     map[Username]mysql.User
}

func newMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		BiggestId: 0,
		Users:     map[Username]mysql.User{},
	}
}

func (mur *MockUserRepo) GetUser(
	ctx context.Context,
	params string,
) (mysql.User, error) {
	user, ok := mur.User.Users[Username(params)]
	if !ok {
		return mysql.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (mur *MockUserRepo) CreateUser(
	ctx context.Context,
	params mysql.CreateUserParams,
) error {
	id := mur.User.BiggestId + 1
	newUser := mysql.User{
		ID:       id,
		Username: params.Username,
		Password: params.Password,
		Salt:     params.Salt,
	}

	mur.User.BiggestId = id
	mur.User.Users[Username(params.Username)] = newUser
	return nil
}

type Name string

type MockMemberRepository struct {
	BiggestId int64
	Members   map[Name]mysql.Member
}

func newMockMemberRepository() *MockMemberRepository {
	return &MockMemberRepository{
		BiggestId: 0,
		Members:   map[Name]mysql.Member{},
	}
}

func (mmr *MockUserRepo) GetMemberByName(
	ctx context.Context,
	memberName string,
) (ent.Member, error) {
	m, ok := mmr.Member.Members[Name(memberName)]
	if !ok {
		return ent.Member{}, fmt.Errorf("member not found")
	}
	return ent.Member{
		Name:   m.Name,
		Status: m.Status,
	}, nil
}

func (mmr *MockUserRepo) CreateMember(
	ctx context.Context,
	newMember ent.Member,
	userID int64,
) (ent.Member, error) {
	ID := mmr.Member.BiggestId + 1

	mmr.Member.BiggestId = ID
	mmr.Member.Members[Name(newMember.Name)] = mysql.Member{
		ID:     ID,
		Name:   newMember.Name,
		Status: newMember.Status,
		UserID: userID,
	}
	return newMember, nil
}

type UserId int64

type MockFriendRepository struct {
	BiggestId int64
	Friends   map[UserId]interface{}
}

func newMockFriendRepository() *MockFriendRepository {
	return &MockFriendRepository{
		BiggestId: 0,
		Friends:   map[UserId]interface{}{},
	}
}

func (mur *MockUserRepo) GetFriendsByUserId(
	ctx context.Context,
	userId int64,
) ([]ent.Member, error) {
	return nil, nil
}

func (mur *MockUserRepo) CreateFriend(
	ctx context.Context,
	member1Id,
	member2Id int64,
) error {
	return nil
}
