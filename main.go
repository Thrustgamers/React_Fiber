package main

import (
	"api/database"
	"api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var app *fiber.App

func main() {
	database.ConnectToDb()
	app = fiber.New()
	app.Use(cors.New())
	routes.SetupRoutes(app)
	app.Get("/metrics", monitor.New())
	app.Static("/", "./web/dist/")
	app.Listen(":80")
}
