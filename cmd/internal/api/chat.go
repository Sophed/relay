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

// GetChat returns either specified chat component or triggers an error modal
func GetChat(c *fiber.Ctx) error {
	// check auth
	authed, user := auth.Validate(c)
	if !authed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// get target user
	targetID := c.Params("id", "")
	target, err := storage.METHOD.FindUser(&data.SearchableUser{
		ID: targetID,
	})
	if err != nil { // error modals on fail
		if err == storage.ErrNotFound {
			web.SetModalError(c, "No user found with this ID.")
		} else {
			web.SetModalError(c, "Something went wrong, please try again later.")
			lg.Erro(err) // probably bad
		}
		c.SendStatus(fiber.StatusOK)
	}

	return c.SendString(web.Render(components.ChatWindow(user, target)))
}
