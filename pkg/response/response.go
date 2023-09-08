package response

import "github.com/gofiber/fiber/v2"

func InternalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal Server Error",
	})
}

func Ok(ctx *fiber.Ctx, data map[string]interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    data,
	})
}
