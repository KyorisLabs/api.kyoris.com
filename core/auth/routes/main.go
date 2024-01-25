package authRoutes

import (
	sharedApplication "api.kyoris.com/core/_shared/application"
	authHandler "api.kyoris.com/core/auth/handler"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, services *sharedApplication.Services) {
	_authHandler := authHandler.NewHandler(services)

	app.Post("/auth/register", _authHandler.Register())
	app.Post("/auth/login", _authHandler.Login())
	app.Post("/auth/logout", _authHandler.Logout())
}
