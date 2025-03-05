package pages

import (
	"messaging/cmd/internal/components"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

// Login handles the web request to /login
func Login(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(render(pageLogin()))
}

// pageLogin returns the login page content
func pageLogin() Node {
	return base("Login",
		components.Navbar(),
		Link(Rel("stylesheet"), Href("/static/css/login.css")),
		Div(Class("content"),
			Form(Class("controls"), hx.Post("/api/login"), hx.Target("#status"), hx.Swap("innerHTML"),
				P(Text("Login to continue.")),
				// email input field
				Input(Name("email"), ID("email"), Class("field"), Placeholder("Email address..."), MaxLength("64")),
				// password input field
				Input(Name("password"), ID("password"), Class("field"), Type("password"), Placeholder("Password..."), MaxLength("64")),
				// submit button
				Button(Type("submit"), Text("Login")),
				Button(Type("button"), Text("Register instead"), Attr("onclick", components.BtnLink("/register"))),
				A(ID("status"), Class("forgot"), Text("")),
				A(Class("forgot"), Text("Forgot your password?")),
			),
		),
	)
}
