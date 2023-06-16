package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/heartbeat"
)

type HeartbeatService struct {
	repo   repo.HeartbeatRepo
	logger zerolog.Logger
}

func NewHeartbeatService(
	repo repo.HeartbeatRepo,
) *HeartbeatService {
	sublogger := log.With().Str("layer", "service").Str("package", "heartbeat").Logger()

	return &HeartbeatService{
		repo:   repo,
		logger: sublogger,
	}
}

func (s *HeartbeatService) Pulse(ctx context.Context, userId int64) error {
	err := s.repo.UpdateToOnline(ctx, userId)

	if err != nil {
		errMessage := fmt.Errorf("failed to update to online in the heartbeat map")
		s.logger.Error().Err(errMessage)
		return errMessage
	}

	return nil
}

func (s *HeartbeatService) CheckUserStatus(ctx context.Context, userId int64) (string, error) {
	exist, err := s.repo.CheckUserIdExistence(ctx, userId)

	if err != nil {
		errMessage := fmt.Errorf("failed to checkUserIdExistence: %v", userId)
		s.logger.Error().Err(errMessage)
		return "", errMessage
	}

	if exist {
		return "online", nil
	} else {
		return "offline", nil
	}
}
