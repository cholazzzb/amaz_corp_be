package member

import (
	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/member"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

type MemberHandler struct {
	svc *service.Service
}

func NewMemberHandler(svc *service.Service) *MemberHandler {
	return &MemberHandler{svc}
}

type GetMemberByNameRequest struct {
	Name string `json:"name" validate:"required"`
}

func (h *MemberHandler) GetMemberByName(ctx *fiber.Ctx) error {
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

	member, err := h.svc.Member.GetMemberByName(ctx.Context(), req.Name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			err.Error(),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"member":  member,
	})
}

type CreateMemberByUsernameRequest struct {
	Name string `json:"name" validate:"required"`
}

func (h *MemberHandler) CreateMemberByUsername(ctx *fiber.Ctx) error {
	req := new(CreateMemberByUsernameRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	username := ctx.Locals("Username").(string)
	member, err := h.svc.Member.CreateMember(ctx.Context(), member.Member{
		Name:   req.Name,
		Status: "new member",
	}, username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			err.Error(),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"member":  member,
	})
}
