package main

import (
	"log"
	"os"

	"github.com/GOXayyasang/golang/database"
	"github.com/GOXayyasang/golang/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Some error occured. Err: %s", err)
	}
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("HELLO")
	})

	routes.SetupRoleRoute(app)
	routes.SetupUserRoute(app)
	routes.SetupAuthRoute(app)

	app.Listen(os.Getenv("SERVER_PORT"))
}
