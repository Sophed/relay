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

// handles web requests to /api/login
func Login(c *fiber.Ctx) error {
	var req loginReq
	err := json.Unmarshal(c.Body(), &req) // parse JSON as loginReq type
	if err != nil {
		lg.Warn(err) // log on fail
		return c.SendStatus(fiber.StatusBadRequest) // return error
	}
	// check if fields are empty
	if req.Email == "" || req.Password == "" {
		return c.SendString(respErr("One or more missing fields"))
	}
	// find user by email
	user, err := storage.METHOD.FindUser(&data.SearchableUser{
		Email: req.Email,
	})
	if err == storage.ErrNotFound {
		return c.SendString(respErr("No user found with this email"))
	}
	if err != nil {
		lg.Warn(err) // log on fail
		return c.SendString(respInternal()) // return error
	}
	// check if request password matches user password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return c.SendString(respErr("Incorrect login credentials"))
	}
	// create response and convert to JSON
	r, _ := json.MarshalIndent(res{
		SessionID: user.NewSession(),
	}, "", "\n")
	lg.Info("[login] " + user.Username + " - " + user.ID) // log user login
	return c.SendString(string(r)) // return JSON
}
