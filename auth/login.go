package auth

import (
	"errors"
	"messaging/data"
	"messaging/data/entities"
	"messaging/data/sessions"
	"messaging/data/storage"

	"github.com/sophed/lg"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmptyFields = errors.New("One or more fields are empty")
	ErrNoUser      = errors.New("No user found with this email address")
	ErrBadLogin    = errors.New("Incorrect login credentials")
	ErrInternal    = errors.New("An internal error occurred, please try again later")
)

func Login(email, password string) (*sessions.Token, error) {
	// check if fields are empty
	if email == "" || password == "" {
		return nil, ErrEmptyFields
	}
	user, err := findUser(email)
	if err != nil {
		if err != ErrNoUser {
			return nil, ErrInternal
		}
		return nil, err
	}
	// check if request password matches user password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, ErrBadLogin
	}
	token := user.NewSession()
	lg.Info("[login] " + user.Username + " - " + user.ID) // log user login
	return &token, nil
}

func findUser(email string) (*entities.User, error) {
	// find user by email
	user, err := storage.METHOD.FindUser(&data.SearchableUser{
		Email: email,
	})
	if err == storage.ErrNotFound {
		return nil, ErrNoUser
	}
	return user, err
}
