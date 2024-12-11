package views

import (
	"bytes"
	"messaging/app/components"
	"messaging/util"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ViewIndex(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	page := new(bytes.Buffer)
	PageIndex().Render(page)
	return c.SendString(page.String())
}

func PageIndex() Node {
	return components.View("Home",
		components.Navbar(),
		Div(Class("content"),
			P(Text("Welcome to "+util.APP_NAME)),
		),
	)
}
