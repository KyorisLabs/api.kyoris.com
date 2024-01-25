package authHandler

import (
	"net/http"

	"api.kyoris.com/core/_shared/domain/response"
	authDomain "api.kyoris.com/core/auth/domain"
	tokenUtils "api.kyoris.com/utils/token"
	"github.com/gofiber/fiber/v2"
)

func (a *authHandler) Logout() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")

		// Comprobamos el token
		token, ok := tokenUtils.ValidateToken(token)
		if !ok {
			return ctx.Status(http.StatusBadRequest).JSON(response.Error{
				Error: authDomain.ErrEmptyToken.Error(),
			})
		}

		err := a.authService.Logout(token)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(response.Error{
				Error: err.Error(),
			})
		}

		return ctx.Status(http.StatusOK).JSON(response.Default{
			Message: "Sesión cerrada correctamente",
		})
	}
}
