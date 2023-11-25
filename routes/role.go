package routes

import (
	"github.com/GOXayyasang/golang/controllers"
	"github.com/GOXayyasang/golang/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoleRoute(app *fiber.App) {
	app.Get("/roles", middlewares.VerifyToken, controllers.GetRoles)
	app.Get("/role", controllers.GetRoles)
}
