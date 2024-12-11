package views

import (
	"bytes"
	"messaging/app/components"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ViewLogin(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	page := new(bytes.Buffer)
	PageLogin().Render(page)
	return c.SendString(page.String())
}

func PageLogin() Node {
	return components.View("Login",
		components.Navbar(),
		Link(Rel("stylesheet"), Href("static/styles/login.css")),
		Script(Src("static/js/login.js")),
		Div(Class("content"),
			Div(Class("controls"),
				P(Text("Login")),
				Input(Class("field"), Placeholder("Email address...")),
				Input(Class("field"), Type("password"), Placeholder("Password...")),
				Raw("<button onclick=\"submit()\">Login</button>"),
				Button(Text("Register")),
				A(Class("forgot"), Text("Forgot your password?")),
			),
		),
	)
}
