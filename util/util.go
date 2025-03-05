package util

import (
	"unicode"
)

const APP_NAME = "Relay"

// AlphaNumeric ensures a string contains only letter or number values
func AlphaNumeric(s string) bool {
	for _, c := range s { // loop over every character
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			continue
		} else {
			return false
		}
	}
	return true
}
