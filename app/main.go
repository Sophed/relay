package main

import (
	"messaging/app/views"
	"messaging/data/storage"
	"messaging/routes/api"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

const STORAGE_DIR = "storage/"
const PORT = 1337

func main() {
	// defining JSON storage method with file paths
	storage.METHOD = &storage.StorageJSON{
		UsersFile:         STORAGE_DIR + "users.json",
		MessagesFile:      STORAGE_DIR + "messages.json",
		ConversationsFile: STORAGE_DIR + "conversations.json",
	}
	// test storage method and panic on fail
	err := storage.METHOD.Test()
	if err != nil {
		lg.Fatl(err)
	}
	lg.Info("storage checks passed")

	// create fiber instance
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// page routes
	app.Get("/", views.ViewIndex)
	app.Get("/pricing", views.ViewPricing)
	app.Get("/login", views.ViewLogin)
	app.Get("/register", views.ViewRegister)
	app.Get("/app", views.ViewApp)

	// api routes
	app.Post("/api/login", api.Login)
	app.Post("/api/register", api.Register)

	// serve static assets
	app.Static("/static", "static")

	// start http server
	lg.Info("server started at http://127.0.0.1:" + strconv.Itoa(PORT))
	lg.Fatl(app.Listen(":" + strconv.Itoa(PORT)))
}
