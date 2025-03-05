package main

import (
	"messaging/cmd/internal/api"
	"messaging/cmd/internal/pages"

	"github.com/gofiber/fiber/v2"
)

func routes(app *fiber.App) {
	// page routes
	app.Get("/", pages.Index)
	app.Get("/pricing", pages.Pricing)
	app.Get("/login", pages.Login)
	app.Get("/register", pages.Register)
	app.Get("/app", pages.App)

	// api routes
	app.Post("/api/login", api.Login)
	app.Post("/api/register", api.Register)

	// serve static assets
	app.Static("/static", "static")
}
