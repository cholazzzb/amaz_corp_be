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

func (h *LocationHandler) CreateBuilding(ctx *fiber.Ctx) error {
	userID, ok, resFactory := validator.CheckUserIDJWT(ctx, h.logger)
	if !ok {
		return resFactory.Create()
	}

	req := new(ent.BuildingCommand)
	ok, resFactory = validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	err := h.svc.CreateBuilding(ctx.Context(), req.Name, userID)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, nil)
}

func (h *LocationHandler) GetBuildingByID(ctx *fiber.Ctx) error {
	buildingID := ctx.Params("buildingID")
	building, err := h.svc.GetBuildingByID(ctx.Context(), buildingID)
	if err != nil {
		return response.InternalServerError(ctx)
	}
	return response.Ok(ctx, building)
}

func (h *LocationHandler) GetBuildings(ctx *fiber.Ctx) error {
	bs, err := h.svc.GetBuildings(ctx.Context())
	if err != nil {
		return response.InternalServerError(ctx)
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

func (h *LocationHandler) GetMyInvitation(ctx *fiber.Ctx) error {
	userID, ok, resFactory := validator.CheckUserIDJWT(ctx, h.logger)
	if !ok {
		return resFactory.Create()
	}

	bldngs, err := h.svc.GetMyInvitation(ctx.Context(), userID)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, bldngs)
}

func (h *LocationHandler) GetListMyOwnedBuilding(ctx *fiber.Ctx) error {
	userID, ok, resFactory := validator.CheckUserIDJWT(ctx, h.logger)
	if !ok {
		return resFactory.Create()
	}

	bldngs, err := h.svc.GetListMyOwnedBuilding(ctx.Context(), userID)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, bldngs)
}

type GetBuildingsByUserIDRequest struct {
	UserID string `json:"userID" validate:"required"`
}

func (h *LocationHandler) GetBuildingsByUserID(ctx *fiber.Ctx) error {
	userID, success := ctx.Locals("UserId").(string)
	if !success {
		err := errors.New("failed to get userID from JWT")
		h.logger.Error(err.Error())
		return response.InternalServerError(ctx)
	}

	if len(userID) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "userID is missing from the request",
		})
	}

	req := GetBuildingsByUserIDRequest{userID}
	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	lbs, err := h.svc.GetBuildingsByUserID(ctx.Context(), userID)
	if err != nil {
		h.logger.Error(err.Error())
		return response.InternalServerError(ctx)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "ok",
		"buildings": lbs,
	})
}

func (h *LocationHandler) InviteMemberToBuilding(
	ctx *fiber.Ctx,
) error {
	req := new(ent.InviteMemberToBuildingCommand)
	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	exist, err := h.svc.CheckMemberBuildingExist(ctx.Context(), req.UserID, req.BuildingID)
	if err != nil {
		return response.InternalServerError(ctx)
	}
	if exist {
		return response.BadRequest(ctx, "user already joined in the building")
	}

	err = h.svc.InviteMemberToBuilding(ctx.Context(), "New Member", req.UserID, req.BuildingID)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return nil
}

func (h *LocationHandler) JoinBuildingById(ctx *fiber.Ctx) error {
	req := new(ent.JoinBuildingCommand)
	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	err := h.svc.JoinBuilding(ctx.Context(), req.MemberID, req.BuildingID)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, "ok")
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
		return response.InternalServerError(ctx)
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

func (h *LocationHandler) EditMemberName(ctx *fiber.Ctx) error {
	req := new(ent.RenameMemberCommand)
	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	err := h.svc.EditMemberName(ctx.Context(), req.MemberID, req.Name)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, nil)
}

type GetMemberByNameRequest struct {
	Name string `query:"name"`
}

func (h *LocationHandler) GetMemberByName(ctx *fiber.Ctx) error {
	queryParams := new(GetMemberByNameRequest)
	ok, resFactory := validator.CheckQueryParams(ctx, queryParams)
	if !ok {
		return resFactory.Create()
	}

	member, err := h.svc.GetMemberByName(ctx.Context(), queryParams.Name)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, member)
}

type GetMemberByIDRequest struct {
	MemberID string `json:"memberID" validate:"required"`
}

func (h *LocationHandler) GetMemberByID(ctx *fiber.Ctx) error {
	memberID := ctx.Params("memberID")
	if len(memberID) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "memberID must be filled",
		})
	}

	req := GetMemberByIDRequest{memberID}
	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	member, err := h.svc.GetMemberByID(ctx.Context(), req.MemberID)
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
