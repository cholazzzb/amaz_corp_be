package validator

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/pkg/response"
)

type ok bool

func CheckUserIDJWT(ctx *fiber.Ctx, logger *slog.Logger) (string, ok, response.ResponseFactory) {
	uID, success := ctx.Locals("UserId").(string)
	if !success {
		err := errors.New("failed to get userId from JWT")
		logger.Error(err.Error())
		return "", false, response.Make(response.EInternalServerError, ctx, nil)
	}

	if len(uID) == 0 {
		err := errors.New("len(userID) is 0 from JWT")
		logger.Error(err.Error())
		return "", false, response.Make(response.EBadRequest, ctx, err)
	}

	return uID, true, response.Make(response.Noop, ctx, nil)
}

func CheckSchema[T comparable](ctx *fiber.Ctx, reqSchema T) (ok, response.ResponseFactory) {
	if errs := Validate(reqSchema); len(errs) > 0 {
		return false, response.Make(response.ERequestNotValid, ctx, errs)
	}

	return true, response.Make(response.Noop, ctx, nil)
}

func CheckReqBodySchema(
	ctx *fiber.Ctx,
	reqSchema interface{},
) (ok, response.ResponseFactory) {
	if err := ctx.BodyParser(&reqSchema); err != nil {
		return false, response.Make(response.EBadRequest, ctx, "body request is in wrong format")
	}

	return CheckSchema(ctx, reqSchema)
}
