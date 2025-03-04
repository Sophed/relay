package views

import (
	"messaging/app/components"
	"messaging/util"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// ViewLogin handles the web request to /login
func ViewLogin(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(util.Render(pageLogin()))
}


// pageLogin returns the login page content
func pageLogin() Node {
	return components.View("Login",
		components.Navbar(),
		Link(Rel("stylesheet"), Href("static/styles/login.css")),
		Script(Src("static/js/login.js")),
		Div(Class("content"),
			Div(Class("controls"),
				P(Text("Login to continue.")),
				// email input field
				Input(ID("inp-email"), Class("field"), Placeholder("Email address..."), MaxLength("64")),
				// password input field
				Input(ID("inp-password"), Class("field"), Type("password"), Placeholder("Password..."), MaxLength("64")),
				// submit button
				Button(Text("Login"), Attr("onclick", "submit()")),
				Button(Text("Register instead"), Attr("onclick", "register()")),
				A(Class("forgot"), Text("Forgot your password?")),
			),
		),
	)
}
