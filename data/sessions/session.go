package sessions

import "github.com/google/uuid"

const SESSION_COOKIE_KEY = "relay-session"

var sessionMap = make(map[string][]Token)

type Token struct {
	Value string
}

// New creates a new session for a given user ID and returns the generated token
func New(id string) Token {
	token := uuid.NewString()
	sessionMap[id] = append(sessionMap[id], Token{token}) // append to session map
	return Token{token}
}

// ValidToken evaluates whether a given token is valid and returns the user ID if true
func ValidToken(token Token) (bool, string) {
	for id, list := range sessionMap {
		for _, t := range list {
			if t.Value == token.Value {
				// return true & ID if token exists in the map
				return true, id
			}
		}
	}
	// assume false if no results
	return false, ""
}

// GetTokens returns a list of valid tokens for a given user ID
func GetTokens(id string) []Token {
	return sessionMap[id]
}
