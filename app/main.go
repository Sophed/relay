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
	storage.METHOD = &storage.StorageJSON{
		UsersFile:         STORAGE_DIR + "users.json",
		MessagesFile:      STORAGE_DIR + "messages.json",
		ConversationsFile: STORAGE_DIR + "conversations.json",
	}
	err := storage.METHOD.Test()
	if err != nil {
		lg.Fatl(err)
	}
	lg.Info("storage checks passed")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", views.ViewIndex)
	app.Get("/login", views.ViewLogin)
	app.Get("/"+strconv.Itoa(fiber.StatusNotFound), views.ViewErrorNotFound)
	app.Get("/"+strconv.Itoa(fiber.StatusInternalServerError), views.ViewErrorInternal)

	app.Post("/api/register", api.Register)
	app.Post("/api/login", api.Login)

	app.Static("/static", "static")

	lg.Info("server started at http://127.0.0.1:" + strconv.Itoa(PORT))
	lg.Fatl(app.Listen(":" + strconv.Itoa(PORT)))
}
