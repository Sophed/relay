package web

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func modalTrigger(event, msg string) string {
	m := map[string]string{
		event: msg,
	}
	data, _ := json.Marshal(m)
	return string(data)
}

func SetModalError(c *fiber.Ctx, msg string) {
	c.Set("HX-Trigger", modalTrigger("modalError", msg))
}
