package auth

import (
	"messaging/data"
	"messaging/data/entities"
	"messaging/data/sessions"
	"messaging/data/storage"

	"github.com/gofiber/fiber/v2"
)

// Validate returns the auth state of a request along with the user associated
func Validate(c *fiber.Ctx) (bool, *entities.User) {
	// get ID from request cookie
	state, id := sessions.ValidToken(sessions.Token{
		Value: c.Cookies(sessions.SESSION_COOKIE_KEY, ""),
	})
	if !state {
		return false, nil
	}

	// get user using ID
	user, err := storage.METHOD.FindUser(&data.SearchableUser{
		ID: id,
	})
	if err != nil {
		return false, nil
	}
	return true, user
}
