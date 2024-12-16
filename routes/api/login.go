package api

import (
	"encoding/json"
	"messaging/data"
	"messaging/data/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
	"golang.org/x/crypto/bcrypt"
)

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type res struct {
	Message string `json:"message"`
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
	user, err := storage.METHOD.FindUser(&data.SearchableUser{
		Email: req.Email,
	})
	if err == storage.ErrNotFound {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if err != nil {
		lg.Warn(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	r, err := json.MarshalIndent(res{
		Message: "valid",
	}, "", "\n")
	if err != nil {
		lg.Warn(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	lg.Info("[login] " + user.Username + " - " + user.ID)
	return c.SendString(string(r))
}
