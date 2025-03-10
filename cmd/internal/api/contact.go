package api

import (
	"messaging/auth"
	"messaging/cmd/internal/components"
	"messaging/data"
	"messaging/data/storage"
	"messaging/web"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

// AddContact will either return an updated contacts component or trigger an error modal
func AddContact(c *fiber.Ctx) error {
	// check auth
	authed, user := auth.Validate(c)
	if !authed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// get target user
	targetUsername := c.FormValue("target", "")
	target, err := storage.METHOD.FindUser(&data.SearchableUser{
		Username: targetUsername,
	})
	if err != nil { // error modals on fail
		if err == storage.ErrNotFound {
			web.SetModalError(c, "No user found with this name.")
		} else {
			web.SetModalError(c, "Something went wrong, please try again later.")
			lg.Warn(err) // probably bad
		}
	}

	target.GetTokens()
	// TODO: add target to users contact list - we need one of those

	return c.SendString(web.Render(components.ContactsSidebar(user)))
}
