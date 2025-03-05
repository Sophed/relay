package api

import (
	"messaging/auth"
	"messaging/data/entities"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type registerReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// handles web requests to /api/register
func Register(c *fiber.Ctx) error {
	username := c.FormValue("username", "")
	email := c.FormValue("email", "")
	password := c.FormValue("password", "")

	// enforce lowercase usernames
	token, err := auth.Register(strings.ToLower(username), email, password)
	if err != nil {
		return c.SendString(err.Error())
	}

	// set response
	c.Set("HX-Redirect", "/app")
	c.Cookie(entities.SessionCookie(token.Value))
	return c.SendString("")
}
