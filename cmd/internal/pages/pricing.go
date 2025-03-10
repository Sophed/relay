package pages

import (
	"messaging/cmd/internal/components"
	"messaging/util"
	"messaging/web"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// Pricing handles the web request to /pricing
func Pricing(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(web.Render(pagePricing()))
}

// pagePricing returns the pricing page content
func pagePricing() Node {
	return base("Pricing",
		components.Navbar(),
		Div(Class("content"),
			P(Text("Welcome to "+util.APP_NAME)),
		),
	)
}
