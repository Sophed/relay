package util

import "strconv"

func Template(t string) string {
	return "{{" + t + "}}"
}

func SmallError(code int, message string) string {
	return strconv.Itoa(code) + " - " + message
}
