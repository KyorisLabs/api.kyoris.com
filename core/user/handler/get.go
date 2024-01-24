package userHandler

import "github.com/gofiber/fiber/v2"

// Get implements UserHandler.
func (*userHandler) Get() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	}
}
