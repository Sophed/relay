package views

import (
	"messaging/app/components"
	"messaging/util"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// ViewPricing handles the web request to /pricing
func ViewPricing(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(util.Render(pagePricing()))
}

// pagePricing returns the pricing page content
func pagePricing() Node {
	return components.View("Pricing",
		components.Navbar(),
		Div(Class("content"),
			P(Text("Welcome to "+util.APP_NAME)),
		),
	)
}
