package api

import (
	"encoding/json"
)

type ResponseError struct {
	ErrorMsg string `json:"error"`
}

func respErr(msg string) string {
	e := ResponseError{
		ErrorMsg: msg,
	}
	data, _ := json.Marshal(e)
	return string(data)
}

func respInternal() string {
	return respErr("Something went wrong on our end")
}
