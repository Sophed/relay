package pages

import (
	"messaging/static"
	"messaging/util"
	"os"
	"strings"
)

func LoginPage() (string, error) {
	data, err := os.ReadFile(static.PAGES_DIR + "login.html")
	if err != nil {
		return "", err
	}
	content := strings.ReplaceAll(string(data), util.Template("app"), static.APP_NAME)
	styles, err := static.GlobalStyles()
	if err != nil {
		return "", err
	}
	content = strings.ReplaceAll(content, util.Template("styles"), styles)
	return content, nil
}
