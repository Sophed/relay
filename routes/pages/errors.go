package pages

import (
	"messaging/static"

	"github.com/gofiber/fiber/v2"
)

func PageErrorInternal(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(static.ErrorPage(fiber.StatusInternalServerError, "Internal Server Error"))
}
