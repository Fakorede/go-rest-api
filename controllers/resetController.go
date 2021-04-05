package controllers

import (
	"goadmin/database"
	"goadmin/models"
	"log"
	"math/rand"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token := GenerateRandStringRunes(12)

	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}

	database.DB.Create(&passwordReset)

	from := "support@goadmin.com"

	to := []string{
		data["email"],
	}

	url := "http://localhost:3000/reset/" + token

	message := "Click <a href=\"" + url + "\">here</a> to reset your password."

	err := smtp.SendMail("0.0.0.0:1025", nil, from, to, []byte(message))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "A password reset token has been sent to your email",
	})
}

func Reset(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirmation"] {
		return c.Status(400).JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}

	var passwordReset = models.PasswordReset{}

	err := database.DB.Where("token = ?", data["token"]).Last(&passwordReset)
	if err.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Token!",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{}

	database.DB.Where("email = ?", passwordReset.Email).First(&user)
	database.DB.Model(&user).Where("email = ?", passwordReset.Email).Update("password", password)

	if user.Id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Credentials!",
		})
	}

	log.Print(user.FirstName)

	return c.JSON(fiber.Map{
		"message": "Password reset successful!",
	})

}

func GenerateRandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
