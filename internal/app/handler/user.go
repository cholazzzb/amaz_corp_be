package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/model"
	"github.com/cholazzzb/amaz_corp_be/internal/app/service"

	custom_logger "github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/response"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

type UserHandler struct {
	svc    *service.UserService
	logger *slog.Logger
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	sublogger := custom_logger.Get().With(slog.String("domain", "user"), slog.String("layer", "handler"))

	return &UserHandler{
		svc:    svc,
		logger: sublogger,
	}
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	req := new(model.RegisterRequest)

	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	err := h.svc.Register(ctx.Context(), req)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, nil)
}
