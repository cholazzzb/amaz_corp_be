package response

import (
	"github.com/gofiber/fiber/v2"
)

// body request / query params / path params is in wrong format
func RequestNotValid(ctx *fiber.Ctx, validationError interface{}) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message":          "Request Not Valid",
		"validation error": validationError,
	})
}

// cannot parse body request / query params / path params
func BadRequest(ctx *fiber.Ctx, message interface{}) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
	})
}

func InternalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal Server Error",
	})
}

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    data,
	})
}
