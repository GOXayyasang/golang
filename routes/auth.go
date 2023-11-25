package routes

import (
	"github.com/GOXayyasang/golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoute(app *fiber.App) {
	app.Post("/login", controllers.Login)
}
