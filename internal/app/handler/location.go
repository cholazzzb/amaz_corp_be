package handler

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/location"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/response"
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
	MemberId   string `json:"memberID" validate:"required"`
	BuildingId string `json:"buildingID" validate:"required"`
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

func (h *LocationHandler) JoinBuildingById(ctx *fiber.Ctx) error {
	userID, success := ctx.Locals("UserId").(string)
	if !success {
		err := errors.New("failed to get userId from JWT")
		h.logger.Error(err.Error())
		return response.InternalServerError(ctx)
	}
	req := new(ent.JoinBuildingCommand)
	if err := ctx.BodyParser(req); err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	exist, err := h.svc.CheckMemberBuildingExist(ctx.Context(), userID, req.BuildingId)
	if err != nil {
		h.logger.Error(err.Error())
		return response.InternalServerError(ctx)
	}
	if exist {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user already joined in the building",
		})
	}

	err = h.svc.JoinBuilding(ctx.Context(), req.Name, userID, req.BuildingId)
	if err != nil {
		h.logger.Error(err.Error())
		return response.InternalServerError(ctx)
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

func (h *LocationHandler) GetListMemberByBuildingID(ctx *fiber.Ctx) error {
	buildingID := ctx.Params("buildingID")
	if len(buildingID) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "buildingID is missing from the request",
		})
	}

	ms, err := h.svc.GetListMemberByBuildingID(ctx.Context(), buildingID)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"members": ms,
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

type GetMemberByNameRequest struct {
	Name string `json:"name" validate:"required"`
}

func (h *LocationHandler) GetMemberByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	if len(name) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "name must be filled",
		})
	}

	req := GetMemberByNameRequest{name}
	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	member, err := h.svc.GetMemberByName(ctx.Context(), req.Name)
	if err != nil {
		h.logger.Error(err.Error())
		return response.InternalServerError(ctx)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"member":  member,
	})
}

type GetFriendsByMemberIdRequest struct {
	MemberID string `json:"memberId" validate:"required"`
}

func (h *LocationHandler) GetFriendsByMemberId(ctx *fiber.Ctx) error {
	mID := ctx.Params("memberId")
	if len(mID) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "memberId is missing from the request",
		})
	}

	req := GetFriendsByMemberIdRequest{mID}
	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	fs, err := h.svc.GetFriendsByMemberId(ctx.Context(), mID)
	if err != nil {
		h.logger.Error(err.Error())
		return response.InternalServerError(ctx)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"friends": fs,
	})
}
