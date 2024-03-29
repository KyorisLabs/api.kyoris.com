package routes

import (
	sharedApplication "api.kyoris.com/core/_shared/application"
	authRoutes "api.kyoris.com/core/auth/routes"
	userRoutes "api.kyoris.com/core/user/routes"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, services *sharedApplication.Services) {

	userRoutes.Setup(app, services)
	authRoutes.Setup(app, services)

}
