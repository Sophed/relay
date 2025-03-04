package entities

import (
	"messaging/data/sessions"

	"github.com/sophed/lg"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string   `json:"id" bson:"_id"`
	Username       string   `json:"username" bson:"username"`
	DisplayName    string   `json:"display_name" bson:"display_name"`
	Email          string   `json:"email" bson:"email"`
	ProfilePicture string   `json:"pfp" bson:"pfp"`
	PasswordHash   string   `json:"password_hash" bson:"password_hash"`
	Timestamp      int64    `json:"timestamp" bson:"timestamp"`
	Contacts       []string `json:"contacts"`
}

// NewSession creates a new session for the given user and returns the token
func (u *User) NewSession() string {
	return sessions.New(u.ID)
}

// GetTokens returns all of the stored tokens for the given user
func (u *User) GetTokens() []string {
	return sessions.GetTokens(u.ID)
}

// Hash takes a password as input and returns it in a hashed format for storage
func Hash(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		lg.Fatl(err)
	}
	return string(hash)
}
