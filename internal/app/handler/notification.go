package handler

import (
	"log/slog"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type NotificationHandler struct {
	svc    *service.NotificationService
	logger *slog.Logger
}

func NewNotificationHandler(
	svc *service.NotificationService,
) *NotificationHandler {
	sublogger := logger.Get().With(
		slog.String("domain", "notification"),
		slog.String("layer", "handler"),
	)

	return &NotificationHandler{
		svc:    svc,
		logger: sublogger,
	}
}
