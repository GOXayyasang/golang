package controllers

import (
	"fmt"

	"github.com/GOXayyasang/golang/database"
	"github.com/GOXayyasang/golang/middlewares"
	"github.com/GOXayyasang/golang/models"
	"github.com/gofiber/fiber/v2"
)

func GetRoles(c *fiber.Ctx) error {
	db := database.DB
	role := c.Locals("user")
	roleDecoded := role.(models.Role)
	fmt.Printf("%d", *roleDecoded.UpdatedBy)
	var roles []models.Role
	// db, err := database.InitDB()
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"resCode": 500,
	// 		"resDesc": fmt.Sprintf("CONNECT DB ERROR %s", err.Error()),
	// 	})
	// }
	err := db.Select(&roles, "SELECT * FROM Roles")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resCode": 400,
			"resDesc": fmt.Sprintf("Failed to retrieve roles: %s", err.Error()),
		})
	}
	// defer rows.Close()
	// for rows.Next() {
	// 	var role models.Role
	// 	err := rows.Scan(&role.ID, &role.Name, &role.Status, &role.CreatedBy, &role.CreatedAt, &role.UpdatedAt, &role.UpdatedBy)
	// 	if err != nil {
	// 		fmt.Println("Error scanning role:", err)
	// 		continue
	// 	}
	// 	roles = append(roles, role)
	// }

	// if err := rows.Err(); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"resCode": 500,
	// 		"resDesc": fmt.Sprintf("Failed to retrieve roles: %s", err.Error()),
	// 	})
	// }
	token, err := middlewares.GenerateToken(models.User{})
	if err != nil {
		fmt.Println(err)
	}
	roles[0].Token = token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resCode": 200,
		"resDesc": "SUCCESS",
		"data":    roles,
	})
}
