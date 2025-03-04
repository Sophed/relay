package views

import (
	"messaging/app/components"
	"messaging/util"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// ViewRegister handles the web request to /register
func ViewRegister(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(util.Render(pageRegister()))
}

// pageRegister returns the register page content
func pageRegister() Node {
	return components.View("Register",
		components.Navbar(),
		Link(Rel("stylesheet"), Href("static/styles/login.css")),
		Script(Src("static/js/register.js")),
		Div(Class("content"),
			Div(Class("controls"),
				P(Text("Create an account to continue.")),
				// username input field
				Input(ID("inp-username"), Class("field"), Placeholder("Username..."), MaxLength("64")),
				// email input field
				Input(ID("inp-email"), Class("field"), Placeholder("Email address..."), MaxLength("64")),
				// password input field
				Input(ID("inp-password"), Class("field"), Type("password"), Placeholder("Password..."), MaxLength("64")),
				// submit button
				Button(Text("Register"), Attr("onclick", "submit()")),
				Button(Text("Login instead"), Attr("onclick", "login()")),
			),
		),
	)
}
