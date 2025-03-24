package components

import (
	"messaging/data/entities"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

const ID_CHAT_WINDOW = "chat-window"

func ChatWindow(user, target *entities.User) Node {
	if user == nil || target == nil {
		return Div(ID(ID_CHAT_WINDOW), Class("pane"),
			Text("Select a contact to get started"),
		)
	}
	return Div(ID(ID_CHAT_WINDOW), Class("pane"),
		Text("chat between "+user.DisplayName+" and "+target.DisplayName),
		Form(ID("message-compose"),
			Input(ID("chat-input"), Placeholder("Write a messsage to @"+target.Username)),
			Button(ID("chat-send"), Text("Send")),
		),
	)
}
