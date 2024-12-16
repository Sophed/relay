package views

import (
	"bytes"
	"messaging/app/components"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ViewRegister(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	page := new(bytes.Buffer)
	PageRegister().Render(page)
	return c.SendString(page.String())
}

func PageRegister() Node {
	return components.View("Register",
		components.Navbar(),
		Link(Rel("stylesheet"), Href("static/styles/login.css")),
		Script(Src("static/js/register.js")),
		Div(Class("content"),
			Div(Class("controls"),
				P(Text("Create an account to continue.")),
				Input(ID("inp-username"), Class("field"), Placeholder("Username...")),
				Input(ID("inp-email"), Class("field"), Placeholder("Email address...")),
				Input(ID("inp-password"), Class("field"), Type("password"), Placeholder("Password...")),
				Button(Text("Register"), Attr("onclick", "submit()")),
				Button(Text("Login instead"), Attr("onclick", "login()")),
			),
		),
	)
}
