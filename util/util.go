package util

import (
	"bytes"
	"unicode"

	"github.com/sophed/lg"
	"maragu.dev/gomponents"
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

// Render converts a HTML node to a usable string which can be sent to clients
func Render(page gomponents.Node) string {
	buf := new(bytes.Buffer) // empty buffer
	err := page.Render(buf) // render to buffer
	if err != nil {
		lg.Fatl(err) // panic on fail
	}
	return buf.String() // return buffer content as string
}
