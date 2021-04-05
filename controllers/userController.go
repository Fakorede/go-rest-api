package controllers

import (
	"goadmin/database"
	"goadmin/middlewares"
	"goadmin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Users(c *fiber.Ctx) error {
	if err := middlewares.Authorize(c, "users"); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.User{}, page))
}

func GetUser(c *fiber.Ctx) error {
	if err := middlewares.Authorize(c, "users"); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id := c.Params("id")

	var user models.User
	database.DB.Preload("Role").Where("id = ?", id).First(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found!",
		})
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	if err := middlewares.Authorize(c, "users"); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "FirstName, LastName and Email is required",
		})
	}

	user.SetPassword("password")

	database.DB.Create(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already exists!",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	if err := middlewares.Authorize(c, "users"); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Model(&user).Updates(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	if err := middlewares.Authorize(c, "users"); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return c.SendStatus(fiber.StatusNoContent)
}
