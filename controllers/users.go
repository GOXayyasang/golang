package controllers

import (
	"database/sql"
	"fmt"

	"github.com/GOXayyasang/golang/database"
	"github.com/GOXayyasang/golang/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUsers(c *fiber.Ctx) error {
	db := database.DB
	var user models.User
	var req models.UserReq
	userDecoded, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 400,
			"resDesc": fmt.Sprintf("Failed to TOKEN"),
		})
	}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 400,
			"resDesc": fmt.Sprintf("Failed to parse body: %s", err.Error()),
		})
	}
	err = db.Get(&user, `SELECT * FROM Users WHERE username = @username`, sql.Named("username", req.Username))
	if err != nil {
		if err != sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"resCode": 400,
				"resDesc": fmt.Sprintf("Failed to check user: %s", err.Error()),
			})
		}
	}
	if user.Name != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 400,
			"resDesc": fmt.Sprintf("This user aready exist"),
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 400,
			"resDesc": fmt.Sprintf("Failed to hash password: %s", err.Error()),
		})
	}
	req.Password = string(hashedPassword)
	query := `
			INSERT INTO Users
				(name,surname,birthdate,username,password,createdBy) 
			VALUES
				(:name,:surname,:birthdate,:username,:password,:createdBy)`

	result, err := db.NamedExec(query, map[string]interface{}{
		"name":      req.Name,
		"surname":   req.Surname,
		"birthdate": req.Birthdate,
		"username":  req.Username,
		"password":  req.Password,
		"createdBy": userDecoded.ID,
	})
	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"resCode": 500,
			"resDesc": fmt.Sprintf("Failed to create user: %s", err.Error()),
		})
	}
	_, err = result.RowsAffected()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"resCode": 500,
			"resDesc": fmt.Sprintf("Failed to create user: %s", err.Error()),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resCode": 200,
		"resDesc": "SUCCESS",
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	db := database.DB
	err := db.Select(&users, "SELECT * FROM Users")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 400,
			"resDesc": fmt.Sprintf("Failed to retrieve roles: %s", err.Error()),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resCode": 200,
		"resDesc": "SUCCESS",
		"data":    users,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var users models.User
	db := database.DB
	err := db.Get(&users, "SELECT * FROM Users WHERE id = @id", sql.Named("id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"resCode": 404,
				"resDesc": fmt.Sprintf("NOT FOUND THIS USER"),
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"resCode": 400,
				"resDesc": fmt.Sprintf("Failed to check user: %s", err.Error()),
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resCode": 200,
		"resDesc": "SUCCESS",
		"data":    users,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	userDecoded := c.Locals("user").(models.User)
	id := c.Params("id")
	var hashedPassword []byte
	db := database.DB
	var user models.User
	var req models.UserReq
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 400,
			"resDesc": fmt.Sprintf("Failed to parse body: %s", err.Error()),
		})
	}
	err = db.Get(&user, `SELECT * FROM Users WHERE username = @username AND id != @id`, sql.Named("username", req.Username), sql.Named("id", id))
	if err != nil {
		if err != sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"resCode": 400,
				"resDesc": fmt.Sprintf("Failed to check user: %s", err.Error()),
			})
		}
	}
	if user.Name != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 400,
			"resDesc": fmt.Sprintf("This user aready exist"),
		})
	}
	query := `UPDATE Users SET name = :name, surname = :surname, birthdate = :birthdate, username = :username, updatedBy = :updatedBy `
	arg := map[string]interface{}{
		"name":      req.Name,
		"surname":   req.Surname,
		"birthdate": req.Birthdate,
		"username":  req.Username,
		"updatedBy": userDecoded.ID,
		"id":        id,
	}
	if req.Password != "" {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"resCode": 400,
				"resDesc": fmt.Sprintf("Failed to hash password: %s", err.Error()),
			})
		}
		query += ", password = :password "
		arg["password"] = string(hashedPassword)
	}
	query += "WHERE id = :id"

	result, err := db.NamedExec(query, arg)
	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"resCode": 500,
			"resDesc": fmt.Sprintf("Failed to update user: %s", err.Error()),
		})
	}
	_, err = result.RowsAffected()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"resCode": 500,
			"resDesc": fmt.Sprintf("Failed to create user: %s", err.Error()),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resCode": 200,
		"resDesc": "SUCCESS",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	_, err := db.Exec("DELETE FROM Users WHERE id = @id", sql.Named("id", id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 400,
			"resDesc": fmt.Sprintf("Failed to delete this user : %s", err.Error()),
		})

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resCode": 200,
		"resDesc": "SUCCESS",
	})
}
