package trainerHandler

import (
	sharedApplication "api.kyoris.com/core/_shared/application"
	trainerApplication "api.kyoris.com/core/trainer/application"
	"github.com/gofiber/fiber/v2"
)

type TrainerHandler interface {
	Create() fiber.Handler
}

type trainerHandler struct {
	trainerService trainerApplication.TrainerService
}

func NewHandler(services *sharedApplication.Services) TrainerHandler {
	return &trainerHandler{
		trainerService: services.TrainerService,
	}
}
