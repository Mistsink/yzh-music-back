package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

var errCodes = make(map[int]struct{}, 10)

func NewError(code int, msg string) *Error {
	if _, ok := errCodes[code]; ok {
		panic(fmt.Sprintf("err code %d has been declared.\n", code))
	}
	errCodes[code] = struct{}{}
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	return e.Msg
}

func (e *Error) Errorf(data ...interface{}) string {
	return fmt.Sprintf(e.Msg, data...)
}

//	TO-TEST: 不知道复制前后的detail是否生效
func (e *Error) WithDetails(details []string) *Error {
	newE := *e
	newE.Details = details
	return &newE
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case NotFound.Code:
		return http.StatusNotFound
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenError.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
