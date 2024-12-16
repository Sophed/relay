package util

import "unicode"

const APP_NAME = "Relay"

func AlphaNumeric(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			continue
		} else {
			return false
		}
	}
	return true
}
