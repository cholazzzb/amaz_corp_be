package user

import (
	"github.com/cholazzzb/amaz_corp_be/validator"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc *UserService
}

func NewUserHandler(svc *UserService) *UserHandler {
	return &UserHandler{svc}
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
