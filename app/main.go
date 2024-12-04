package main

import (
	"messaging/routes/pages"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/login", pages.PageLogin)
	app.Get("/"+strconv.Itoa(fiber.StatusInternalServerError), pages.PageErrorInternal)

	lg.Fatl(app.Listen(":1337"))
}
