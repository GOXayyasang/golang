package routes

import (
	"github.com/GOXayyasang/golang/controllers"
	"github.com/GOXayyasang/golang/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoute(app *fiber.App) {
	app.Post("/user", middlewares.VerifyToken, controllers.CreateUsers)
	app.Get("/users", middlewares.VerifyToken, controllers.GetUsers)
	app.Get("/user/:id", middlewares.VerifyToken, controllers.GetUser)
	app.Put("/user/:id", middlewares.VerifyToken, controllers.UpdateUser)
	app.Delete("/user/:id", middlewares.VerifyToken, controllers.DeleteUser)
}
