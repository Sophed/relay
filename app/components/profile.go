package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ProfileSidebar() Node {
	return Div(Class("pane profile-sidebar"),
		Text("profile sidebar"),
	)
}
