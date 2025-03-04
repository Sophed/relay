package views

import (
	"messaging/app/components"
	"messaging/util"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// ViewIndex handles the web request to /
func ViewIndex(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(util.Render(pageIndex()))
}

// pageIndex returns the home page content
func pageIndex() Node {
	return components.View("Home",
		components.Navbar(),
		Div(Class("content"),
			P(Text("Welcome to "+util.APP_NAME)),
		),
	)
}
