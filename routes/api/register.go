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

// handles web requests to /api/register
func Register(c *fiber.Ctx) error {
	var req registerReq
	err := json.Unmarshal(c.Body(), &req) // parse JSON as registerReq type
	if err != nil {
		lg.Warn(err) // log on fail
		return c.SendStatus(fiber.StatusBadRequest) // return error
	}
	// check if fields are empty
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return c.SendString(respErr("One or more fields are empty"))
	}
	// check if fields pass validation checks
	if !validEmail(req.Email) || !validUsername(req.Username) || !validPassword(req.Password) {
		return c.SendString(respErr("One or more fields have an invalid format"))
	}
	req.Username = strings.ToLower(req.Username) // enforce lowercase usernames
	// check if username or email already exists
	exists, err := alreadyExists(&data.SearchableUser{
		Email:    req.Email,
		Username: req.Username,
	})
	if err != nil {
		lg.Warn(err) // log on fail
		return c.SendString(respInternal()) // return error
	}
	if exists {
		return c.SendString(respErr("A user with this email or username already exists"))
	}
	id, err := genID() // generate user ID
	if err != nil {
		lg.Warn(err) // log on fail
		return c.SendString(respInternal()) // return error
	}
	// create user object with request details
	user := entities.User{
		ID:             id,
		Username:       req.Username,
		DisplayName:    req.Username,
		Email:          req.Email,
		ProfilePicture: "",
		PasswordHash:   entities.Hash([]byte(req.Password)), // hashed password
		Timestamp:      time.Now().UnixMilli(), // current unix timestamp
	}
	// store user
	err = storage.METHOD.AddUser(&user)
	if err != nil {
		lg.Warn(err) // log on fail
		return c.SendString(respInternal()) // return error
	}
	// create response and convert to JSON
	r, _ := json.MarshalIndent(res{
		SessionID: "valid",
	}, "", "\n")
	lg.Info("[register] " + user.Username + " - " + user.ID) // log user registration
	return c.SendString(string(r)) // return JSON
}

// validUsername checks if a username matches validation
func validUsername(username string) bool {
	if len(username) > 32 {
		return false
	}
	return util.AlphaNumeric(username)
}

// validEmail checks if an email matches validation
func validEmail(email string) bool {
	if len(email) > 64 {
		return false
	}
	_, err := mail.ParseAddress(email) // TODO: replace this
	return err == nil
}

// validPassword checks if a password matches validation
func validPassword(password string) bool {
	if len(password) > 64 {
		return false
	}
	return true
}

// alreadyExists checks if a user exists in storage
func alreadyExists(query *data.SearchableUser) (bool, error) {
	_, err := storage.METHOD.FindUser(query) // find user with details
	if err != nil {
		if err == storage.ErrNotFound {
			return false, nil // return other error
		} else {
			return false, err // return not found
		}
	}
	return true, nil // user exists
}

// genID generates a new user ID ensuring it does not already exist
func genID() (string, error) {
	id := uuid.New().String() // generate ID
	// search for user by ID
	exists, err := alreadyExists(&data.SearchableUser{
		ID: id,
	})
	if err != nil {
		return "", err // return error on fail
	}
	if exists {
		return genID() // recursion - attempt again if the generated ID exists
	}
	return id, nil // return ID, no user exists already
}
