package middlewares

import (
	"net/http"

	sharedApplication "api.kyoris.com/core/_shared/application"
	"api.kyoris.com/core/_shared/domain/response"
	"github.com/gofiber/fiber/v2"
)

func CheckToken(services *sharedApplication.Services) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userId int
		var err error

		// Obtenemos el header de Autorización
		authToken := c.Get("Authorization")

		// Validamos el token
		if userId, err = services.AuthService.CheckToken(authToken); err != nil {
			return c.Status(http.StatusUnauthorized).JSON(response.Error{
				Error: err.Error(),
			})
		}

		c.Locals("userID", userId)

		return c.Next()
	}
}
