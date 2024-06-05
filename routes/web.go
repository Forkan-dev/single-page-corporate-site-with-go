package routes

import (
	"blog/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func SetupRoutes() *fiber.App {
	app := fiber.New()
	engine := html.New("view", ".html")

	app = fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/assets", "./public/assets")
	app.Static("/js", "./public/js")

	// ==============Routes started=================

	// User routes
	app.Get("/", controllers.GetHome)

	// Add more routes as needed

	return app
}
