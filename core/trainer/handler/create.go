package trainerHandler

import (
	"strings"

	sharedDomain "api.kyoris.com/core/_shared/domain"
	"api.kyoris.com/core/_shared/domain/response"
	trainerDomain "api.kyoris.com/core/trainer/domain"
	"github.com/gofiber/fiber/v2"
)

func (t *trainerHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload *trainerDomain.CreateTrainerDTO
		if err := c.BodyParser(&payload); err != nil {
			// En caso de error, devolvemos un error 400 y el codigo de error
			return c.Status(fiber.StatusBadRequest).JSON(response.Error{
				Error: sharedDomain.ErrBodyNotValid.Error(),
			})
		}

		if payload.Name == "" {
			return c.Status(fiber.StatusBadRequest).JSON(response.Error{
				Error: trainerDomain.ErrTrainerNameIsEmpty.Error(),
			})
		}

		payload.Name = strings.TrimSpace(payload.Name)

		// Añadimos el usuario ID del header obtenido en el middleware
		payload.UserID = 1

		err := t.trainerService.Create(payload)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.Error{
				Error: err.Error(),
			})
		}
		return nil
	}
}
