package api

import (
	"messaging/auth"
	"messaging/data/entities"

	"github.com/gofiber/fiber/v2"
)

// handles web requests to /api/login
func Login(c *fiber.Ctx) error {
	email := c.FormValue("email", "")
	password := c.FormValue("password", "")

	token, err := auth.Login(email, password)
	if err != nil {
		return c.SendString(err.Error())
	}

	c.Set("HX-Redirect", "/app")
	c.Cookie(entities.SessionCookie(token.Value))
	return c.SendString("")
}
