package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ChatWindow() Node {
	return Div(Class("pane chat-window"),
		Text("chat window"),
	)
}
