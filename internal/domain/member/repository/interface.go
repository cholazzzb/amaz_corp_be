package member

import (
	"context"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/member"
)

type MemberRepository interface {
	GetMemberByName(ctx context.Context, memberName string) (member.Member, error)
	CreateMember(ctx context.Context, newMember member.Member, userID int64) (member.Member, error)
}
