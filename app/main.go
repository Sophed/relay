package main

import (
	"messaging/routes/views"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/login", views.PageLogin)
	app.Get("/"+strconv.Itoa(fiber.StatusInternalServerError), views.PageErrorInternal)

	lg.Fatl(app.Listen(":1337"))
}
