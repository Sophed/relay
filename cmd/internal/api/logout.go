package api

import (
	"messaging/auth"
	"messaging/data/sessions"
	"time"

	"github.com/sophed/lg"
	"github.com/gofiber/fiber/v2"
)

// handles web requests to /api/logout
func Logout(c *fiber.Ctx) error {
	// check auth
	authed, user := auth.Validate(c)
	if authed {
		lg.Info("[logout] " + user.Username + " - " + user.ID) // log user login
	}
	c.Set("HX-Redirect", "/") // redirect to /
	c.Cookie(&fiber.Cookie{   // send an empty cookie to override existing sessions
		Name:     sessions.SESSION_COOKIE_KEY,
		Value:    "",
		Secure:   true,
		HTTPOnly: true,
		Expires:  time.Now().AddDate(0, 0, -1),
	})
	return c.SendStatus(fiber.StatusOK)
}
