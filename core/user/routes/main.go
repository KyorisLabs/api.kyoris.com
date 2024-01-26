package userRoutes

import (
	sharedApplication "api.kyoris.com/core/_shared/application"
	userHandler "api.kyoris.com/core/user/handler"
	"api.kyoris.com/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, services *sharedApplication.Services) {
	_userHandler := userHandler.NewHandler(services)

	app.Get("/user", middlewares.CheckToken(services), _userHandler.GetDataByToken())
}
