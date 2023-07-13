package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/internal/domain/user"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

type UserHandler struct {
	svc    *service.UserService
	logger zerolog.Logger
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	sublogger := log.With().Str("layer", "handler").Str("package", "user").Logger()

	return &UserHandler{svc: svc, logger: sublogger}
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=8,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	err := h.svc.RegisterUser(ctx.Context(), req.Username, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

type LoginRequest struct {
	Username string
	Password string
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := h.svc.Login(ctx.Context(), req.Username, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

type GetMemberByNameRequest struct {
	Name string `json:"name" validate:"required"`
}

func (h *UserHandler) GetMemberByName(ctx *fiber.Ctx) error {
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

func (h *UserHandler) CreateMemberByUsername(ctx *fiber.Ctx) error {
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
	member, err := h.svc.CreateMember(ctx.Context(), user.Member{
		Name:   req.Name,
		Status: "new member",
	}, username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"member":  member,
	})
}

type GetFriendsByMemberIdRequest struct {
	MemberID string `json:"memberId" validate:"required"`
}

func (h *UserHandler) GetFriendsByMemberId(ctx *fiber.Ctx) error {
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"friends": fs,
	})
}
