package errors

import (
	"fmt"
)

const (
	ERROR_DB = 3000
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Error) String() string {
	return fmt.Sprintf("%d,%s", e.Code, e.Msg)
}

func New(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func ThrowIfNotNil(err error, codes ...int) {
	if err != nil {
		code := 400
		if len(codes) > 0 {
			code = codes[0]
		}
		panic(New(code, err.Error()))
	}
}

func Throw(code int, msg string) {
	panic(New(code, msg))
}

func ThrowCommon(msg string) {
	panic(New(400, msg))
}
