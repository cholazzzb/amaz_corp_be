package handler

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

type LocationHandler struct {
	svc    *service.LocationService
	logger *slog.Logger
}

func NewLocationHandler(svc *service.LocationService) *LocationHandler {
	sublogger := logger.Get().With(slog.String("domain", "location"), slog.String("layer", "handler"))

	return &LocationHandler{svc: svc, logger: sublogger}
}

func (h *LocationHandler) GetBuildings(ctx *fiber.Ctx) error {
	bs, err := h.svc.GetBuildings(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "ok",
		"buildings": bs,
	})
}

type DeleteBuildingRequest struct {
	MemberId   string `json:"memberId" validate:"required"`
	BuildingId string `json:"buildingId" validate:"required"`
}

func (h *LocationHandler) DeleteBuilding(ctx *fiber.Ctx) error {
	req := new(DeleteBuildingRequest)
	if err := ctx.BodyParser(req); err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err := h.svc.DeleteBuilding(ctx.Context(), req.BuildingId, req.MemberId)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return nil
}

type GetBuildingsByMemberIdRequest struct {
	MemberId string `json:"memberId" validate:"required"`
}

func (h *LocationHandler) GetBuildingsByMemberId(ctx *fiber.Ctx) error {
	memberId, success := ctx.Locals("UserId").(string)
	if !success {
		err := errors.New("failed to get userId from JWT")
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			err.Error(),
		)
	}

	if len(memberId) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "memberId is missing from the request",
		})
	}

	req := GetBuildingsByMemberIdRequest{memberId}
	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	lbs, err := h.svc.GetBuildingsByMemberId(ctx.Context(), memberId)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "ok",
		"buildings": lbs,
	})
}

type JoinBuildingRequest struct {
	MemberId   string `json:"memberId" validate:"required"`
	BuildingId string `json:"buildingId" validate:"required"`
}

func (h *LocationHandler) JoinBuildingById(ctx *fiber.Ctx) error {
	req := new(JoinBuildingRequest)
	if err := ctx.BodyParser(req); err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := h.svc.JoinBuilding(ctx.Context(), req.MemberId, req.BuildingId)
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

type GetRoomsByBuildingIdRequest struct {
	BuildingId string `json:"buildingId" validate:"required"`
}

func (h *LocationHandler) GetRoomsByBuildingId(ctx *fiber.Ctx) error {
	buildingId := ctx.Params("buildingId")
	if len(buildingId) == 0 {
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
		h.logger.Error(err.Error())
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
	RoomId string `json:"roomId" validate:"required"`
}

func (h *LocationHandler) GetListOnlineMembers(ctx *fiber.Ctx) error {
	roomId := ctx.Params("roomId")
	if len(roomId) == 0 {
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
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"members": loms,
	})
}
