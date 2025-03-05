package pages

import (
	"messaging/cmd/internal/components"
	"messaging/util"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// Index handles the web request to /
func Index(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(render(pageIndex()))
}

// pageIndex returns the home page content
func pageIndex() Node {
	return base("Home",
		components.Navbar(),
		Div(Class("content"),
			P(Text("Welcome to "+util.APP_NAME)),
		),
	)
}
