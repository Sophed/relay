package api

import (
	"messaging/data/sessions"
	"time"

	"github.com/gofiber/fiber/v2"
)

// handles web requests to /api/logout
func Logout(c *fiber.Ctx) error {
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
