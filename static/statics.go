package static

import (
	"os"
)

const APP_NAME = "Relay"
const STATIC_DIR = "static/"
const PAGES_DIR = STATIC_DIR + "pages/"

func GlobalStyles() (string, error) {
	data, err := os.ReadFile(STATIC_DIR + "global.css")
	return string(data), err
}

func Nav() (string, error) {
	styles, err := os.ReadFile(STATIC_DIR + "nav.css")
	if err != nil {
		return "", err
	}
	content, err := os.ReadFile(STATIC_DIR + "nav.html")
	if err != nil {
		return "", err
	}
	return "<style>" + string(styles) + "</style>\n" + string(content), err
}
