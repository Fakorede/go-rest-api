package controllers

import (
	"goadmin/database"
	"goadmin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product
	database.DB.Where("id = ?", id).First(&product)

	if product.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found!",
		})
	}

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if product.Title == "" || product.Description == "" || product.Price == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Title, Description and Price is required",
		})
	}

	database.DB.Create(&product)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Model(&product).Updates(&product)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"product": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)

	return c.SendStatus(fiber.StatusNoContent)
}
