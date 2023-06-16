package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

type LocationHandler struct {
	svc    *service.LocationService
	logger zerolog.Logger
}

func NewLocationHandler(svc *service.LocationService) *LocationHandler {
	sublogger := log.With().Str("layer", "handler").Str("package", "location").Logger()

	return &LocationHandler{svc: svc, logger: sublogger}
}

type GetBuildingsByMemberIdRequest struct {
	MemberId int64 `json:"memberId" validate:"required"`
}

func (h *LocationHandler) GetBuildingsByMemberId(ctx *fiber.Ctx) error {
	memberId, err := strconv.ParseInt(ctx.Query("memberId"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "memberId is invalid from the request",
		})
	}

	req := GetBuildingsByMemberIdRequest{memberId}
	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	lbs, err := h.svc.GetBuildingsByMemberId(ctx.Context(), memberId)
	if err != nil {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "ok",
		"buildings": lbs,
	})
}

type GetRoomsByBuildingIdRequest struct {
	BuildingId int64 `json:"buildingId" validate:"required"`
}

func (h *LocationHandler) GetRoomsByBuildingId(ctx *fiber.Ctx) error {
	buildingId, err := strconv.ParseInt(ctx.Params("buildingId"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "buildingId is missing from the request",
		})
	}

	req := GetRoomsByBuildingIdRequest{buildingId}
	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	rs, err := h.svc.GetRoomsByBuildingId(ctx.Context(), buildingId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"rooms":   rs,
	})
}

type GetListOnlineMembersRequest struct {
	RoomId int64 `json:"roomId" validate:"required"`
}

func (h *LocationHandler) GetListOnlineMembers(ctx *fiber.Ctx) error {
	roomId, err := strconv.ParseInt(ctx.Params("roomId"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "roomId is missing from the request",
		})
	}

	req := GetListOnlineMembersRequest{roomId}
	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	loms, err := h.svc.GetListOnlineMembers(ctx.Context(), roomId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"members": loms,
	})
}
