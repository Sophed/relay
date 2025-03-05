package auth

import (
	"errors"
	"messaging/data"
	"messaging/data/entities"
	"messaging/data/sessions"
	"messaging/data/storage"
	"messaging/util"
	"net/mail"
	"time"

	"github.com/google/uuid"
	"github.com/sophed/lg"
)

var (
	ErrBadUsername = errors.New("Invalid username format, must only contain alphanumeric characters and '.'")
	ErrBadEmail    = errors.New("Invalid email format")
	ErrBadPassword = errors.New("Invalid password")
	ErrConflict    = errors.New("A user with this email or username already exists")
)

func Register(username, email, password string) (*sessions.Token, error) {
	// check if fields are empty
	if username == "" || email == "" || password == "" {
		return nil, ErrEmptyFields
	}

	// check if fields pass validation checks
	if !validUsername(username) {
		return nil, ErrBadUsername
	}
	if !validEmail(email) {
		return nil, ErrBadEmail
	}
	if !validPassword(password) {
		return nil, ErrBadPassword
	}

	// check if user already exists
	exists, err := alreadyExists(&data.SearchableUser{
		Email:    email,
		Username: username,
	})
	if err != nil {
		return nil, ErrInternal
	}
	if exists {
		return nil, ErrConflict
	}

	// generate user ID
	id, err := genID()
	if err != nil {
		return nil, ErrInternal
	}

	// create user object with request details
	user := entities.User{
		ID:             id,
		Username:       username,
		DisplayName:    username,
		Email:          email,
		ProfilePicture: "",
		PasswordHash:   entities.Hash([]byte(password)), // hashed password
		Timestamp:      time.Now().UnixMilli(),          // current unix timestamp
	}

	// store user
	err = storage.METHOD.AddUser(&user)
	if err != nil {
		return nil, ErrInternal
	}

	token := sessions.New(id)
	lg.Info("[register] " + user.Username + " - " + user.ID) // log user registration
	return &token, nil
}

// alreadyExists checks if a user exists in storage
func alreadyExists(query *data.SearchableUser) (bool, error) {
	_, err := storage.METHOD.FindUser(query)
	if err != nil {
		if err == storage.ErrNotFound {
			return false, nil // user does not exist
		}
		return false, err // other error
	}
	return true, nil // user exists
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
