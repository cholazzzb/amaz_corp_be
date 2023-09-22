package handler

import (
	"log/slog"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type RemoteConfigHandler struct {
	svc    *service.RemoteConfigService
	logger *slog.Logger
}

func NewRemoteConfigHandler(svc *service.RemoteConfigService) *RemoteConfigHandler {
	sublogger := logger.Get().With(
		slog.String("domain", "remoteconfig"),
		slog.String("layer", "handler"),
	)

	return &RemoteConfigHandler{
		svc:    svc,
		logger: sublogger,
	}
}

func (h *RemoteConfigHandler) GetAPKVersion(ctx *fiber.Ctx) error {
	av, err := h.svc.GetAPKVersion(ctx.Context())
	if err != nil {
		h.logger.Error(err.Error())
		return err
	}

	return response.Ok(ctx, av)
}
