package friend

import (
	"context"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/member"
)

type FriendRepository interface {
	GetFriendsByUserId(ctx context.Context, userId int64) ([]member.Member, error)
	CreateFriend(ctx context.Context, member1Id, member2Id int64) error
}
