package components

import (
	"messaging/util"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func View(title string, elements ...Node) Node {
	return HTML5(HTML5Props{
		Title:    util.APP_NAME + " ~ " + title,
		Language: "en",
		Head: []Node{
			Link(Rel("icon"), Type("image/x-icon"), Href("/static/favicon.png")),
			Link(Rel("stylesheet"), Href("static/styles/global.css")),
		},
		Body: []Node{
			Div(Class("container"),
				Group(elements),
			),
		},
	})
}
