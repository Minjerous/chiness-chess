package errdef

import (
	"chess-common/httpcode"
	"fmt"
)

type Err struct {
	HttpCode    int
	StatusCode  int
	Description string
}

var Nil = New(httpcode.StatusAccepted, 0, "")

func (e Err) Error() string {
	return fmt.Sprintf("HttpCode: %d StatusCode: %d Description: %s\n", e.HttpCode, e.StatusCode, e.Description)
}

func New(httpCode, statusCode int, des string) Err {
	return Err{
		HttpCode:    httpCode,
		StatusCode:  statusCode,
		Description: des,
	}
}

func Errorf(httpCode, statusCode int, format string, a ...any) Err {
	return Err{
		HttpCode:    httpCode,
		StatusCode:  statusCode,
		Description: fmt.Sprintf(format, a...),
	}
}
