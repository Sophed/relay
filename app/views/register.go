package views

import (
	"messaging/app/components"
	"messaging/util"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ViewRegister(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(util.Render(pageRegister()))
}

func pageRegister() Node {
	return components.View("Register",
		components.Navbar(),
		Link(Rel("stylesheet"), Href("static/styles/login.css")),
		Script(Src("static/js/register.js")),
		Div(Class("content"),
			Div(Class("controls"),
				P(Text("Create an account to continue.")),
				Input(ID("inp-username"), Class("field"), Placeholder("Username..."), MaxLength("64")),
				Input(ID("inp-email"), Class("field"), Placeholder("Email address..."), MaxLength("64")),
				Input(ID("inp-password"), Class("field"), Type("password"), Placeholder("Password..."), MaxLength("64")),
				Button(Text("Register"), Attr("onclick", "submit()")),
				Button(Text("Login instead"), Attr("onclick", "login()")),
			),
		),
	)
}
