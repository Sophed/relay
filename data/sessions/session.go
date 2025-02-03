package sessions

import "github.com/google/uuid"

var sessionMap = make(map[string][]string)

func New(id string) string {
	token := uuid.NewString()
	sessionMap[id] = append(sessionMap[id], token)
	return token
}

func ValidToken(token string) bool {
	for _, list := range sessionMap {
		for _, t := range list {
			if t == token {
				return true
			}
		}
	}
	return false
}

func GetTokens(id string) []string {
	return sessionMap[id]
}
