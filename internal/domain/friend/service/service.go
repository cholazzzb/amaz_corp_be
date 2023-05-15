package friend

import (
	"context"
	"fmt"

	"github.com/cholazzzb/amaz_corp_be/internal/app/repository"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type FriendService struct {
	repo   *repository.Repository
	logger zerolog.Logger
}

func NewFriendService(
	repo *repository.Repository,
) *FriendService {
	sublogger := log.With().Str("layer", "service").Str("package", "member").Logger()

	return &FriendService{
		repo:   repo,
		logger: sublogger,
	}
}

func (svc *FriendService) GetFriendsByMemberId(ctx context.Context, userId int64) ([]user.Member, error) {
	fs, err := svc.repo.Friend.GetFriendsByUserId(ctx, userId)
	if err != nil {
		svc.logger.Error().Err(err)
		return nil, fmt.Errorf("cannot find friends with name %s", fs)
	}
	return fs, nil
}
