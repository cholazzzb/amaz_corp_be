package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

type UserHandler struct {
	svc    *service.UserService
	logger *slog.Logger
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	sublogger := logger.Get().With(slog.String("domain", "user"), slog.String("layer", "handler"))

	return &UserHandler{svc: svc, logger: sublogger}
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=8,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := ctx.BodyParser(req); err != nil {
		h.logger.Error(err.Error())
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
	Username string `json:"username" validate:"required,min=8,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	req := new(LoginRequest)
	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	token, err := h.svc.Login(ctx.Context(), req.Username, req.Password)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// TODO: Handle to give response if user not exist

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *UserHandler) CheckUserExistance(ctx *fiber.Ctx) error {
	username := ctx.Params("userId")
	if len(username) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "userId must be filled",
		})
	}

	exist, err := h.svc.CheckUserExistance(ctx.Context(), username)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			err.Error(),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"exist":   exist,
	})
}
