package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	files := form.File["images"]
	var filename string

	for _, file := range files {
		filename = file.Filename

		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/" + filename,
	})
}
