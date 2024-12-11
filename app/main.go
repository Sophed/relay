package main

import (
	"messaging/app/views"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

func main() {
	app := fiber.New()

	app.Get("/", views.ViewIndex)
	app.Get("/login", views.ViewLogin)
	app.Get("/"+strconv.Itoa(fiber.StatusNotFound), views.ViewErrorNotFound)
	app.Get("/"+strconv.Itoa(fiber.StatusInternalServerError), views.ViewErrorInternal)

	app.Static("/static", "static")

	lg.Fatl(app.Listen(":1337"))
}
