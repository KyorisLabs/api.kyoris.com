package authHandler

import (
	sharedApplication "api.kyoris.com/core/_shared/application"
	authApplication "api.kyoris.com/core/auth/application"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Register() fiber.Handler
	// Login() fiber.Handler
}

type authHandler struct {
	authService authApplication.AuthService
}

func NewHandler(services *sharedApplication.Services) AuthHandler {
	return &authHandler{
		authService: services.AuthService,
	}
}
