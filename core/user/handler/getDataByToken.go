package userHandler

import (
	"net/http"

	"api.kyoris.com/core/_shared/domain/response"
	userDomain "api.kyoris.com/core/user/domain"
	"github.com/gofiber/fiber/v2"
)

// Get implements UserHandler.
func (h *userHandler) GetDataByToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID, _ := ctx.Locals("userID").(int)

		userModel := h.userService.GetById(userID)

		if userModel == nil {
			return ctx.Status(http.StatusBadRequest).JSON(response.Error{
				Error: userDomain.ErrGetUserData.Error(),
			})
		}

		return ctx.JSON(response.Data{
			Data:    userModel,
			Message: "User found",
		})

	}
}
