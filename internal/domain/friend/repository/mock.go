package friend

import (
	"context"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
)

type UserId int64

type MockFriendRepository struct {
	BiggestId int64
	Friends   map[UserId]interface{}
}

func NewMockFriendRepository() *MockFriendRepository {
	return &MockFriendRepository{
		BiggestId: 0,
		Friends:   map[UserId]interface{}{},
	}
}

func (mfr *MockFriendRepository) GetFriendsByUserId(ctx context.Context, userId int64) ([]user.Member, error) {
	return nil, nil
}

func (mfr *MockFriendRepository) CreateFriend(ctx context.Context, member1Id, member2Id int64) error {
	return nil
}
