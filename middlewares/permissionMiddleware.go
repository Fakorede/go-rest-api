package middlewares

import (
	"errors"
	"goadmin/database"
	"goadmin/models"
	util "goadmin/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Authorize(c *fiber.Ctx, model string) error {
	cookie := c.Cookies("jwt")

	id, err := util.ParseJWT(cookie)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	database.DB.Preload("Permissions").Find(&role)

	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+model || permission.Name == "edit_"+model {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			if permission.Name == "edit_"+model {
				return nil
			}
		}
	}

	return errors.New("you are not authorized to perform this action")
}
