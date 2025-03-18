package components

import (
	"messaging/data"
	"messaging/data/entities"
	"messaging/data/storage"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func ContactsSidebar(user *entities.User) Node {
	return Div(Class("pane contacts-sidebar"),
		Div(Class("top"),
			Div(Class("controls"),
				Form(hx.Put("/api/contacts"), hx.Target("#contacts-list"), hx.Swap("outerHTML"),
					Input(Name("target"), Placeholder("Search by username...")),
					Button(Type("submit"), Text("+")),
				),
			),
			ContactsList(user),
		),
	)
}

func ContactsList(user *entities.User) Node {
	contacts := fetchContacts(user.Contacts)
	return Ul(ID("contacts-list"),
		Map(contacts, func(u entities.User) Node {
			return contactCard(&u)
		}),
	)
}

func contactCard(user *entities.User) Node {
	return Div(Class("contact-card"),
		Img(Src(user.ProfilePicture), Class("pfp")),
		P(Class("label"), Text(user.DisplayName)),
		P(Class("sublabel"), Text(user.Username)),
	)
}

func fetchContacts(userIDs []string) []entities.User {
	list := []entities.User{}
	for _, s := range userIDs {
		user, err := storage.METHOD.FindUser(&data.SearchableUser{
			ID: s,
		})
		if err == nil {
			list = append(list, *user)
		}
	}
	return list
}
