package api

import (
	"encoding/json"
	"messaging/data"
	"messaging/data/entities"
	"messaging/data/storage"
	"messaging/util"
	"net/mail"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sophed/lg"
)

type registerReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var req registerReq
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		lg.Warn(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return c.SendString(respErr("One or more fields are empty"))
	}
	if !validEmail(req.Email) || !validUsername(req.Username) || !validPassword(req.Password) {
		return c.SendString(respErr("One or more fields have an invalid format"))
	}
	req.Username = strings.ToLower(req.Username)
	exists, err := alreadyExists(&data.SearchableUser{
		Email:    req.Email,
		Username: req.Username,
	})
	if err != nil {
		lg.Warn(err)
		return c.SendString(respInternal())
	}
	if exists {
		return c.SendString(respErr("A user with this email or username already exists"))
	}
	id, err := genID()
	if err != nil {
		lg.Warn(err)
		return c.SendString(respInternal())
	}
	user := entities.User{
		ID:             id,
		Username:       req.Username,
		DisplayName:    req.Username,
		Email:          req.Email,
		ProfilePicture: "",
		PasswordHash:   entities.Hash([]byte(req.Password)),
		Timestamp:      time.Now().UnixMilli(),
	}
	err = storage.METHOD.AddUser(&user)
	if err != nil {
		lg.Warn(err)
		return c.SendString(respInternal())
	}
	r, _ := json.MarshalIndent(res{
		SessionID: "valid",
	}, "", "\n")
	lg.Info("[register] " + user.Username + " - " + user.ID)
	return c.SendString(string(r))
}

func validUsername(username string) bool {
	if len(username) > 32 {
		return false
	}
	return util.AlphaNumeric(username)
}

func validEmail(email string) bool {
	if len(email) > 64 {
		return false
	}
	_, err := mail.ParseAddress(email) // TODO: replace this
	return err == nil
}

func validPassword(password string) bool {
	if len(password) > 64 {
		return false
	}
	return true
}

func alreadyExists(query *data.SearchableUser) (bool, error) {
	_, err := storage.METHOD.FindUser(query)
	if err != nil {
		if err == storage.ErrNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func genID() (string, error) {
	id := uuid.New().String()
	exists, err := alreadyExists(&data.SearchableUser{
		ID: id,
	})
	if err != nil {
		return "", err
	}
	if exists {
		return genID()
	}
	return id, nil
}
