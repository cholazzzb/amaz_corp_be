package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

type UserHandler struct {
	svc    *service.Service
	logger zerolog.Logger
}

func NewUserHandler(svc *service.Service) *UserHandler {
	sublogger := log.With().Str("layer", "repository").Str("package", "user").Logger()

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

	err := h.svc.User.RegisterUser(ctx.Context(), req.Username, req.Password)
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

	token, err := h.svc.User.Login(ctx.Context(), req.Username, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
