package pages

import (
	"messaging/auth"
	"messaging/cmd/internal/components"
	"messaging/web"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

// Register handles the web request to /register
func Register(c *fiber.Ctx) error {
	authed, _ := auth.Validate(c)
	if !authed {
		return c.Redirect("/app") // redirect to /app if not logged in
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(web.Render(pageRegister()))
}

// pageRegister returns the register page content
func pageRegister() Node {
	return base("Register",
		components.Navbar(),
		Link(Rel("stylesheet"), Href("static/css/login.css")),
		Div(Class("content"),
			Form(Class("controls"), hx.Post("/api/register"), hx.Target("#status"), hx.Swap("innerHTML"),
				P(Text("Create an account to continue.")),
				// username input field
				Input(Name("username"), ID("username"), Class("field"), Placeholder("Username..."), MaxLength("64")),
				// email input field
				Input(Name("email"), ID("email"), Class("field"), Placeholder("Email address..."), MaxLength("64")),
				// password input field
				Input(Name("password"), ID("password"), Class("field"), Type("password"), Placeholder("Password..."), MaxLength("64")),
				// submit button
				Button(Type("submit"), Text("Register")),
				Button(Type("button"), Text("Login instead"), Attr("onclick", components.BtnLink("/login"))),
				A(ID("status"), Class("forgot"), Text("")),
			),
		),
	)
}
