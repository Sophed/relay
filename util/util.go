package util

import (
	"bytes"
	"unicode"

	"github.com/sophed/lg"
	"maragu.dev/gomponents"
)

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

func Render(page gomponents.Node) string {
	buf := new(bytes.Buffer)
	err := page.Render(buf)
	if err != nil {
		lg.Fatl(err)
	}
	return buf.String()
}
