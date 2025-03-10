package pages

import (
	"messaging/auth"
	"messaging/cmd/internal/components"
	"messaging/data/entities"
	"messaging/web"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// App handles the web request to /app
func App(c *fiber.Ctx) error {
	authed, user := auth.Validate(c)
	if !authed {
		return c.Redirect("/") // redirect to home if not logged in
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(web.Render(pageApp(user)))
}

// pageApp returns the app page content
func pageApp(user *entities.User) Node {
	return base("App",
		Link(Rel("stylesheet"), Href("/static/css/app.css")),
		Div(Class("panes"),
			components.ContactsSidebar(user),
			components.ChatWindow(),
			components.ProfileSidebar(),
		),
	)
}
