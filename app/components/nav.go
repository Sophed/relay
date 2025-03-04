package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// Navbar returns navbar UI component
func Navbar() Node {
	return Div(
		Link(Rel("stylesheet"), Href("static/styles/nav.css")),
		Ul(Class("nav"),
			navItem("Home", "/", "left"),
			navItem("Pricing", "/pricing", "left"),
			navItem("Sign Up", "/register", "right"),
			navItem("Login", "/login", "right"),
		),
	)
}

// navItem returns navigation item component
func navItem(name, path string, pos string) Node {
	return Li(Style("float:"+pos),
		A(Href(path), Text(name)),
	)
}
