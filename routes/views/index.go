package views

import (
	"messaging/static/pages"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

func PageIndex(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	page, err := pages.LoginPage()
	if err != nil {
		lg.Warn(err)
		return c.Redirect("/"+strconv.Itoa(fiber.StatusInternalServerError), fiber.StatusTemporaryRedirect)
	}
	return c.SendString(page)
}
