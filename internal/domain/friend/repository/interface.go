package friend

import (
	"context"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
)

type FriendRepository interface {
	GetFriendsByUserId(ctx context.Context, userId int64) ([]user.Member, error)
	CreateFriend(ctx context.Context, member1Id, member2Id int64) error
}
