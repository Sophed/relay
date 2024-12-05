package pages

import (
	"messaging/static"
	"messaging/util"
	"os"
	"strconv"
	"strings"
)

func ErrorPage(code int, message string) string {
	data, err := os.ReadFile(static.PAGES_DIR + "error.html")
	if err != nil {
		return util.SmallError(code, message)
	}
	content := strings.ReplaceAll(string(data), util.Template("app"), static.APP_NAME)
	styles, err := static.GlobalStyles()
	if err != nil {
		return util.SmallError(code, message)
	}
	content = strings.ReplaceAll(content, util.Template("styles"), styles)
	content = strings.ReplaceAll(content, util.Template("code"), strconv.Itoa(code))
	return strings.ReplaceAll(content, util.Template("message"), message)
}
