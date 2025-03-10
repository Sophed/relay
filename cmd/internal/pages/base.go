package pages

import (
	"messaging/util"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

// base returns the base HTML page template
func base(title string, elements ...Node) Node {
	return HTML5(HTML5Props{
		Title:    util.APP_NAME + " | " + title,
		Language: "en",
		Head: []Node{
			Link(Rel("icon"), Type("image/x-icon"), Href("/static/favicon.png")),
			Link(Rel("stylesheet"), Href("/static/css/global.css")),
			Script(Src("/static/lib/htmx.min.js"), Defer()),
		},
		Body: []Node{
			Div(Class("container"),
				Group(elements),
			),
		},
	})
}
