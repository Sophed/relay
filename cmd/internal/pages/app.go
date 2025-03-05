package pages

import (
	"messaging/cmd/internal/components"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// App handles the web request to /app
func App(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(render(pageApp()))
}

// pageApp returns the app page content
func pageApp() Node {
	return base("App",
		Link(Rel("stylesheet"), Href("/static/css/app.css")),
		Div(Class("panes"),
			components.ContactsSidebar(),
			components.ChatWindow(),
			components.ProfileSidebar(),
		),
	)
}
