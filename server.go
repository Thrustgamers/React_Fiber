package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Print(c.IP())
		return c.SendStatus(200)
	})

	app.Listen(":3000")
}
