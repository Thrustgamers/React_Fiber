package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./web/dist/")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("test")
	})

	app.Get("/user/checklogin/:id", func(c *fiber.Ctx) error {

		return c.SendString("test")
	})

	app.Listen(":3000")
}
