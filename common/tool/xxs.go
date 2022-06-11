package tool

import (
	"golang.org/x/net/html"
)

func UnescapeString(str string) string {
	return html.UnescapeString(html.EscapeString(str))
}
