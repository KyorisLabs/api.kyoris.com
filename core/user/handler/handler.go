package userHandler

import (
	sharedApplication "api.kyoris.com/core/_shared/application"
	userApplication "api.kyoris.com/core/user/application"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Register() fiber.Handler
	Get() fiber.Handler
}

type userHandler struct {
	userService userApplication.UserService
}

func NewHandler(services *sharedApplication.Services) UserHandler {
	return &userHandler{
		userService: services.UserService,
	}
}
