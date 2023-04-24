package member

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/app/repository"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/member"
)

type MemberService struct {
	repo   *repository.Repository
	logger zerolog.Logger
}

func NewMemberService(
	repo *repository.Repository,
) *MemberService {
	sublogger := log.With().Str("layer", "service").Str("package", "member").Logger()

	return &MemberService{
		repo:   repo,
		logger: sublogger,
	}
}

func (svc *MemberService) GetMemberByName(ctx context.Context, name string) (member.Member, error) {
	member, err := svc.repo.Member.GetMemberByName(ctx, name)
	if err != nil {
		return member, fmt.Errorf("cannot find member with name %s", name)
	}
	return member, nil
}

func (svc *MemberService) CreateMember(ctx context.Context, memberReq member.Member, username string) (member.Member, error) {
	userData, err := svc.repo.User.GetUser(ctx, username)
	if err != nil {
		svc.logger.Error().Err(err)
		return member.Member{}, fmt.Errorf("cannot found user with username %s", username)
	}
	newMember, err := svc.repo.Member.CreateMember(ctx, memberReq, userData.ID)
	if err != nil {
		svc.logger.Error().Err(err)
		return member.Member{}, fmt.Errorf("failed to create member %v", memberReq)
	}
	return newMember, nil
}
