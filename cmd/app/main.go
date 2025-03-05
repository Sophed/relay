package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

const PORT = 1337

func main() {
	// test storage method and panic on fail
	err := storageType("json")
	if err != nil {
		lg.Fatl(err)
	}
	lg.Info("storage checks passed")

	// create fiber instance
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	routes(app)

	// start http server
	lg.Info("server started at http://127.0.0.1:" + strconv.Itoa(PORT))
	lg.Fatl(app.Listen(":" + strconv.Itoa(PORT)))
}
