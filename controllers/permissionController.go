package controllers

import (
	"goadmin/database"
	"goadmin/models"

	"github.com/gofiber/fiber/v2"
)

func Permissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}

func CreatePermission(c *fiber.Ctx) error {
	var permission models.Permission

	if err := c.BodyParser(&permission); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if permission.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Name is required",
		})
	}

	database.DB.Create(&permission)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Permission created successfully",
		"permission": permission,
	})
}
