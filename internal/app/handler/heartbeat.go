package handler

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
)

type HeartbeatHandler struct {
	svc    *service.HeartbeatService
	logger *slog.Logger
}

func NewHeartBeatHandler(svc *service.HeartbeatService) *HeartbeatHandler {
	sublogger := logger.Get().With(slog.String("layer", "handler"), slog.String("package", "heartbeat"))

	return &HeartbeatHandler{svc: svc, logger: sublogger}
}

func (h *HeartbeatHandler) Pulse(ctx *fiber.Ctx) error {
	userId, success := ctx.Locals("UserId").(string)

	if !success {
		err := errors.New("failed to get userId from JWT")
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			err.Error(),
		)
	}

	err := h.svc.Pulse(ctx.Context(), userId)

	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			err.Error(),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

func (h *HeartbeatHandler) GetStatusByUserId(ctx *fiber.Ctx) error {
	userId := string(
		[]byte(strings.Trim(ctx.Params("userId"), " ")),
	)

	status, err := h.svc.CheckUserStatus(ctx.Context(), userId)

	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			err.Error(),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		status:    status,
	})
}
