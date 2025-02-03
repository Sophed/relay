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
	SessionID string `json:"session_id"`
}

func Login(c *fiber.Ctx) error {
	var req loginReq
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		lg.Warn(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if req.Email == "" || req.Password == "" {
		return c.SendString(respErr("One or more missing fields"))
	}
	user, err := storage.METHOD.FindUser(&data.SearchableUser{
		Email: req.Email,
	})
	if err == storage.ErrNotFound {
		return c.SendString(respErr("No user found with this email"))
	}
	if err != nil {
		lg.Warn(err)
		return c.SendString(respInternal())
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return c.SendString(respErr("Incorrect login credentials"))
	}
	r, _ := json.MarshalIndent(res{
		SessionID: user.NewSession(),
	}, "", "\n")
	lg.Info("[login] " + user.Username + " - " + user.ID)
	return c.SendString(string(r))
}
