package response

import "github.com/gofiber/fiber/v2"

func InternalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(
		"Internal Server Error",
	)
}
