package routes

import (
	"goadmin/controllers"
	"goadmin/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// public routes
	app.Get("/", controllers.Home)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/forgot", controllers.Forgot)
	app.Post("/api/reset", controllers.Reset)

	app.Use(middlewares.Authenticate)

	// private routes
	// users
	app.Put("/api/users/profile", controllers.UpdateProfile)
	app.Put("/api/users/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.Users)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	// roles
	app.Get("/api/roles", controllers.Roles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	// permissions
	app.Get("/api/permissions", controllers.Permissions)
	app.Post("/api/permissions", controllers.CreatePermission)

	// products
	app.Get("/api/products", controllers.Products)
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

	// orders
	app.Get("/api/orders", controllers.Orders)
	app.Get("/api/chart", controllers.Chart)
	app.Post("/api/export", controllers.Export)

	// image
	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/uploads", "./uploads")

	// 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
}
