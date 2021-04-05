package controllers

import (
	"goadmin/database"
	"goadmin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Roles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Preload("Permissions").Find(&roles)

	return c.JSON(roles)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var role models.Role
	database.DB.Preload("Permissions").Where("id = ?", id).First(&role)

	if role.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Role not found!",
		})
	}

	return c.JSON(role)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	list := roleDTO["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		// id, _ := strconv.Atoi(permissionId.(string))
		id := int(permissionId.(float64))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Role created successfully",
		"role":    role,
	})
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	list := roleDTO["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	var result interface{}
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}

	// database.DB.Where("id = ?", id).First(&role)

	// if role.Id == 0 {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"message": "Role not found!",
	// 	})
	// }

	database.DB.Model(&role).Updates(role)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Role updated successfully",
		"role":    role,
	})
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return c.SendStatus(fiber.StatusNoContent)
}
