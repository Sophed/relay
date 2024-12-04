package static

import (
	"messaging/util"
	"os"
	"strconv"
	"strings"
)

const APP_NAME = "Relay"
const STATIC_DIR = "static/"
const PAGES_DIR = STATIC_DIR + "pages/"

func GlobalStyles() (string, error) {
	data, err := os.ReadFile(STATIC_DIR + "global.css")
	return string(data), err
}

func ErrorPage(code int, message string) string {
	data, err := os.ReadFile(PAGES_DIR + "error.html")
	if err != nil {
		return util.SmallError(code, message)
	}
	content := strings.ReplaceAll(string(data), util.Template("app"), APP_NAME)
	styles, err := GlobalStyles()
	if err != nil {
		return util.SmallError(code, message)
	}
	content = strings.ReplaceAll(content, util.Template("styles"), styles)
	content = strings.ReplaceAll(content, util.Template("code"), strconv.Itoa(code))
	return strings.ReplaceAll(content, util.Template("message"), message)
}

func LoginPage() (string, error) {
	data, err := os.ReadFile(PAGES_DIR + "login.html")
	if err != nil {
		return "", err
	}
	content := strings.ReplaceAll(string(data), util.Template("app"), APP_NAME)
	styles, err := GlobalStyles()
	if err != nil {
		return "", err
	}
	content = strings.ReplaceAll(content, util.Template("styles"), styles)
	return content, nil
}
