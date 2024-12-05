package views

import (
	"messaging/static/pages"

	"github.com/gofiber/fiber/v2"
)

func PageErrorInternal(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(pages.ErrorPage(fiber.StatusInternalServerError, "Internal Server Error"))
}
