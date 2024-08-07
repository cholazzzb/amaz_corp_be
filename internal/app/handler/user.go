package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/response"
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

type RegisterResponse struct {
	UserID string `json:"userID"`
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := ctx.BodyParser(req); err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	userID, err := h.svc.RegisterUser(ctx.Context(), req.Username, req.Password, 2)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, RegisterResponse{userID})
}

func (h *UserHandler) RegisterAdmin(ctx *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := ctx.BodyParser(req); err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	userID, err := h.svc.RegisterUser(ctx.Context(), req.Username, req.Password, 1)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, RegisterResponse{userID})
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

	token, err := h.svc.Login(ctx.Context(), req.Username, req.Password, 2)
	if err != nil {
		h.logger.Error(err.Error())
		return response.InternalServerError(ctx)
	}
	// TODO: Handle to give response if user not exist

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *UserHandler) LoginAdmin(ctx *fiber.Ctx) error {
	req := new(LoginRequest)
	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	token, err := h.svc.Login(ctx.Context(), req.Username, req.Password, 1)
	if err != nil {
		return response.InternalServerError(ctx)
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

type GetUserByNameReq struct {
	username string `query:"username"`
}

func (h *UserHandler) GetListUserByUsername(ctx *fiber.Ctx) error {
	queryParams := new(GetUserByNameReq)
	ok, resFactory := validator.CheckQueryParams(ctx, queryParams)
	if !ok {
		return resFactory.Create()
	}

	res, err := h.svc.GetListUserByUsername(ctx.Context(), queryParams.username)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, res)
}
