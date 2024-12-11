package api

import (
	"encoding/json"
	"messaging/data"
	"messaging/data/entities"
	"messaging/data/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req loginReq
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		lg.Warn(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if req.Email == "" || req.Password == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	password := entities.Hash([]byte(req.Password))
	user, err := storage.METHOD.FindUser(&data.SearchableUser{
		Email: req.Email,
	})
	if err == storage.ErrNotFound {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if err != nil {
		lg.Warn(err)
	}
	if user.PasswordHash != password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.SendStatus(fiber.StatusOK)
}
