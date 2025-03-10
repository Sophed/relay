package components

import (
	"messaging/data/entities"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ContactsSidebar(*entities.User) Node {
	return Div(Class("pane contacts-sidebar"),
		Div(Class("top"),
			Div(Class("controls"),
				Input(Placeholder("Search by username...")),
				Button(Text("+")),
			),
		),
	)
}
