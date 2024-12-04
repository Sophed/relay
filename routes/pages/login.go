package pages

import (
	"messaging/static"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

func PageLogin(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	page, err := static.LoginPage()
	if err != nil {
		lg.Warn(err)
		return c.Redirect("/"+strconv.Itoa(fiber.StatusInternalServerError), fiber.StatusTemporaryRedirect)
	}
	return c.SendString(page)
}
