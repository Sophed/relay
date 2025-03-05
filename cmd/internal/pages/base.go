package pages

import (
	"bytes"
	"messaging/util"

	"github.com/sophed/lg"
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

// render converts a HTML node to a usable string which can be sent to clients
func render(page Node) string {
	buf := new(bytes.Buffer) // empty buffer
	err := page.Render(buf)  // render to buffer
	if err != nil {
		lg.Fatl(err) // panic on fail
	}
	return buf.String() // return buffer content as string
}
