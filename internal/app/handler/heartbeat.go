package handler

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
)

type HeartbeatHandler struct {
	svc    *service.HeartbeatService
	logger zerolog.Logger
}

func NewHeartBeatHandler(svc *service.HeartbeatService) *HeartbeatHandler {
	sublogger := log.With().Str("layer", "handler").Str("package", "heartbeat").Logger()

	return &HeartbeatHandler{svc: svc, logger: sublogger}
}

func (h *HeartbeatHandler) Pulse(ctx *fiber.Ctx) error {
	userId, success := ctx.Locals("UserId").(float64)

	if !success {
		err := errors.New("userId from JWT is not float")
		h.logger.Error().Err(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			err.Error(),
		)
	}

	err := h.svc.Pulse(ctx.Context(), (int64(userId)))

	if err != nil {
		h.logger.Error().Err(err)
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

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		h.logger.Error().Err(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			err.Error(),
		)
	}

	status, err := h.svc.CheckUserStatus(ctx.Context(), int64(userIdInt))

	if err != nil {
		h.logger.Error().Err(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			err.Error(),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		status:    status,
	})
}
