package web

import (
	"bytes"

	"github.com/sophed/lg"
	"maragu.dev/gomponents"
)

// Render converts a HTML node to a usable string which can be sent to clients
func Render(page gomponents.Node) string {
	buf := new(bytes.Buffer) // empty buffer
	err := page.Render(buf)  // render to buffer
	if err != nil {
		lg.Fatl(err) // panic on fail
	}
	return buf.String() // return buffer content as string
}
