package service

import (
	"context"
	"fmt"
	"log/slog"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/heartbeat"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type HeartbeatService struct {
	repo   repo.HeartbeatRepo
	logger *slog.Logger
}

func NewHeartbeatService(
	repo repo.HeartbeatRepo,
) *HeartbeatService {
	sublogger := logger.Get().With(slog.String("domain", "heartbeat"), slog.String("layer", "svc"))

	return &HeartbeatService{
		repo:   repo,
		logger: sublogger,
	}
}

func (s *HeartbeatService) Pulse(ctx context.Context, userId string) error {
	err := s.repo.UpdateToOnline(ctx, userId)

	if err != nil {
		errMessage := fmt.Errorf("failed to update to online in the heartbeat map")
		s.logger.Error(err.Error())
		return errMessage
	}

	return nil
}

func (s *HeartbeatService) CheckUserStatus(ctx context.Context, userId string) (string, error) {
	exist, err := s.repo.CheckUserIdExistence(ctx, userId)

	if err != nil {
		errMessage := fmt.Errorf("failed to checkUserIdExistence: %v", userId)
		s.logger.Error(err.Error())
		return "", errMessage
	}

	if exist {
		return "online", nil
	} else {
		return "offline", nil
	}
}
