package controllers

import (
	"database/sql"
	"fmt"

	"github.com/GOXayyasang/golang/database"
	"github.com/GOXayyasang/golang/middlewares"
	"github.com/GOXayyasang/golang/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var req models.ReqLogin
	var user models.User
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 400,
			"resDesc": fmt.Sprintf("Failed to parse body: %s", err.Error()),
		})
	}
	db := database.DB
	err = db.Get(&user, `SELECT * FROM Users WHERE username = @username`, sql.Named("username", req.Username))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"resCode": 404,
				"resDesc": fmt.Sprintf("Not found this user: %s", req.Username),
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"resCode": 401,
				"resDesc": fmt.Sprintf("Failed to check user: %s", err.Error()),
			})
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 402,
			"resDesc": fmt.Sprintf("Username or password are not found"),
		})
	}
	user.Password = ""
	token, err := middlewares.GenerateToken(user)
	if err != nil {
		fmt.Println(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resCode": 200,
		"resDesc": "SUCCESS",
		"data":    user,
		"token":   token,
	})

}
