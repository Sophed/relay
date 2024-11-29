package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	lg.Fatl(app.Listen(":1337"))
}
