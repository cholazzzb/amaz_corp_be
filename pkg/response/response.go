package response

import (
	"github.com/gofiber/fiber/v2"
)

func RequestNotValid(ctx *fiber.Ctx, validationError interface{}) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message":          "Request Not Valid",
		"validation error": validationError,
	})
}

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
