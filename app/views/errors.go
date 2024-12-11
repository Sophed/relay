package views

import (
	"bytes"
	"messaging/app/components"
	"strconv"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ViewErrorInternal(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	page := new(bytes.Buffer)
	PageError(fiber.StatusInternalServerError, "Internal server error").Render(page)
	return c.SendString(page.String())
}

func ViewErrorNotFound(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	page := new(bytes.Buffer)
	PageError(fiber.StatusNotFound, "Not found").Render(page)
	return c.SendString(page.String())
}

func PageError(code int, message string) Node {
	return components.View("Home",
		components.Navbar(),
		Div(Class("content"),
			P(Text("Error "+strconv.Itoa(code))),
			P(Text(message)),
		),
	)
}
