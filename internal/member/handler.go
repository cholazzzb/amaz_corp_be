package member

import (
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type MemberHandler struct {
	svc *MemberService
}

func NewMemberHandler(svc *MemberService) *MemberHandler {
	return &MemberHandler{svc}
}

type GetMemberByNameRequest struct {
	Name string `json:"name" validate:"required"`
}

func (h *MemberHandler) GetMemberByName(ctx *fiber.Ctx) error {
	// username := ctx.Locals("Username")

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
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"member":  member,
	})
}
