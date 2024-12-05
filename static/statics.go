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
