package views

import (
	"messaging/app/components"
	"messaging/util"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// ViewApp handles the web request to /app
func ViewApp(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(util.Render(pageApp()))
}

// pageApp returns the app page content
func pageApp() Node {
	return components.View("App",
		Link(Rel("stylesheet"), Href("static/styles/app.css")),
		Div(Class("panes"),
			components.ContactsSidebar(),
			components.ChatWindow(),
			components.ProfileSidebar(),
		),
	)
}
