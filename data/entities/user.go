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

func (u *User) NewSession() string {
	return sessions.New(u.ID)
}

func (u *User) GetTokens() []string {
	return sessions.GetTokens(u.ID)
}

func Hash(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		lg.Fatl(err)
	}
	return string(hash)
}
