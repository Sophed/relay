package sessions

import "github.com/google/uuid"

var sessionMap = make(map[string][]string)

// New creates a new session for a given user ID and returns the generated token
func New(id string) string {
	token := uuid.NewString()
	sessionMap[id] = append(sessionMap[id], token) // append to session map
	return token
}

// ValidToken evaluates whether a given token is valid
func ValidToken(token string) bool {
	for _, list := range sessionMap {
		for _, t := range list {
			if t == token {
				// return true if a token exists anywhere in the map
				return true
			}
		}
	}
	// assume false if no results
	return false
}

// GetTokens returns a list of valid tokens for a given user ID
func GetTokens(id string) []string {
	return sessionMap[id]
}
