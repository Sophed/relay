package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ContactsSidebar() Node {
	return Div(Class("pane contacts-sidebar"),
		Div(Class("top"),
			Div(Class("controls"),
				Input(Placeholder("Search by username...")),
				Button(Text("+")),
			),
		),
	)
}
