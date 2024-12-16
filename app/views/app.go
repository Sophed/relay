package views

import (
	"bytes"
	"messaging/app/components"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ViewApp(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	page := new(bytes.Buffer)
	PageApp().Render(page)
	return c.SendString(page.String())
}

func PageApp() Node {
	return components.View("App",
		P(Text("app!!!")),
	)
}
