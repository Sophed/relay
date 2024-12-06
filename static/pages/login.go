package pages

import (
	"messaging/static"
	"messaging/util"
	"os"
	"strings"

	"github.com/sophed/lg"
)

func LoginPage() (string, error) {
	data, err := os.ReadFile(static.PAGES_DIR + "login.html")
	if err != nil {
		lg.Warn(err)
		return "", err
	}
	content := strings.ReplaceAll(string(data), util.Template("app"), static.APP_NAME)
	styles, err := static.GlobalStyles()
	if err != nil {
		lg.Warn(err)
		return "", err
	}
	nav, err := static.Nav()
	if err != nil {
		lg.Warn(err)
		return "", err
	}
	specifics, err := os.ReadFile(static.STYLES_DIR + "login.css")
	if err != nil {
		lg.Warn(err)
		return "", err
	}
	styles += string(specifics)
	content = strings.ReplaceAll(content, util.Template("nav"), nav)
	content = strings.ReplaceAll(content, util.Template("styles"), styles)
	return content, nil
}
