package friend

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

type FriendHandler struct {
	svc *service.Service
}

func NewFriendHandler(svc *service.Service) *FriendHandler {
	return &FriendHandler{svc}
}

type GetFriendsByMemberIdRequest struct {
	MemberID int64 `json:"memberId" validate:"required"`
}

func (h *FriendHandler) GetFriendsByMemberId(ctx *fiber.Ctx) error {
	mID, err := strconv.ParseInt(ctx.Params("memberId"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "memberId is missing from the request",
		})
	}

	req := GetFriendsByMemberIdRequest{mID}
	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	fs, err := h.svc.Friend.GetFriendsByMemberId(ctx.Context(), mID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"friends": fs,
	})
}
