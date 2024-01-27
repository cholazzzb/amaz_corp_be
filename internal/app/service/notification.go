package service

import (
	"log/slog"

	repo "github.com/cholazzzb/amaz_corp_be/internal/app/repository/notification"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type NotificationService struct {
	repo   repo.NotificationRepo
	logger *slog.Logger
}

func NewNotificationService(
	repo repo.NotificationRepo,
) *NotificationService {
	sublogger := logger.Get().With(
		slog.String("domain", "notification"),
		slog.String("layer", "svc"),
	)

	return &NotificationService{
		repo:   repo,
		logger: sublogger,
	}
}
