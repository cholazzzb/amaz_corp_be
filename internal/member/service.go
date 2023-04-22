package member

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type MemberService struct {
	repo   *MySQLMemberRepository
	logger zerolog.Logger
}

func NewMemberService(
	repo *MySQLMemberRepository,
) *MemberService {
	sublogger := log.With().Str("layer", "service").Str("package", "member").Logger()

	return &MemberService{
		repo:   repo,
		logger: sublogger,
	}
}

func (svc *MemberService) GetMemberByName(ctx context.Context, name string) (Member, error) {
	return Member{}, nil
}

func (svc *MemberService) CreateMember(ctx context.Context, member Member) (Member, error) {
	return Member{}, nil
}
